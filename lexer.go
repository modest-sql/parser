package parser

import (
	"fmt"
	"strconv"
)
import (
	"bufio"
	"io"
	"strings"
)

type frame struct {
	i            int
	s            string
	line, column int
}
type Lexer struct {
	// The lexer runs in its own goroutine, and communicates via channel 'ch'.
	ch      chan frame
	ch_stop chan bool
	// We record the level of nesting because the action could return, and a
	// subsequent call expects to pick up where it left off. In other words,
	// we're simulating a coroutine.
	// TODO: Support a channel-based variant that compatible with Go's yacc.
	stack []frame
	stale bool

	// The 'l' and 'c' fields were added for
	// https://github.com/wagerlabs/docker/blob/65694e801a7b80930961d70c69cba9f2465459be/buildfile.nex
	// Since then, I introduced the built-in Line() and Column() functions.
	l, c int

	parseResult interface{}

	// The following line makes it easy for scripts to insert fields in the
	// generated code.
	// [NEX_END_OF_LEXER_STRUCT]
}

// NewLexerWithInit creates a new Lexer object, runs the given callback on it,
// then returns it.
func NewLexerWithInit(in io.Reader, initFun func(*Lexer)) *Lexer {
	yylex := new(Lexer)
	if initFun != nil {
		initFun(yylex)
	}
	yylex.ch = make(chan frame)
	yylex.ch_stop = make(chan bool, 1)
	var scan func(in *bufio.Reader, ch chan frame, ch_stop chan bool, family []dfa, line, column int)
	scan = func(in *bufio.Reader, ch chan frame, ch_stop chan bool, family []dfa, line, column int) {
		// Index of DFA and length of highest-precedence match so far.
		matchi, matchn := 0, -1
		var buf []rune
		n := 0
		checkAccept := func(i int, st int) bool {
			// Higher precedence match? DFAs are run in parallel, so matchn is at most len(buf), hence we may omit the length equality check.
			if family[i].acc[st] && (matchn < n || matchi > i) {
				matchi, matchn = i, n
				return true
			}
			return false
		}
		var state [][2]int
		for i := 0; i < len(family); i++ {
			mark := make([]bool, len(family[i].startf))
			// Every DFA starts at state 0.
			st := 0
			for {
				state = append(state, [2]int{i, st})
				mark[st] = true
				// As we're at the start of input, follow all ^ transitions and append to our list of start states.
				st = family[i].startf[st]
				if -1 == st || mark[st] {
					break
				}
				// We only check for a match after at least one transition.
				checkAccept(i, st)
			}
		}
		atEOF := false
		stopped := false
		for {
			if n == len(buf) && !atEOF {
				r, _, err := in.ReadRune()
				switch err {
				case io.EOF:
					atEOF = true
				case nil:
					buf = append(buf, r)
				default:
					panic(err)
				}
			}
			if !atEOF {
				r := buf[n]
				n++
				var nextState [][2]int
				for _, x := range state {
					x[1] = family[x[0]].f[x[1]](r)
					if -1 == x[1] {
						continue
					}
					nextState = append(nextState, x)
					checkAccept(x[0], x[1])
				}
				state = nextState
			} else {
			dollar: // Handle $.
				for _, x := range state {
					mark := make([]bool, len(family[x[0]].endf))
					for {
						mark[x[1]] = true
						x[1] = family[x[0]].endf[x[1]]
						if -1 == x[1] || mark[x[1]] {
							break
						}
						if checkAccept(x[0], x[1]) {
							// Unlike before, we can break off the search. Now that we're at the end, there's no need to maintain the state of each DFA.
							break dollar
						}
					}
				}
				state = nil
			}

			if state == nil {
				lcUpdate := func(r rune) {
					if r == '\n' {
						line++
						column = 0
					} else {
						column++
					}
				}
				// All DFAs stuck. Return last match if it exists, otherwise advance by one rune and restart all DFAs.
				if matchn == -1 {
					if len(buf) == 0 { // This can only happen at the end of input.
						break
					}
					lcUpdate(buf[0])
					buf = buf[1:]
				} else {
					text := string(buf[:matchn])
					buf = buf[matchn:]
					matchn = -1
					for {
						sent := false
						select {
						case ch <- frame{matchi, text, line, column}:
							{
								sent = true
							}
						case stopped = <-ch_stop:
							{
							}
						default:
							{
								// nothing
							}
						}
						if stopped || sent {
							break
						}
					}
					if stopped {
						break
					}
					if len(family[matchi].nest) > 0 {
						scan(bufio.NewReader(strings.NewReader(text)), ch, ch_stop, family[matchi].nest, line, column)
					}
					if atEOF {
						break
					}
					for _, r := range text {
						lcUpdate(r)
					}
				}
				n = 0
				for i := 0; i < len(family); i++ {
					state = append(state, [2]int{i, 0})
				}
			}
		}
		ch <- frame{-1, "", line, column}
	}
	go scan(bufio.NewReader(in), yylex.ch, yylex.ch_stop, dfas, 0, 0)
	return yylex
}

type dfa struct {
	acc          []bool           // Accepting states.
	f            []func(rune) int // Transitions.
	startf, endf []int            // Transitions at start and end of input.
	nest         []dfa
}

var dfas = []dfa{
	// [ \t\f\r\n]
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 9:
				return 1
			case 10:
				return 1
			case 12:
				return 1
			case 13:
				return 1
			case 32:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 9:
				return -1
			case 10:
				return -1
			case 12:
				return -1
			case 13:
				return -1
			case 32:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// --[^\n]*
	{[]bool{false, false, true, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 10:
				return -1
			case 45:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 10:
				return -1
			case 45:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 10:
				return -1
			case 45:
				return 3
			}
			return 3
		},
		func(r rune) int {
			switch r {
			case 10:
				return -1
			case 45:
				return 3
			}
			return 3
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

	// \/\*[^*]*\*\/
	{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 42:
				return -1
			case 47:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 42:
				return 2
			case 47:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 42:
				return 3
			case 47:
				return 4
			}
			return 4
		},
		func(r rune) int {
			switch r {
			case 42:
				return -1
			case 47:
				return 5
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 42:
				return 3
			case 47:
				return 4
			}
			return 4
		},
		func(r rune) int {
			switch r {
			case 42:
				return -1
			case 47:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

	// -?[0-9]+
	{[]bool{false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 45:
				return 1
			}
			switch {
			case 48 <= r && r <= 57:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 45:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 45:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 2
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

	// -?[0-9]+\.[0-9]*|-?\.[0-9]+|-?[0-9]+[Ee][-+]?[0-9]+|-?[0-9]+\.[0-9]*[Ee][-+]?[0-9]+|-?\.[0-9]*[Ee][-+]?[0-9]+
	{[]bool{false, false, false, false, true, false, false, true, false, true, false, true, false, true, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return 1
			case 46:
				return 2
			case 69:
				return -1
			case 101:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return 2
			case 69:
				return -1
			case 101:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return -1
			case 69:
				return 12
			case 101:
				return 12
			}
			switch {
			case 48 <= r && r <= 57:
				return 13
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return 4
			case 69:
				return 5
			case 101:
				return 5
			}
			switch {
			case 48 <= r && r <= 57:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return -1
			case 69:
				return 8
			case 101:
				return 8
			}
			switch {
			case 48 <= r && r <= 57:
				return 9
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return 6
			case 45:
				return 6
			case 46:
				return -1
			case 69:
				return -1
			case 101:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 7
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return -1
			case 69:
				return -1
			case 101:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 7
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return -1
			case 69:
				return -1
			case 101:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 7
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return 10
			case 45:
				return 10
			case 46:
				return -1
			case 69:
				return -1
			case 101:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 11
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return -1
			case 69:
				return 8
			case 101:
				return 8
			}
			switch {
			case 48 <= r && r <= 57:
				return 9
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return -1
			case 69:
				return -1
			case 101:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 11
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return -1
			case 69:
				return -1
			case 101:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 11
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return 14
			case 45:
				return 14
			case 46:
				return -1
			case 69:
				return -1
			case 101:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 15
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return -1
			case 69:
				return 12
			case 101:
				return 12
			}
			switch {
			case 48 <= r && r <= 57:
				return 13
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return -1
			case 69:
				return -1
			case 101:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 15
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			case 45:
				return -1
			case 46:
				return -1
			case 69:
				return -1
			case 101:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return 15
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

	// \+
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 43:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 43:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// \-
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 45:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 45:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// \*
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 42:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 42:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// \/
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 47:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 47:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// \(
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 40:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 40:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// \)
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 41:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 41:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// \.
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 46:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 46:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// ;
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 59:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 59:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// \,
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 44:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 44:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// =
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 61:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 61:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// \<
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 60:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 60:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// \>
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 62:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 62:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

	// \>=
	{[]bool{false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 61:
				return -1
			case 62:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 61:
				return 2
			case 62:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 61:
				return -1
			case 62:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

	// \<=
	{[]bool{false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 60:
				return 1
			case 61:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 60:
				return -1
			case 61:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 60:
				return -1
			case 61:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

	// \<\>
	{[]bool{false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 60:
				return 1
			case 62:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 60:
				return -1
			case 62:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 60:
				return -1
			case 62:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

	// [lL][iI][kK][eE]
	{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 75:
				return -1
			case 76:
				return 1
			case 101:
				return -1
			case 105:
				return -1
			case 107:
				return -1
			case 108:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return 2
			case 75:
				return -1
			case 76:
				return -1
			case 101:
				return -1
			case 105:
				return 2
			case 107:
				return -1
			case 108:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 75:
				return 3
			case 76:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 107:
				return 3
			case 108:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return 4
			case 73:
				return -1
			case 75:
				return -1
			case 76:
				return -1
			case 101:
				return 4
			case 105:
				return -1
			case 107:
				return -1
			case 108:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 75:
				return -1
			case 76:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 107:
				return -1
			case 108:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

	// [bB][eE][tT][wW][eE][eE][nN]
	{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 66:
				return 1
			case 69:
				return -1
			case 78:
				return -1
			case 84:
				return -1
			case 87:
				return -1
			case 98:
				return 1
			case 101:
				return -1
			case 110:
				return -1
			case 116:
				return -1
			case 119:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 66:
				return -1
			case 69:
				return 2
			case 78:
				return -1
			case 84:
				return -1
			case 87:
				return -1
			case 98:
				return -1
			case 101:
				return 2
			case 110:
				return -1
			case 116:
				return -1
			case 119:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 66:
				return -1
			case 69:
				return -1
			case 78:
				return -1
			case 84:
				return 3
			case 87:
				return -1
			case 98:
				return -1
			case 101:
				return -1
			case 110:
				return -1
			case 116:
				return 3
			case 119:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 66:
				return -1
			case 69:
				return -1
			case 78:
				return -1
			case 84:
				return -1
			case 87:
				return 4
			case 98:
				return -1
			case 101:
				return -1
			case 110:
				return -1
			case 116:
				return -1
			case 119:
				return 4
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 66:
				return -1
			case 69:
				return 5
			case 78:
				return -1
			case 84:
				return -1
			case 87:
				return -1
			case 98:
				return -1
			case 101:
				return 5
			case 110:
				return -1
			case 116:
				return -1
			case 119:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 66:
				return -1
			case 69:
				return 6
			case 78:
				return -1
			case 84:
				return -1
			case 87:
				return -1
			case 98:
				return -1
			case 101:
				return 6
			case 110:
				return -1
			case 116:
				return -1
			case 119:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 66:
				return -1
			case 69:
				return -1
			case 78:
				return 7
			case 84:
				return -1
			case 87:
				return -1
			case 98:
				return -1
			case 101:
				return -1
			case 110:
				return 7
			case 116:
				return -1
			case 119:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 66:
				return -1
			case 69:
				return -1
			case 78:
				return -1
			case 84:
				return -1
			case 87:
				return -1
			case 98:
				return -1
			case 101:
				return -1
			case 110:
				return -1
			case 116:
				return -1
			case 119:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

	// [aA][nN][dD]
	{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return 1
			case 68:
				return -1
			case 78:
				return -1
			case 97:
				return 1
			case 100:
				return -1
			case 110:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 78:
				return 2
			case 97:
				return -1
			case 100:
				return -1
			case 110:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return 3
			case 78:
				return -1
			case 97:
				return -1
			case 100:
				return 3
			case 110:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 78:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 110:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

	// [oO][rR]
	{[]bool{false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 79:
				return 1
			case 82:
				return -1
			case 111:
				return 1
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 79:
				return -1
			case 82:
				return 2
			case 111:
				return -1
			case 114:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 79:
				return -1
			case 82:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

	// [nN][oO][tT]
	{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 78:
				return 1
			case 79:
				return -1
			case 84:
				return -1
			case 110:
				return 1
			case 111:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 78:
				return -1
			case 79:
				return 2
			case 84:
				return -1
			case 110:
				return -1
			case 111:
				return 2
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 78:
				return -1
			case 79:
				return -1
			case 84:
				return 3
			case 110:
				return -1
			case 111:
				return -1
			case 116:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 78:
				return -1
			case 79:
				return -1
			case 84:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

	// [dD][eE][fF][aA][uU][lL][lT]
	{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return 1
			case 69:
				return -1
			case 70:
				return -1
			case 76:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 97:
				return -1
			case 100:
				return 1
			case 101:
				return -1
			case 102:
				return -1
			case 108:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return 2
			case 70:
				return -1
			case 76:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return 2
			case 102:
				return -1
			case 108:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 70:
				return 3
			case 76:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 102:
				return 3
			case 108:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 4
			case 68:
				return -1
			case 69:
				return -1
			case 70:
				return -1
			case 76:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 97:
				return 4
			case 100:
				return -1
			case 101:
				return -1
			case 102:
				return -1
			case 108:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 70:
				return -1
			case 76:
				return -1
			case 84:
				return -1
			case 85:
				return 5
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 102:
				return -1
			case 108:
				return -1
			case 117:
				return 5
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 70:
				return -1
			case 76:
				return 6
			case 84:
				return -1
			case 85:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 102:
				return -1
			case 108:
				return 6
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 70:
				return -1
			case 76:
				return -1
			case 84:
				return 7
			case 85:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 102:
				return -1
			case 108:
				return 7
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 70:
				return -1
			case 76:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 102:
				return -1
			case 108:
				return -1
			case 117:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

	// [iI][nN][tT][eE][gG][eE][rR]
	{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 71:
				return -1
			case 73:
				return 1
			case 78:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 101:
				return -1
			case 103:
				return -1
			case 105:
				return 1
			case 110:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 71:
				return -1
			case 73:
				return -1
			case 78:
				return 2
			case 82:
				return -1
			case 84:
				return -1
			case 101:
				return -1
			case 103:
				return -1
			case 105:
				return -1
			case 110:
				return 2
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 71:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return -1
			case 84:
				return 3
			case 101:
				return -1
			case 103:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return -1
			case 116:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return 4
			case 71:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 101:
				return 4
			case 103:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 71:
				return 5
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 101:
				return -1
			case 103:
				return 5
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return 6
			case 71:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 101:
				return 6
			case 103:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 71:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return 7
			case 84:
				return -1
			case 101:
				return -1
			case 103:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return 7
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 71:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 101:
				return -1
			case 103:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

	// [cC][hH][aA][rR]
	{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return 1
			case 72:
				return -1
			case 82:
				return -1
			case 97:
				return -1
			case 99:
				return 1
			case 104:
				return -1
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 72:
				return 2
			case 82:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 104:
				return 2
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 3
			case 67:
				return -1
			case 72:
				return -1
			case 82:
				return -1
			case 97:
				return 3
			case 99:
				return -1
			case 104:
				return -1
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 72:
				return -1
			case 82:
				return 4
			case 97:
				return -1
			case 99:
				return -1
			case 104:
				return -1
			case 114:
				return 4
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 72:
				return -1
			case 82:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 104:
				return -1
			case 114:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

	// [cC][rR][eE][aA][tT][eE]
	{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return 1
			case 69:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 99:
				return 1
			case 101:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 82:
				return 2
			case 84:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 114:
				return 2
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return 3
			case 82:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return 3
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 4
			case 67:
				return -1
			case 69:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 97:
				return 4
			case 99:
				return -1
			case 101:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 82:
				return -1
			case 84:
				return 5
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 114:
				return -1
			case 116:
				return 5
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return 6
			case 82:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return 6
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

	// [tT][aA][bB][lL][eE]
	{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 84:
				return 1
			case 97:
				return -1
			case 98:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 116:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 2
			case 66:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 84:
				return -1
			case 97:
				return 2
			case 98:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return 3
			case 69:
				return -1
			case 76:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 98:
				return 3
			case 101:
				return -1
			case 108:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return -1
			case 69:
				return -1
			case 76:
				return 4
			case 84:
				return -1
			case 97:
				return -1
			case 98:
				return -1
			case 101:
				return -1
			case 108:
				return 4
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return -1
			case 69:
				return 5
			case 76:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 98:
				return -1
			case 101:
				return 5
			case 108:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 98:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

	// [dD][eE][lL][eE][tT][eE]
	{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 68:
				return 1
			case 69:
				return -1
			case 76:
				return -1
			case 84:
				return -1
			case 100:
				return 1
			case 101:
				return -1
			case 108:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 68:
				return -1
			case 69:
				return 2
			case 76:
				return -1
			case 84:
				return -1
			case 100:
				return -1
			case 101:
				return 2
			case 108:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 68:
				return -1
			case 69:
				return -1
			case 76:
				return 3
			case 84:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 108:
				return 3
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 68:
				return -1
			case 69:
				return 4
			case 76:
				return -1
			case 84:
				return -1
			case 100:
				return -1
			case 101:
				return 4
			case 108:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 68:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 84:
				return 5
			case 100:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 116:
				return 5
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 68:
				return -1
			case 69:
				return 6
			case 76:
				return -1
			case 84:
				return -1
			case 100:
				return -1
			case 101:
				return 6
			case 108:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 68:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 84:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

	// [iI][nN][sS][eE][rR][tT]
	{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return 1
			case 78:
				return -1
			case 82:
				return -1
			case 83:
				return -1
			case 84:
				return -1
			case 101:
				return -1
			case 105:
				return 1
			case 110:
				return -1
			case 114:
				return -1
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 78:
				return 2
			case 82:
				return -1
			case 83:
				return -1
			case 84:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 110:
				return 2
			case 114:
				return -1
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return -1
			case 83:
				return 3
			case 84:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return -1
			case 115:
				return 3
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return 4
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return -1
			case 83:
				return -1
			case 84:
				return -1
			case 101:
				return 4
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return -1
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return 5
			case 83:
				return -1
			case 84:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return 5
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return -1
			case 83:
				return -1
			case 84:
				return 6
			case 101:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return -1
			case 115:
				return -1
			case 116:
				return 6
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return -1
			case 83:
				return -1
			case 84:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return -1
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

	// [iI][nN][tT][oO]
	{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 73:
				return 1
			case 78:
				return -1
			case 79:
				return -1
			case 84:
				return -1
			case 105:
				return 1
			case 110:
				return -1
			case 111:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 78:
				return 2
			case 79:
				return -1
			case 84:
				return -1
			case 105:
				return -1
			case 110:
				return 2
			case 111:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 84:
				return 3
			case 105:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 116:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 78:
				return -1
			case 79:
				return 4
			case 84:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 111:
				return 4
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 84:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

	// [vV][aA][lL][uU][eE][sS]
	{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 83:
				return -1
			case 85:
				return -1
			case 86:
				return 1
			case 97:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 115:
				return -1
			case 117:
				return -1
			case 118:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 2
			case 69:
				return -1
			case 76:
				return -1
			case 83:
				return -1
			case 85:
				return -1
			case 86:
				return -1
			case 97:
				return 2
			case 101:
				return -1
			case 108:
				return -1
			case 115:
				return -1
			case 117:
				return -1
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 76:
				return 3
			case 83:
				return -1
			case 85:
				return -1
			case 86:
				return -1
			case 97:
				return -1
			case 101:
				return -1
			case 108:
				return 3
			case 115:
				return -1
			case 117:
				return -1
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 83:
				return -1
			case 85:
				return 4
			case 86:
				return -1
			case 97:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 115:
				return -1
			case 117:
				return 4
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return 5
			case 76:
				return -1
			case 83:
				return -1
			case 85:
				return -1
			case 86:
				return -1
			case 97:
				return -1
			case 101:
				return 5
			case 108:
				return -1
			case 115:
				return -1
			case 117:
				return -1
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 83:
				return 6
			case 85:
				return -1
			case 86:
				return -1
			case 97:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 115:
				return 6
			case 117:
				return -1
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 83:
				return -1
			case 85:
				return -1
			case 86:
				return -1
			case 97:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 115:
				return -1
			case 117:
				return -1
			case 118:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

	// [sS][eE][lL][eE][cC][tT]
	{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 83:
				return 1
			case 84:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 115:
				return 1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 69:
				return 2
			case 76:
				return -1
			case 83:
				return -1
			case 84:
				return -1
			case 99:
				return -1
			case 101:
				return 2
			case 108:
				return -1
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 69:
				return -1
			case 76:
				return 3
			case 83:
				return -1
			case 84:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 108:
				return 3
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 69:
				return 4
			case 76:
				return -1
			case 83:
				return -1
			case 84:
				return -1
			case 99:
				return -1
			case 101:
				return 4
			case 108:
				return -1
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return 5
			case 69:
				return -1
			case 76:
				return -1
			case 83:
				return -1
			case 84:
				return -1
			case 99:
				return 5
			case 101:
				return -1
			case 108:
				return -1
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 83:
				return -1
			case 84:
				return 6
			case 99:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 115:
				return -1
			case 116:
				return 6
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 83:
				return -1
			case 84:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

	// [uU][pP][dD][aA][tT][eE]
	{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 80:
				return -1
			case 84:
				return -1
			case 85:
				return 1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 112:
				return -1
			case 116:
				return -1
			case 117:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 80:
				return 2
			case 84:
				return -1
			case 85:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 112:
				return 2
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return 3
			case 69:
				return -1
			case 80:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 97:
				return -1
			case 100:
				return 3
			case 101:
				return -1
			case 112:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 4
			case 68:
				return -1
			case 69:
				return -1
			case 80:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 97:
				return 4
			case 100:
				return -1
			case 101:
				return -1
			case 112:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 80:
				return -1
			case 84:
				return 5
			case 85:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 112:
				return -1
			case 116:
				return 5
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return 6
			case 80:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return 6
			case 112:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 80:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 112:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

	// [wW][hH][eE][rR][eE]
	{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 72:
				return -1
			case 82:
				return -1
			case 87:
				return 1
			case 101:
				return -1
			case 104:
				return -1
			case 114:
				return -1
			case 119:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 72:
				return 2
			case 82:
				return -1
			case 87:
				return -1
			case 101:
				return -1
			case 104:
				return 2
			case 114:
				return -1
			case 119:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return 3
			case 72:
				return -1
			case 82:
				return -1
			case 87:
				return -1
			case 101:
				return 3
			case 104:
				return -1
			case 114:
				return -1
			case 119:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 72:
				return -1
			case 82:
				return 4
			case 87:
				return -1
			case 101:
				return -1
			case 104:
				return -1
			case 114:
				return 4
			case 119:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return 5
			case 72:
				return -1
			case 82:
				return -1
			case 87:
				return -1
			case 101:
				return 5
			case 104:
				return -1
			case 114:
				return -1
			case 119:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 72:
				return -1
			case 82:
				return -1
			case 87:
				return -1
			case 101:
				return -1
			case 104:
				return -1
			case 114:
				return -1
			case 119:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

	// [tT][rR][uU][eE]
	{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 82:
				return -1
			case 84:
				return 1
			case 85:
				return -1
			case 101:
				return -1
			case 114:
				return -1
			case 116:
				return 1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 82:
				return 2
			case 84:
				return -1
			case 85:
				return -1
			case 101:
				return -1
			case 114:
				return 2
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return 3
			case 101:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return 4
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 101:
				return 4
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 101:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

	// [fF][aA][lL][sS][eE]
	{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 70:
				return 1
			case 76:
				return -1
			case 83:
				return -1
			case 97:
				return -1
			case 101:
				return -1
			case 102:
				return 1
			case 108:
				return -1
			case 115:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 2
			case 69:
				return -1
			case 70:
				return -1
			case 76:
				return -1
			case 83:
				return -1
			case 97:
				return 2
			case 101:
				return -1
			case 102:
				return -1
			case 108:
				return -1
			case 115:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 70:
				return -1
			case 76:
				return 3
			case 83:
				return -1
			case 97:
				return -1
			case 101:
				return -1
			case 102:
				return -1
			case 108:
				return 3
			case 115:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 70:
				return -1
			case 76:
				return -1
			case 83:
				return 4
			case 97:
				return -1
			case 101:
				return -1
			case 102:
				return -1
			case 108:
				return -1
			case 115:
				return 4
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return 5
			case 70:
				return -1
			case 76:
				return -1
			case 83:
				return -1
			case 97:
				return -1
			case 101:
				return 5
			case 102:
				return -1
			case 108:
				return -1
			case 115:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 70:
				return -1
			case 76:
				return -1
			case 83:
				return -1
			case 97:
				return -1
			case 101:
				return -1
			case 102:
				return -1
			case 108:
				return -1
			case 115:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

	// [fF][rR][oO][mM]
	{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 70:
				return 1
			case 77:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 102:
				return 1
			case 109:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 70:
				return -1
			case 77:
				return -1
			case 79:
				return -1
			case 82:
				return 2
			case 102:
				return -1
			case 109:
				return -1
			case 111:
				return -1
			case 114:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 70:
				return -1
			case 77:
				return -1
			case 79:
				return 3
			case 82:
				return -1
			case 102:
				return -1
			case 109:
				return -1
			case 111:
				return 3
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 70:
				return -1
			case 77:
				return 4
			case 79:
				return -1
			case 82:
				return -1
			case 102:
				return -1
			case 109:
				return 4
			case 111:
				return -1
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 70:
				return -1
			case 77:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 102:
				return -1
			case 109:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

	// [sS][eE][tT]
	{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 83:
				return 1
			case 84:
				return -1
			case 101:
				return -1
			case 115:
				return 1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return 2
			case 83:
				return -1
			case 84:
				return -1
			case 101:
				return 2
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 83:
				return -1
			case 84:
				return 3
			case 101:
				return -1
			case 115:
				return -1
			case 116:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 83:
				return -1
			case 84:
				return -1
			case 101:
				return -1
			case 115:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

	// [dD][rR][oO][pP]
	{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 68:
				return 1
			case 79:
				return -1
			case 80:
				return -1
			case 82:
				return -1
			case 100:
				return 1
			case 111:
				return -1
			case 112:
				return -1
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 68:
				return -1
			case 79:
				return -1
			case 80:
				return -1
			case 82:
				return 2
			case 100:
				return -1
			case 111:
				return -1
			case 112:
				return -1
			case 114:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 68:
				return -1
			case 79:
				return 3
			case 80:
				return -1
			case 82:
				return -1
			case 100:
				return -1
			case 111:
				return 3
			case 112:
				return -1
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 68:
				return -1
			case 79:
				return -1
			case 80:
				return 4
			case 82:
				return -1
			case 100:
				return -1
			case 111:
				return -1
			case 112:
				return 4
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 68:
				return -1
			case 79:
				return -1
			case 80:
				return -1
			case 82:
				return -1
			case 100:
				return -1
			case 111:
				return -1
			case 112:
				return -1
			case 114:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

	// [iI][nN][nN][eE][rR]
	{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return 1
			case 78:
				return -1
			case 82:
				return -1
			case 101:
				return -1
			case 105:
				return 1
			case 110:
				return -1
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 78:
				return 2
			case 82:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 110:
				return 2
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 78:
				return 3
			case 82:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 110:
				return 3
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return 4
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return -1
			case 101:
				return 4
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return 5
			case 101:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return 5
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 69:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 82:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 114:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

	// [jJ][oO][iI][nN]
	{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 74:
				return 1
			case 78:
				return -1
			case 79:
				return -1
			case 105:
				return -1
			case 106:
				return 1
			case 110:
				return -1
			case 111:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 74:
				return -1
			case 78:
				return -1
			case 79:
				return 2
			case 105:
				return -1
			case 106:
				return -1
			case 110:
				return -1
			case 111:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return 3
			case 74:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 105:
				return 3
			case 106:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 74:
				return -1
			case 78:
				return 4
			case 79:
				return -1
			case 105:
				return -1
			case 106:
				return -1
			case 110:
				return 4
			case 111:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 74:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 105:
				return -1
			case 106:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

	// [hH][aA][vV][iI][nN][gG]
	{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 71:
				return -1
			case 72:
				return 1
			case 73:
				return -1
			case 78:
				return -1
			case 86:
				return -1
			case 97:
				return -1
			case 103:
				return -1
			case 104:
				return 1
			case 105:
				return -1
			case 110:
				return -1
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 2
			case 71:
				return -1
			case 72:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 86:
				return -1
			case 97:
				return 2
			case 103:
				return -1
			case 104:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 71:
				return -1
			case 72:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 86:
				return 3
			case 97:
				return -1
			case 103:
				return -1
			case 104:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 118:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 71:
				return -1
			case 72:
				return -1
			case 73:
				return 4
			case 78:
				return -1
			case 86:
				return -1
			case 97:
				return -1
			case 103:
				return -1
			case 104:
				return -1
			case 105:
				return 4
			case 110:
				return -1
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 71:
				return -1
			case 72:
				return -1
			case 73:
				return -1
			case 78:
				return 5
			case 86:
				return -1
			case 97:
				return -1
			case 103:
				return -1
			case 104:
				return -1
			case 105:
				return -1
			case 110:
				return 5
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 71:
				return 6
			case 72:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 86:
				return -1
			case 97:
				return -1
			case 103:
				return 6
			case 104:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 71:
				return -1
			case 72:
				return -1
			case 73:
				return -1
			case 78:
				return -1
			case 86:
				return -1
			case 97:
				return -1
			case 103:
				return -1
			case 104:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			case 118:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

	// [sS][uU][mM]
	{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 77:
				return -1
			case 83:
				return 1
			case 85:
				return -1
			case 109:
				return -1
			case 115:
				return 1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 77:
				return -1
			case 83:
				return -1
			case 85:
				return 2
			case 109:
				return -1
			case 115:
				return -1
			case 117:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 77:
				return 3
			case 83:
				return -1
			case 85:
				return -1
			case 109:
				return 3
			case 115:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 77:
				return -1
			case 83:
				return -1
			case 85:
				return -1
			case 109:
				return -1
			case 115:
				return -1
			case 117:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

	// [cC][oO][uU][nN][tT]
	{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 67:
				return 1
			case 78:
				return -1
			case 79:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 99:
				return 1
			case 110:
				return -1
			case 111:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 78:
				return -1
			case 79:
				return 2
			case 84:
				return -1
			case 85:
				return -1
			case 99:
				return -1
			case 110:
				return -1
			case 111:
				return 2
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 84:
				return -1
			case 85:
				return 3
			case 99:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 116:
				return -1
			case 117:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 78:
				return 4
			case 79:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 99:
				return -1
			case 110:
				return 4
			case 111:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 84:
				return 5
			case 85:
				return -1
			case 99:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 116:
				return 5
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 99:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

	// [aA][vV][gG]
	{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return 1
			case 71:
				return -1
			case 86:
				return -1
			case 97:
				return 1
			case 103:
				return -1
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 71:
				return -1
			case 86:
				return 2
			case 97:
				return -1
			case 103:
				return -1
			case 118:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 71:
				return 3
			case 86:
				return -1
			case 97:
				return -1
			case 103:
				return 3
			case 118:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 71:
				return -1
			case 86:
				return -1
			case 97:
				return -1
			case 103:
				return -1
			case 118:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

	// [mM][iI][nN]
	{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 77:
				return 1
			case 78:
				return -1
			case 105:
				return -1
			case 109:
				return 1
			case 110:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return 2
			case 77:
				return -1
			case 78:
				return -1
			case 105:
				return 2
			case 109:
				return -1
			case 110:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return 3
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

	// [mM][aA][xX]
	{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 77:
				return 1
			case 88:
				return -1
			case 97:
				return -1
			case 109:
				return 1
			case 120:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 2
			case 77:
				return -1
			case 88:
				return -1
			case 97:
				return 2
			case 109:
				return -1
			case 120:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 77:
				return -1
			case 88:
				return 3
			case 97:
				return -1
			case 109:
				return -1
			case 120:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 77:
				return -1
			case 88:
				return -1
			case 97:
				return -1
			case 109:
				return -1
			case 120:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

	// [aA][lL][tT][eE][rR]
	{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return 1
			case 69:
				return -1
			case 76:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 97:
				return 1
			case 101:
				return -1
			case 108:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 76:
				return 2
			case 82:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 101:
				return -1
			case 108:
				return 2
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 82:
				return -1
			case 84:
				return 3
			case 97:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 114:
				return -1
			case 116:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return 4
			case 76:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 101:
				return 4
			case 108:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 82:
				return 5
			case 84:
				return -1
			case 97:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 114:
				return 5
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

	// [nN][uU][lL][lL]
	{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 76:
				return -1
			case 78:
				return 1
			case 85:
				return -1
			case 108:
				return -1
			case 110:
				return 1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 76:
				return -1
			case 78:
				return -1
			case 85:
				return 2
			case 108:
				return -1
			case 110:
				return -1
			case 117:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 76:
				return 3
			case 78:
				return -1
			case 85:
				return -1
			case 108:
				return 3
			case 110:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 76:
				return 4
			case 78:
				return -1
			case 85:
				return -1
			case 108:
				return 4
			case 110:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 76:
				return -1
			case 78:
				return -1
			case 85:
				return -1
			case 108:
				return -1
			case 110:
				return -1
			case 117:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

	// [iI][nN]
	{[]bool{false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 73:
				return 1
			case 78:
				return -1
			case 105:
				return 1
			case 110:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 78:
				return 2
			case 105:
				return -1
			case 110:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 78:
				return -1
			case 105:
				return -1
			case 110:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

	// [iI][sS]
	{[]bool{false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 73:
				return 1
			case 83:
				return -1
			case 105:
				return 1
			case 115:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 83:
				return 2
			case 105:
				return -1
			case 115:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 73:
				return -1
			case 83:
				return -1
			case 105:
				return -1
			case 115:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

	// [aA][sS]
	{[]bool{false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return 1
			case 83:
				return -1
			case 97:
				return 1
			case 115:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 83:
				return 2
			case 97:
				return -1
			case 115:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 83:
				return -1
			case 97:
				return -1
			case 115:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

	// [bB][yY]
	{[]bool{false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 66:
				return 1
			case 89:
				return -1
			case 98:
				return 1
			case 121:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 66:
				return -1
			case 89:
				return 2
			case 98:
				return -1
			case 121:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 66:
				return -1
			case 89:
				return -1
			case 98:
				return -1
			case 121:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

	// [bB][oO][oO][lL][eE][aA][nN]
	{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return 1
			case 69:
				return -1
			case 76:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 97:
				return -1
			case 98:
				return 1
			case 101:
				return -1
			case 108:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 78:
				return -1
			case 79:
				return 2
			case 97:
				return -1
			case 98:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 110:
				return -1
			case 111:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 78:
				return -1
			case 79:
				return 3
			case 97:
				return -1
			case 98:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 110:
				return -1
			case 111:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return -1
			case 69:
				return -1
			case 76:
				return 4
			case 78:
				return -1
			case 79:
				return -1
			case 97:
				return -1
			case 98:
				return -1
			case 101:
				return -1
			case 108:
				return 4
			case 110:
				return -1
			case 111:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return -1
			case 69:
				return 5
			case 76:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 97:
				return -1
			case 98:
				return -1
			case 101:
				return 5
			case 108:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 6
			case 66:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 97:
				return 6
			case 98:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 78:
				return 7
			case 79:
				return -1
			case 97:
				return -1
			case 98:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 110:
				return 7
			case 111:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 66:
				return -1
			case 69:
				return -1
			case 76:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 97:
				return -1
			case 98:
				return -1
			case 101:
				return -1
			case 108:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

	// [oO][nN]
	{[]bool{false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 78:
				return -1
			case 79:
				return 1
			case 110:
				return -1
			case 111:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 78:
				return 2
			case 79:
				return -1
			case 110:
				return 2
			case 111:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 78:
				return -1
			case 79:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

	// [aA][dD][dD]
	{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return 1
			case 68:
				return -1
			case 97:
				return 1
			case 100:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return 2
			case 97:
				return -1
			case 100:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return 3
			case 97:
				return -1
			case 100:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

	// [cC][oO][lL][uU][mM][nN]
	{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 67:
				return 1
			case 76:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 85:
				return -1
			case 99:
				return 1
			case 108:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 76:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return 2
			case 85:
				return -1
			case 99:
				return -1
			case 108:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return 2
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 76:
				return 3
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 85:
				return -1
			case 99:
				return -1
			case 108:
				return 3
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 76:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 85:
				return 4
			case 99:
				return -1
			case 108:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 117:
				return 4
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 76:
				return -1
			case 77:
				return 5
			case 78:
				return -1
			case 79:
				return -1
			case 85:
				return -1
			case 99:
				return -1
			case 108:
				return -1
			case 109:
				return 5
			case 110:
				return -1
			case 111:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 76:
				return -1
			case 77:
				return -1
			case 78:
				return 6
			case 79:
				return -1
			case 85:
				return -1
			case 99:
				return -1
			case 108:
				return -1
			case 109:
				return -1
			case 110:
				return 6
			case 111:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 67:
				return -1
			case 76:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 85:
				return -1
			case 99:
				return -1
			case 108:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 117:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

	// [gG][rR][oO][uU][pP]
	{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 71:
				return 1
			case 79:
				return -1
			case 80:
				return -1
			case 82:
				return -1
			case 85:
				return -1
			case 103:
				return 1
			case 111:
				return -1
			case 112:
				return -1
			case 114:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 71:
				return -1
			case 79:
				return -1
			case 80:
				return -1
			case 82:
				return 2
			case 85:
				return -1
			case 103:
				return -1
			case 111:
				return -1
			case 112:
				return -1
			case 114:
				return 2
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 71:
				return -1
			case 79:
				return 3
			case 80:
				return -1
			case 82:
				return -1
			case 85:
				return -1
			case 103:
				return -1
			case 111:
				return 3
			case 112:
				return -1
			case 114:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 71:
				return -1
			case 79:
				return -1
			case 80:
				return -1
			case 82:
				return -1
			case 85:
				return 4
			case 103:
				return -1
			case 111:
				return -1
			case 112:
				return -1
			case 114:
				return -1
			case 117:
				return 4
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 71:
				return -1
			case 79:
				return -1
			case 80:
				return 5
			case 82:
				return -1
			case 85:
				return -1
			case 103:
				return -1
			case 111:
				return -1
			case 112:
				return 5
			case 114:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 71:
				return -1
			case 79:
				return -1
			case 80:
				return -1
			case 82:
				return -1
			case 85:
				return -1
			case 103:
				return -1
			case 111:
				return -1
			case 112:
				return -1
			case 114:
				return -1
			case 117:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

	// [dD][aA][tT][eE][tT][iI][mM][eE]
	{[]bool{false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return 1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 100:
				return 1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 2
			case 68:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 84:
				return -1
			case 97:
				return 2
			case 100:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 84:
				return 3
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 116:
				return 3
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return 4
			case 73:
				return -1
			case 77:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return 4
			case 105:
				return -1
			case 109:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 84:
				return 5
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 116:
				return 5
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 73:
				return 6
			case 77:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 105:
				return 6
			case 109:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return 7
			case 84:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return 7
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return 8
			case 73:
				return -1
			case 77:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return 8
			case 105:
				return -1
			case 109:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 68:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 100:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

	// [fF][lL][oO][aA][tT]
	{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 70:
				return 1
			case 76:
				return -1
			case 79:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 102:
				return 1
			case 108:
				return -1
			case 111:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 70:
				return -1
			case 76:
				return 2
			case 79:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 102:
				return -1
			case 108:
				return 2
			case 111:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 70:
				return -1
			case 76:
				return -1
			case 79:
				return 3
			case 84:
				return -1
			case 97:
				return -1
			case 102:
				return -1
			case 108:
				return -1
			case 111:
				return 3
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return 4
			case 70:
				return -1
			case 76:
				return -1
			case 79:
				return -1
			case 84:
				return -1
			case 97:
				return 4
			case 102:
				return -1
			case 108:
				return -1
			case 111:
				return -1
			case 116:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 70:
				return -1
			case 76:
				return -1
			case 79:
				return -1
			case 84:
				return 5
			case 97:
				return -1
			case 102:
				return -1
			case 108:
				return -1
			case 111:
				return -1
			case 116:
				return 5
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 70:
				return -1
			case 76:
				return -1
			case 79:
				return -1
			case 84:
				return -1
			case 97:
				return -1
			case 102:
				return -1
			case 108:
				return -1
			case 111:
				return -1
			case 116:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

	// [aA][uU][tT][oO]_[iI][nN][cC][rR][eE][mM][eE][nN][tT]
	{[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 65:
				return 1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return 1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return 2
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return 3
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return 3
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return 4
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return 4
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return 5
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return 6
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return 6
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return 7
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return 7
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return 8
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return 8
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return 9
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return 9
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return 10
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return 10
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return 11
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return 11
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return 12
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return 12
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return 13
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return 13
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return 14
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return 14
			case 117:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 65:
				return -1
			case 67:
				return -1
			case 69:
				return -1
			case 73:
				return -1
			case 77:
				return -1
			case 78:
				return -1
			case 79:
				return -1
			case 82:
				return -1
			case 84:
				return -1
			case 85:
				return -1
			case 95:
				return -1
			case 97:
				return -1
			case 99:
				return -1
			case 101:
				return -1
			case 105:
				return -1
			case 109:
				return -1
			case 110:
				return -1
			case 111:
				return -1
			case 114:
				return -1
			case 116:
				return -1
			case 117:
				return -1
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

	// '[^']*'
	{[]bool{false, false, true, false}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 39:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 39:
				return 2
			}
			return 3
		},
		func(r rune) int {
			switch r {
			case 39:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 39:
				return 2
			}
			return 3
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

	// "[^"]*"
	{[]bool{false, false, true, false}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 34:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 34:
				return 2
			}
			return 3
		},
		func(r rune) int {
			switch r {
			case 34:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 34:
				return 2
			}
			return 3
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

	// [a-zA-Z][a-zA-Z0-9_]*
	{[]bool{false, true, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 95:
				return -1
			}
			switch {
			case 48 <= r && r <= 57:
				return -1
			case 65 <= r && r <= 90:
				return 1
			case 97 <= r && r <= 122:
				return 1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 95:
				return 2
			}
			switch {
			case 48 <= r && r <= 57:
				return 2
			case 65 <= r && r <= 90:
				return 2
			case 97 <= r && r <= 122:
				return 2
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 95:
				return 2
			}
			switch {
			case 48 <= r && r <= 57:
				return 2
			case 65 <= r && r <= 90:
				return 2
			case 97 <= r && r <= 122:
				return 2
			}
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

	// '(\\.|[^'])*$
	{[]bool{false, false, false, false, true, false, false, false}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 39:
				return 1
			case 92:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 39:
				return -1
			case 92:
				return 2
			}
			return 3
		},
		func(r rune) int {
			switch r {
			case 39:
				return 5
			case 92:
				return 6
			}
			return 7
		},
		func(r rune) int {
			switch r {
			case 39:
				return -1
			case 92:
				return 2
			}
			return 3
		},
		func(r rune) int {
			switch r {
			case 39:
				return -1
			case 92:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 39:
				return -1
			case 92:
				return 2
			}
			return 3
		},
		func(r rune) int {
			switch r {
			case 39:
				return 5
			case 92:
				return 6
			}
			return 7
		},
		func(r rune) int {
			switch r {
			case 39:
				return -1
			case 92:
				return 2
			}
			return 3
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, 4, 4, 4, -1, 4, 4, 4}, nil},

	// \"(\\.|[^"])*$
	{[]bool{false, false, false, false, true, false, false, false}, []func(rune) int{ // Transitions
		func(r rune) int {
			switch r {
			case 34:
				return 1
			case 92:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 34:
				return -1
			case 92:
				return 2
			}
			return 3
		},
		func(r rune) int {
			switch r {
			case 34:
				return 5
			case 92:
				return 6
			}
			return 7
		},
		func(r rune) int {
			switch r {
			case 34:
				return -1
			case 92:
				return 2
			}
			return 3
		},
		func(r rune) int {
			switch r {
			case 34:
				return -1
			case 92:
				return -1
			}
			return -1
		},
		func(r rune) int {
			switch r {
			case 34:
				return -1
			case 92:
				return 2
			}
			return 3
		},
		func(r rune) int {
			switch r {
			case 34:
				return 5
			case 92:
				return 6
			}
			return 7
		},
		func(r rune) int {
			switch r {
			case 34:
				return -1
			case 92:
				return 2
			}
			return 3
		},
	}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, 4, 4, 4, -1, 4, 4, 4}, nil},

	// .
	{[]bool{false, true}, []func(rune) int{ // Transitions
		func(r rune) int {
			return 1
		},
		func(r rune) int {
			return -1
		},
	}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},
}

func NewLexer(in io.Reader) *Lexer {
	return NewLexerWithInit(in, nil)
}

func (yyLex *Lexer) Stop() {
	yyLex.ch_stop <- true
}

// Text returns the matched text.
func (yylex *Lexer) Text() string {
	return yylex.stack[len(yylex.stack)-1].s
}

// Line returns the current line number.
// The first line is 0.
func (yylex *Lexer) Line() int {
	if len(yylex.stack) == 0 {
		return 0
	}
	return yylex.stack[len(yylex.stack)-1].line
}

// Column returns the current column number.
// The first column is 0.
func (yylex *Lexer) Column() int {
	if len(yylex.stack) == 0 {
		return 0
	}
	return yylex.stack[len(yylex.stack)-1].column
}

func (yylex *Lexer) next(lvl int) int {
	if lvl == len(yylex.stack) {
		l, c := 0, 0
		if lvl > 0 {
			l, c = yylex.stack[lvl-1].line, yylex.stack[lvl-1].column
		}
		yylex.stack = append(yylex.stack, frame{0, "", l, c})
	}
	if lvl == len(yylex.stack)-1 {
		p := &yylex.stack[lvl]
		*p = <-yylex.ch
		yylex.stale = false
	} else {
		yylex.stale = true
	}
	return yylex.stack[lvl].i
}
func (yylex *Lexer) pop() {
	yylex.stack = yylex.stack[:len(yylex.stack)-1]
}

// Lex runs the lexer. Always returns 0.
// When the -s option is given, this function is not generated;
// instead, the NN_FUN macro runs the lexer.
func (yylex *Lexer) Lex(lval *yySymType) int {
OUTER0:
	for {
		switch yylex.next(0) {
		case 0:
			{ /* Skip white space */
			}
		case 1:
			{ /* Single line comment */
			}
		case 2:
			{ /* Multi line comment */
			}
		case 3:
			{
				lval.int_t, _ = strconv.Atoi(yylex.Text())
				return INT_LIT
			}
		case 4:
			{
				lval.float_t, _ = strconv.ParseFloat(yylex.Text(), 64)
				return FLOAT_LIT
			}
		case 5:
			{
				return TK_PLUS
			}
		case 6:
			{
				return TK_MINUS
			}
		case 7:
			{
				return TK_STAR
			}
		case 8:
			{
				return TK_DIV
			}
		case 9:
			{
				return TK_LEFT_PAR
			}
		case 10:
			{
				return TK_RIGHT_PAR
			}
		case 11:
			{
				return TK_DOT
			}
		case 12:
			{
				return TK_SEMICOLON
			}
		case 13:
			{
				return TK_COMMA
			}
		case 14:
			{
				return TK_EQ
			}
		case 15:
			{
				return TK_LT
			}
		case 16:
			{
				return TK_GT
			}
		case 17:
			{
				return TK_GTE
			}
		case 18:
			{
				return TK_LTE
			}
		case 19:
			{
				return TK_NE
			}
		case 20:
			{
				return KW_LIKE
			}
		case 21:
			{
				return KW_BETWEEN
			}
		case 22:
			{
				return KW_AND
			}
		case 23:
			{
				return KW_OR
			}
		case 24:
			{
				return KW_NOT
			}
		case 25:
			{
				return KW_DEFAULT
			}
		case 26:
			{
				return KW_INTEGER
			}
		case 27:
			{
				return KW_CHAR
			}
		case 28:
			{
				return KW_CREATE
			}
		case 29:
			{
				return KW_TABLE
			}
		case 30:
			{
				return KW_DELETE
			}
		case 31:
			{
				return KW_INSERT
			}
		case 32:
			{
				return KW_INTO
			}
		case 33:
			{
				return KW_VALUES
			}
		case 34:
			{
				return KW_SELECT
			}
		case 35:
			{
				return KW_UPDATE
			}
		case 36:
			{
				return KW_WHERE
			}
		case 37:
			{
				return KW_TRUE
			}
		case 38:
			{
				return KW_FALSE
			}
		case 39:
			{
				return KW_FROM
			}
		case 40:
			{
				return KW_SET
			}
		case 41:
			{
				return KW_DROP
			}
		case 42:
			{
				return KW_INNER
			}
		case 43:
			{
				return KW_JOIN
			}
		case 44:
			{
				return KW_HAVING
			}
		case 45:
			{
				return KW_SUM
			}
		case 46:
			{
				return KW_COUNT
			}
		case 47:
			{
				return KW_AVG
			}
		case 48:
			{
				return KW_MIN
			}
		case 49:
			{
				return KW_MAX
			}
		case 50:
			{
				return KW_ALTER
			}
		case 51:
			{
				return KW_NULL
			}
		case 52:
			{
				return KW_IN
			}
		case 53:
			{
				return KW_IS
			}
		case 54:
			{
				return KW_AS
			}
		case 55:
			{
				return KW_BY
			}
		case 56:
			{
				return KW_BOOLEAN
			}
		case 57:
			{
				return KW_ON
			}
		case 58:
			{
				return KW_ADD
			}
		case 59:
			{
				return KW_COLUMN
			}
		case 60:
			{
				return KW_GROUP
			}
		case 61:
			{
				return KW_DATETIME
			}
		case 62:
			{
				return KW_FLOAT
			}
		case 63:
			{
				return KW_AUTO_INCREMENT
			}
		case 64:
			{
				lval.string_t = strings.Replace(yylex.Text(), "'", "", -1)
				return STR_LIT
			}
		case 65:
			{
				lval.string_t = strings.Replace(yylex.Text(), "\"", "", -1)
				return STR_LIT
			}
		case 66:
			{
				lval.string_t = strings.ToLower(yylex.Text())
				return TK_ID
			}
		case 67:
			{
				panic(Error{
					line:    yylex.Line() + 1,
					column:  yylex.Column() + 1,
					message: fmt.Sprintf("syntax error: unterminated string literal"),
				})
			}
		case 68:
			{
				panic(Error{
					line:    yylex.Line() + 1,
					column:  yylex.Column() + 1,
					message: fmt.Sprintf("syntax error: unterminated string literal"),
				})
			}
		case 69:
			{
				panic(Error{
					line:    yylex.Line() + 1,
					column:  yylex.Column() + 1,
					message: fmt.Sprintf("syntax error: unexpected token '%s'", yylex.Text()),
				})
			}
		default:
			break OUTER0
		}
		continue
	}
	yylex.pop()

	return 0
}

//line parser.y:4
package parser

import __yyfmt__ "fmt"

//line parser.y:4
import "io"

/*import "github.com/modest-sql/common"*/

var statements statementList

//line parser.y:13
type yySymType struct {
	yys       int
	int64_t   int64
	string_t  string
	float64_t float64

	expr_t expression

	stmt_list_t       statementList
	stmt_t            statement
	columnSpec_t      columnSpec
	columnSpec_list_t []columnSpec
	col_list_t        columnDefinitions
	col_t             *columnDefinition
	assignment_t      assignment
	data_t            dataType
	assignments_list  []assignment
	obj_list_t        []interface{}
	string_list_t     []string

	obj_t interface{}
}

const TK_PLUS = 57346
const TK_MINUS = 57347
const TK_STAR = 57348
const TK_DIV = 57349
const TK_LT = 57350
const TK_GT = 57351
const TK_GTE = 57352
const TK_LTE = 57353
const TK_EQ = 57354
const TK_NE = 57355
const TK_LEFT_PAR = 57356
const TK_RIGHT_PAR = 57357
const TK_COMMA = 57358
const TK_DOT = 57359
const TK_SEMICOLON = 57360
const KW_OR = 57361
const KW_AND = 57362
const KW_NOT = 57363
const KW_INTEGER = 57364
const KW_FLOAT = 57365
const KW_CHAR = 57366
const KW_BOOLEAN = 57367
const KW_DATETIME = 57368
const KW_CREATE = 57369
const KW_TABLE = 57370
const KW_DELETE = 57371
const KW_INSERT = 57372
const KW_INTO = 57373
const KW_SELECT = 57374
const KW_WHERE = 57375
const KW_FROM = 57376
const KW_UPDATE = 57377
const KW_SET = 57378
const KW_ALTER = 57379
const KW_VALUES = 57380
const KW_BETWEEN = 57381
const KW_LIKE = 57382
const KW_INNER = 57383
const KW_HAVING = 57384
const KW_SUM = 57385
const KW_COUNT = 57386
const KW_AVG = 57387
const KW_MIN = 57388
const KW_MAX = 57389
const KW_NULL = 57390
const KW_IN = 57391
const KW_IS = 57392
const KW_AUTO_INCREMENT = 57393
const KW_JOIN = 57394
const KW_ON = 57395
const KW_GROUP = 57396
const KW_BY = 57397
const KW_DROP = 57398
const KW_DEFAULT = 57399
const KW_TRUE = 57400
const KW_FALSE = 57401
const KW_AS = 57402
const KW_ADD = 57403
const KW_COLUMN = 57404
const INT_LIT = 57405
const FLOAT_LIT = 57406
const TK_ID = 57407
const STR_LIT = 57408

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TK_PLUS",
	"TK_MINUS",
	"TK_STAR",
	"TK_DIV",
	"TK_LT",
	"TK_GT",
	"TK_GTE",
	"TK_LTE",
	"TK_EQ",
	"TK_NE",
	"TK_LEFT_PAR",
	"TK_RIGHT_PAR",
	"TK_COMMA",
	"TK_DOT",
	"TK_SEMICOLON",
	"KW_OR",
	"KW_AND",
	"KW_NOT",
	"KW_INTEGER",
	"KW_FLOAT",
	"KW_CHAR",
	"KW_BOOLEAN",
	"KW_DATETIME",
	"KW_CREATE",
	"KW_TABLE",
	"KW_DELETE",
	"KW_INSERT",
	"KW_INTO",
	"KW_SELECT",
	"KW_WHERE",
	"KW_FROM",
	"KW_UPDATE",
	"KW_SET",
	"KW_ALTER",
	"KW_VALUES",
	"KW_BETWEEN",
	"KW_LIKE",
	"KW_INNER",
	"KW_HAVING",
	"KW_SUM",
	"KW_COUNT",
	"KW_AVG",
	"KW_MIN",
	"KW_MAX",
	"KW_NULL",
	"KW_IN",
	"KW_IS",
	"KW_AUTO_INCREMENT",
	"KW_JOIN",
	"KW_ON",
	"KW_GROUP",
	"KW_BY",
	"KW_DROP",
	"KW_DEFAULT",
	"KW_TRUE",
	"KW_FALSE",
	"KW_AS",
	"KW_ADD",
	"KW_COLUMN",
	"INT_LIT",
	"FLOAT_LIT",
	"TK_ID",
	"STR_LIT",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser.y:262

func init() {
	yyErrorVerbose = true
}

func (l *Lexer) Error(s string) {
	panic(Error{
		line:    l.Line() + 1,
		column:  l.Column() + 1,
		message: s,
	})
}

func Parse(in io.Reader) (commands []interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	yyParse(NewLexer(in))

	return statements.convert(), nil
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 203

var yyAct = [...]int{

	166, 151, 161, 106, 62, 150, 75, 64, 136, 42,
	61, 69, 36, 63, 59, 43, 58, 168, 60, 111,
	167, 39, 110, 69, 37, 24, 108, 147, 145, 38,
	144, 142, 76, 85, 53, 52, 49, 48, 47, 46,
	40, 28, 37, 27, 169, 72, 163, 38, 126, 112,
	109, 55, 146, 153, 73, 70, 71, 72, 81, 83,
	65, 165, 68, 67, 37, 45, 43, 70, 71, 38,
	34, 88, 65, 26, 68, 67, 31, 79, 30, 23,
	102, 101, 80, 155, 25, 29, 148, 87, 33, 154,
	86, 32, 113, 21, 118, 119, 120, 121, 122, 123,
	124, 35, 117, 116, 177, 178, 172, 129, 130, 133,
	132, 127, 128, 17, 50, 15, 14, 39, 13, 143,
	103, 16, 175, 18, 138, 141, 137, 139, 140, 41,
	89, 90, 93, 92, 91, 94, 104, 131, 51, 149,
	134, 135, 19, 114, 115, 162, 156, 77, 54, 157,
	1, 82, 164, 158, 159, 99, 100, 97, 98, 164,
	22, 96, 95, 174, 170, 171, 89, 90, 93, 92,
	91, 94, 164, 164, 3, 176, 84, 20, 56, 179,
	57, 125, 66, 44, 74, 78, 152, 160, 173, 107,
	105, 9, 8, 7, 6, 12, 11, 96, 95, 10,
	5, 4, 2,
}
var yyPact = [...]int{

	86, -1000, 86, 75, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 19, 42, -22, -24, 57, 50, 48,
	73, -1000, 54, -1000, -1000, 4, -25, -18, 29, -26,
	-27, -28, -1000, -29, 19, -1000, -36, -30, -1000, -31,
	134, 33, -1000, -3, 33, -33, 133, 21, -1000, -18,
	-1000, -1000, -1000, -1000, -32, -1000, -1000, 71, 67, -1000,
	9, 158, 153, 149, -1000, -1000, -1000, -1000, 100, 9,
	-1000, -1000, -1000, -1000, 104, -1000, 124, -39, -1000, -12,
	-43, -13, 33, -1000, 128, -1000, -3, -3, 158, 9,
	9, 9, 9, 9, 9, 9, -15, 9, 9, 9,
	9, -1000, 122, -33, 9, 125, -1000, -1000, 102, -34,
	102, -35, -37, -1000, 14, -38, 67, -1000, 153, 153,
	153, 153, 153, 153, 153, -1000, 66, 149, 149, -1000,
	-1000, -1000, -1000, 153, -1000, -39, 32, 132, -1000, -1000,
	-1000, -1000, -1000, 32, 102, 102, 131, -1000, -17, -1000,
	32, -1000, -1000, 13, -46, -1000, -19, 32, 32, 32,
	90, -1000, -46, -1000, -1000, -1000, -1000, -1000, -1000, 107,
	32, 32, 131, 89, -1000, -1000, -1000, -1000, -46, -1000,
}
var yyPgo = [...]int{

	0, 12, 101, 202, 174, 201, 200, 199, 196, 195,
	194, 193, 192, 191, 190, 3, 189, 5, 188, 2,
	187, 1, 186, 185, 8, 184, 183, 6, 0, 7,
	13, 4, 182, 181, 10, 14, 16, 180, 9, 178,
	176, 79, 160, 150,
}
var yyR1 = [...]int{

	0, 43, 43, 3, 3, 3, 4, 4, 6, 6,
	6, 5, 5, 5, 5, 7, 14, 14, 15, 16,
	16, 24, 24, 24, 24, 24, 17, 17, 21, 22,
	22, 22, 8, 23, 23, 23, 23, 23, 23, 23,
	9, 10, 10, 42, 42, 41, 41, 41, 41, 41,
	11, 40, 40, 20, 20, 19, 18, 18, 28, 28,
	12, 12, 2, 2, 38, 38, 39, 13, 26, 25,
	25, 27, 37, 37, 36, 36, 35, 35, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 33, 31, 31,
	31, 30, 30, 30, 29, 29, 29, 29, 29, 29,
	1, 32, 32, 32,
}
var yyR2 = [...]int{

	0, 1, 0, 3, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 6, 3, 1, 1, 3,
	2, 4, 1, 1, 1, 1, 2, 1, 1, 2,
	2, 1, 4, 3, 4, 3, 5, 4, 4, 5,
	3, 6, 5, 3, 1, 1, 2, 1, 3, 2,
	8, 3, 1, 3, 1, 3, 3, 1, 1, 1,
	4, 3, 2, 1, 2, 0, 1, 4, 2, 3,
	1, 3, 3, 1, 3, 1, 2, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 3, 3, 3,
	1, 3, 3, 1, 1, 1, 1, 1, 2, 3,
	2, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -43, -3, -4, -5, -6, -10, -11, -12, -13,
	-7, -8, -9, 32, 30, 29, 35, 27, 37, 56,
	-4, 18, -42, -41, 6, 65, 31, 65, 65, 28,
	28, 28, 18, 34, 16, -2, -1, 60, 65, 17,
	65, -2, -38, 33, -26, 36, 65, 65, 65, 65,
	-41, -2, 65, 65, 14, -38, -39, -37, -36, -35,
	21, -34, -31, -30, -29, 63, -32, 66, 65, 14,
	58, 59, 48, -38, -25, -27, 65, 14, -23, 56,
	61, 37, -2, -38, -40, 65, 19, 20, -34, 8,
	9, 12, 11, 10, 13, 40, 39, 4, 5, 6,
	7, -1, -34, 16, 12, -14, -15, -16, 65, 62,
	65, 62, 62, -38, 15, 16, -36, -35, -31, -31,
	-31, -31, -31, -31, -31, -33, 63, -30, -30, -29,
	-29, 15, -27, -31, 15, 16, -24, 24, 22, 25,
	26, 23, 65, -24, 65, 65, 38, 65, 20, -15,
	-17, -21, -22, 21, 57, 51, 14, -17, -24, -24,
	-20, -19, 14, 63, -21, 48, -28, 66, 63, 63,
	-17, -17, 16, -18, -28, 15, -19, 15, 16, -28,
}
var yyDef = [...]int{

	2, -2, 1, 5, 6, 7, 11, 12, 13, 14,
	8, 9, 10, 0, 0, 0, 0, 0, 0, 0,
	0, 4, 0, 44, 45, 47, 0, 65, 0, 0,
	0, 0, 3, 0, 0, 46, 49, 0, 63, 0,
	0, 65, 61, 0, 65, 0, 0, 0, 40, 65,
	43, 48, 62, 100, 0, 60, 64, 66, 73, 75,
	0, 77, 86, 90, 93, 94, 95, 96, 97, 0,
	101, 102, 103, 67, 68, 70, 0, 0, 32, 0,
	0, 0, 65, 42, 0, 52, 0, 0, 76, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 98, 0, 0, 0, 0, 17, 18, 0, 0,
	0, 0, 0, 41, 0, 0, 72, 74, 78, 79,
	80, 81, 82, 83, 84, 85, 0, 88, 89, 91,
	92, 99, 69, 71, 15, 0, 20, 0, 22, 23,
	24, 25, 33, 35, 0, 0, 0, 51, 0, 16,
	19, 27, 28, 0, 0, 31, 0, 34, 37, 38,
	50, 54, 0, 87, 26, 29, 30, 58, 59, 0,
	36, 39, 0, 0, 57, 21, 53, 55, 0, 56,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:73
		{
			statements = yyDollar[1].stmt_list_t
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:74
		{
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:77
		{
			yyVAL.stmt_list_t = yyDollar[1].stmt_list_t
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[2].stmt_t)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:78
		{
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[1].stmt_t)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:79
		{
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[1].stmt_t)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:82
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:83
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:86
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:87
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:88
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:91
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:92
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:93
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:94
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:97
		{
			yyVAL.stmt_t = &createStatement{yyDollar[3].string_t, yyDollar[5].col_list_t}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:100
		{
			yyVAL.col_list_t = yyDollar[1].col_list_t
			yyVAL.col_list_t = append(yyVAL.col_list_t, yyDollar[3].col_t)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:101
		{
			yyVAL.col_list_t = append(yyVAL.col_list_t, yyDollar[1].col_t)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:104
		{
			yyVAL.col_t = yyDollar[1].col_t
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:107
		{
			yyVAL.col_t = &columnDefinition{yyDollar[1].string_t, yyDollar[2].data_t, yyDollar[3].obj_list_t}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:108
		{
			yyVAL.col_t = &columnDefinition{yyDollar[1].string_t, yyDollar[2].data_t, nil}
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:111
		{
			yyVAL.data_t = &charType{yyDollar[3].int64_t}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:112
		{
			yyVAL.data_t = &integerType{}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:113
		{
			yyVAL.data_t = &booleanType{}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:114
		{
			yyVAL.data_t = &datetimeType{}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:115
		{
			yyVAL.data_t = &floatType{}
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:118
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[2].obj_t)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:119
		{
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[1].obj_t)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:122
		{
			yyVAL.obj_t = yyDollar[1].obj_t
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:125
		{
			yyVAL.obj_t = &notNullConstraint{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:126
		{
			yyVAL.obj_t = &defaultConstraint{yyDollar[2].obj_t}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:127
		{
			yyVAL.obj_t = &autoincrementConstraint{}
		}
	case 32:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:130
		{
			yyVAL.stmt_t = &alterStatement{yyDollar[3].string_t, yyDollar[4].obj_t}
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:133
		{
			yyVAL.obj_t = &alterDrop{yyDollar[3].string_t}
		}
	case 34:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:134
		{
			yyVAL.obj_t = &alterAdd{yyDollar[2].string_t, yyDollar[3].data_t, yyDollar[4].obj_list_t}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:135
		{
			yyVAL.obj_t = &alterAdd{yyDollar[2].string_t, yyDollar[3].data_t, nil}
		}
	case 36:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:136
		{
			yyVAL.obj_t = &alterAdd{yyDollar[3].string_t, yyDollar[4].data_t, yyDollar[5].obj_list_t}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:137
		{
			yyVAL.obj_t = &alterAdd{yyDollar[3].string_t, yyDollar[4].data_t, nil}
		}
	case 38:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:138
		{
			yyVAL.obj_t = &alterModify{yyDollar[3].string_t, yyDollar[4].data_t, nil}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:139
		{
			yyVAL.obj_t = &alterModify{yyDollar[3].string_t, yyDollar[4].data_t, yyDollar[5].obj_list_t}
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:142
		{
			yyVAL.stmt_t = &dropStatement{yyDollar[3].string_t}
		}
	case 41:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:145
		{
			yyVAL.stmt_t = &selectStatement{yyDollar[4].string_t, yyDollar[5].string_t, yyDollar[2].columnSpec_list_t, yyDollar[6].expr_t}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:146
		{
			yyVAL.stmt_t = &selectStatement{yyDollar[4].string_t, "", yyDollar[2].columnSpec_list_t, yyDollar[5].expr_t}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:149
		{
			yyVAL.columnSpec_list_t = yyDollar[1].columnSpec_list_t
			yyVAL.columnSpec_list_t = append(yyVAL.columnSpec_list_t, yyDollar[3].columnSpec_t)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:150
		{
			yyVAL.columnSpec_list_t = append(yyVAL.columnSpec_list_t, yyDollar[1].columnSpec_t)
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:153
		{
			yyVAL.columnSpec_t = columnSpec{true, "", "", ""}
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:154
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, "", yyDollar[2].string_t}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:155
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, "", ""}
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:156
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, yyDollar[2].string_t, yyDollar[3].string_t}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:157
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, yyDollar[2].string_t, " "}
		}
	case 50:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:160
		{
			yyVAL.stmt_t = &insertStatement{yyDollar[3].string_t, yyDollar[5].string_list_t, yyDollar[8].obj_list_t}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:163
		{
			yyVAL.string_list_t = yyDollar[1].string_list_t
			yyVAL.string_list_t = append(yyVAL.string_list_t, yyDollar[3].string_t)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:164
		{
			yyVAL.string_list_t = append(yyVAL.string_list_t, yyDollar[1].string_t)
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:167
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[3].obj_list_t)
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:168
		{
			yyVAL.obj_list_t = append(yyDollar[1].obj_list_t)
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:171
		{
			yyVAL.obj_list_t = yyDollar[2].obj_list_t
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:174
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[3].obj_t)
		}
	case 57:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:175
		{
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[1].obj_t)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:178
		{
			yyVAL.obj_t = yyDollar[1].string_t
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:179
		{
			yyVAL.obj_t = yyDollar[1].int64_t
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:182
		{
			yyVAL.stmt_t = &deleteStatement{yyDollar[2].string_t, yyDollar[3].string_t, yyDollar[4].expr_t}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:183
		{
			yyVAL.stmt_t = &deleteStatement{yyDollar[2].string_t, "", yyDollar[3].expr_t}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:186
		{
			yyVAL.string_t = yyDollar[2].string_t
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:187
		{
			yyVAL.string_t = yyDollar[1].string_t
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:190
		{
			yyVAL.expr_t = yyDollar[2].expr_t
		}
	case 65:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:191
		{
			yyVAL.expr_t = nil
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:194
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 67:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:197
		{
			yyVAL.stmt_t = &updateStatement{yyDollar[2].string_t, yyDollar[3].assignments_list, yyDollar[4].expr_t}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:200
		{
			yyVAL.assignments_list = yyDollar[2].assignments_list
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:203
		{
			yyVAL.assignments_list = yyDollar[1].assignments_list
			yyVAL.assignments_list = append(yyDollar[1].assignments_list, yyDollar[3].assignment_t)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:204
		{
			yyVAL.assignments_list = append(yyVAL.assignments_list, yyDollar[1].assignment_t)
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:207
		{
			yyVAL.assignment_t = assignment{yyDollar[1].string_t, yyDollar[3].expr_t}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:210
		{
			yyVAL.expr_t = &orExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:211
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:214
		{
			yyVAL.expr_t = &andExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:215
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:218
		{
			yyVAL.expr_t = &notExpression{yyDollar[2].expr_t}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:219
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:222
		{
			yyVAL.expr_t = &ltExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:223
		{
			yyVAL.expr_t = &gtExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:224
		{
			yyVAL.expr_t = &eqExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:225
		{
			yyVAL.expr_t = &lteExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:226
		{
			yyVAL.expr_t = &gteExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:227
		{
			yyVAL.expr_t = &neExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:228
		{
			yyVAL.expr_t = &likeExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:229
		{
			yyVAL.expr_t = &betweenExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:230
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:233
		{
			yyVAL.expr_t = &betweenExpression{&intExpression{yyDollar[1].int64_t}, &intExpression{yyDollar[3].int64_t}}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:236
		{
			yyVAL.expr_t = &sumExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 89:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:237
		{
			yyVAL.expr_t = &subExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:238
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:241
		{
			yyVAL.expr_t = &multExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:242
		{
			yyVAL.expr_t = &divExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:243
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:246
		{
			yyVAL.expr_t = &intExpression{yyDollar[1].int64_t}
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:247
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:248
		{
			yyVAL.expr_t = &stringExpression{yyDollar[1].string_t}
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:249
		{
			yyVAL.expr_t = &idExpression{yyDollar[1].string_t, ""}
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:250
		{
			yyVAL.expr_t = &idExpression{yyDollar[1].string_t, yyDollar[2].string_t}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:251
		{
			yyVAL.expr_t = yyDollar[2].expr_t
		}
	case 100:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:254
		{
			yyVAL.string_t = yyDollar[2].string_t
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:257
		{
			yyVAL.expr_t = &trueExpression{}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:258
		{
			yyVAL.expr_t = &falseExpression{}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:259
		{
			yyVAL.expr_t = &nullExpression{}
		}
	}
	goto yystack /* stack new state and value */
}

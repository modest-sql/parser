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
	bool_t    bool

	expr_t expression

	stmt_list_t       statementList
	stmt_t            statement
	columnSpec_t      columnSpec
	columnSpec_list_t []columnSpec
	joinSpec_list_t   []joinSpec
	joinSpec_t        joinSpec
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
const KW_PRIMARY = 57369
const KW_FOREIGN = 57370
const KW_KEY = 57371
const KW_CREATE = 57372
const KW_TABLE = 57373
const KW_DELETE = 57374
const KW_INSERT = 57375
const KW_INTO = 57376
const KW_SELECT = 57377
const KW_WHERE = 57378
const KW_FROM = 57379
const KW_UPDATE = 57380
const KW_SET = 57381
const KW_ALTER = 57382
const KW_VALUES = 57383
const KW_BETWEEN = 57384
const KW_LIKE = 57385
const KW_INNER = 57386
const KW_HAVING = 57387
const KW_SUM = 57388
const KW_COUNT = 57389
const KW_AVG = 57390
const KW_MIN = 57391
const KW_MAX = 57392
const KW_NULL = 57393
const KW_IN = 57394
const KW_IS = 57395
const KW_AUTO_INCREMENT = 57396
const KW_JOIN = 57397
const KW_ON = 57398
const KW_GROUP = 57399
const KW_BY = 57400
const KW_DROP = 57401
const KW_DEFAULT = 57402
const KW_TRUE = 57403
const KW_FALSE = 57404
const KW_AS = 57405
const KW_ADD = 57406
const KW_COLUMN = 57407
const INT_LIT = 57408
const FLOAT_LIT = 57409
const TK_ID = 57410
const STR_LIT = 57411

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
	"KW_PRIMARY",
	"KW_FOREIGN",
	"KW_KEY",
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

//line parser.y:290

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

const yyLast = 229

var yyAct = [...]int{

	57, 171, 35, 67, 86, 170, 124, 150, 109, 87,
	63, 42, 118, 75, 64, 43, 70, 128, 183, 65,
	60, 36, 59, 61, 62, 70, 37, 71, 72, 196,
	41, 38, 126, 129, 24, 125, 37, 39, 111, 51,
	114, 38, 37, 113, 164, 159, 158, 38, 156, 76,
	85, 53, 82, 56, 52, 49, 73, 48, 47, 46,
	40, 83, 28, 71, 72, 27, 192, 186, 66, 140,
	69, 68, 71, 72, 81, 115, 162, 66, 112, 69,
	68, 198, 188, 37, 119, 163, 91, 54, 38, 45,
	43, 104, 127, 79, 23, 105, 25, 26, 80, 31,
	30, 29, 168, 132, 133, 134, 135, 136, 137, 138,
	191, 131, 130, 190, 55, 141, 142, 90, 147, 34,
	146, 157, 143, 144, 89, 32, 21, 39, 160, 50,
	161, 17, 165, 15, 14, 122, 13, 173, 106, 16,
	33, 18, 197, 177, 178, 92, 93, 96, 95, 94,
	97, 107, 145, 152, 155, 151, 153, 154, 169, 88,
	19, 166, 167, 180, 148, 149, 181, 182, 184, 179,
	175, 127, 187, 77, 185, 1, 174, 117, 127, 99,
	98, 189, 187, 120, 121, 116, 195, 193, 194, 92,
	93, 96, 95, 94, 97, 187, 187, 199, 22, 200,
	102, 103, 100, 101, 3, 84, 58, 20, 139, 44,
	74, 176, 78, 172, 123, 110, 108, 9, 8, 7,
	6, 12, 11, 99, 98, 10, 5, 4, 2,
}
var yyPact = [...]int{

	101, -1000, 101, 108, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 28, 63, -3, -6, 70, 69, 68,
	107, -1000, 103, -1000, -1000, 20, -8, -21, 50, -9,
	-10, -11, -1000, -13, 28, -1000, -37, -14, -1000, -17,
	73, 54, -1000, 2, 54, -19, 159, 34, -1000, -21,
	-1000, -1000, -1000, -1000, -18, 145, -1000, -1000, 105, 97,
	-1000, 11, 181, 198, 194, -1000, -1000, -1000, -1000, 110,
	11, -1000, -1000, -1000, 122, -1000, 139, -30, -1000, 13,
	-25, 10, 40, -1000, 168, -1000, 119, -1000, -34, 2,
	2, 181, 11, 11, 11, 11, 11, 11, 11, 3,
	11, 11, 11, 11, -1000, 137, -19, 11, 149, -1000,
	-1000, 131, -20, 131, -22, -23, 54, 40, -1000, 21,
	44, -24, 145, 146, -1000, -1000, -1000, -1000, -1000, -1000,
	97, -1000, 198, 198, 198, 198, 198, 198, 198, -1000,
	82, 194, 194, -1000, -1000, -1000, -1000, 198, -1000, -30,
	116, 155, -1000, -1000, -1000, -1000, -1000, 116, 131, 131,
	-1000, -1000, -50, 145, -1000, -1000, -1000, -34, 1, -1000,
	116, -1000, -1000, 31, -34, -1000, -1000, 84, 81, 0,
	116, 116, 116, -27, 119, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, 127, 116, 116, 25, 2, -1000, 2, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 21, 2, 3, 228, 204, 227, 226, 225, 222,
	221, 220, 219, 218, 217, 216, 8, 215, 5, 214,
	9, 4, 1, 213, 212, 211, 7, 210, 209, 13,
	6, 19, 14, 10, 208, 24, 20, 22, 206, 11,
	0, 205, 94, 198, 185, 177, 12, 175,
}
var yyR1 = [...]int{

	0, 47, 47, 4, 4, 4, 5, 5, 7, 7,
	7, 6, 6, 6, 6, 8, 15, 15, 16, 17,
	17, 26, 26, 26, 26, 26, 18, 18, 22, 23,
	23, 23, 23, 25, 25, 9, 24, 24, 24, 24,
	24, 24, 24, 10, 11, 11, 43, 43, 42, 42,
	42, 42, 42, 12, 12, 41, 41, 21, 21, 20,
	19, 19, 30, 30, 30, 30, 30, 44, 44, 45,
	45, 46, 46, 13, 13, 2, 2, 39, 39, 40,
	14, 28, 27, 27, 29, 38, 38, 37, 37, 36,
	36, 35, 35, 35, 35, 35, 35, 35, 35, 35,
	34, 33, 33, 33, 32, 32, 32, 31, 31, 31,
	31, 31, 31, 1, 3, 3,
}
var yyR2 = [...]int{

	0, 1, 0, 3, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 6, 3, 1, 1, 3,
	2, 4, 1, 1, 1, 1, 2, 1, 1, 2,
	2, 1, 1, 2, 2, 4, 3, 4, 3, 5,
	4, 4, 5, 3, 7, 5, 3, 1, 1, 2,
	1, 3, 2, 8, 5, 3, 1, 3, 1, 3,
	3, 1, 1, 1, 1, 1, 1, 1, 0, 2,
	1, 6, 5, 4, 3, 2, 1, 2, 0, 1,
	4, 2, 3, 1, 3, 3, 1, 3, 1, 2,
	1, 3, 3, 3, 3, 3, 3, 3, 3, 1,
	3, 3, 3, 1, 3, 3, 1, 1, 1, 1,
	1, 2, 3, 2, 1, 1,
}
var yyChk = [...]int{

	-1000, -47, -4, -5, -6, -7, -11, -12, -13, -14,
	-8, -9, -10, 35, 33, 32, 38, 30, 40, 59,
	-5, 18, -43, -42, 6, 68, 34, 68, 68, 31,
	31, 31, 18, 37, 16, -2, -1, 63, 68, 17,
	68, -2, -39, 36, -28, 39, 68, 68, 68, 68,
	-42, -2, 68, 68, 14, 41, -39, -40, -38, -37,
	-36, 21, -35, -33, -32, -31, 66, -3, 69, 68,
	14, 61, 62, -39, -27, -29, 68, 14, -24, 59,
	64, 40, -2, -39, -41, 68, -21, -20, 14, 19,
	20, -35, 8, 9, 12, 11, 10, 13, 43, 42,
	4, 5, 6, 7, -1, -35, 16, 12, -15, -16,
	-17, 68, 65, 68, 65, 65, -44, -45, -46, 44,
	15, 16, 16, -19, -30, 69, 66, -3, 51, 67,
	-37, -36, -33, -33, -33, -33, -33, -33, -33, -34,
	66, -32, -32, -31, -31, 15, -29, -33, 15, 16,
	-26, 24, 22, 25, 26, 23, 68, -26, 68, 68,
	-39, -46, 55, 41, 68, -20, 15, 16, 20, -16,
	-18, -22, -23, 21, 60, 54, -25, 27, 28, 14,
	-18, -26, -26, 68, -21, -30, 66, -22, 51, -30,
	29, 29, 66, -18, -18, -2, 56, 15, 56, -40,
	-40,
}
var yyDef = [...]int{

	2, -2, 1, 5, 6, 7, 11, 12, 13, 14,
	8, 9, 10, 0, 0, 0, 0, 0, 0, 0,
	0, 4, 0, 47, 48, 50, 0, 78, 0, 0,
	0, 0, 3, 0, 0, 49, 52, 0, 76, 0,
	0, 78, 74, 0, 78, 0, 0, 0, 43, 78,
	46, 51, 75, 113, 0, 0, 73, 77, 79, 86,
	88, 0, 90, 99, 103, 106, 107, 108, 109, 110,
	0, 114, 115, 80, 81, 83, 0, 0, 35, 0,
	0, 0, 68, 45, 0, 56, 54, 58, 0, 0,
	0, 89, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 111, 0, 0, 0, 0, 17,
	18, 0, 0, 0, 0, 0, 78, 67, 70, 0,
	0, 0, 0, 0, 61, 62, 63, 64, 65, 66,
	85, 87, 91, 92, 93, 94, 95, 96, 97, 98,
	0, 101, 102, 104, 105, 112, 82, 84, 15, 0,
	20, 0, 22, 23, 24, 25, 36, 38, 0, 0,
	44, 69, 0, 0, 55, 57, 59, 0, 0, 16,
	19, 27, 28, 0, 0, 31, 32, 0, 0, 0,
	37, 40, 41, 0, 53, 60, 100, 26, 29, 30,
	33, 34, 0, 39, 42, 0, 0, 21, 0, 72,
	71,
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
	62, 63, 64, 65, 66, 67, 68, 69,
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
		//line parser.y:81
		{
			statements = yyDollar[1].stmt_list_t
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:82
		{
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:85
		{
			yyVAL.stmt_list_t = yyDollar[1].stmt_list_t
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[2].stmt_t)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:86
		{
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[1].stmt_t)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:87
		{
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[1].stmt_t)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:90
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:91
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:94
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:95
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:96
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:99
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:100
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:101
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:102
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:105
		{
			yyVAL.stmt_t = &createStatement{yyDollar[3].string_t, yyDollar[5].col_list_t}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:108
		{
			yyVAL.col_list_t = yyDollar[1].col_list_t
			yyVAL.col_list_t = append(yyVAL.col_list_t, yyDollar[3].col_t)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:109
		{
			yyVAL.col_list_t = append(yyVAL.col_list_t, yyDollar[1].col_t)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:112
		{
			yyVAL.col_t = yyDollar[1].col_t
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:115
		{
			yyVAL.col_t = &columnDefinition{yyDollar[1].string_t, yyDollar[2].data_t, yyDollar[3].obj_list_t}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:116
		{
			yyVAL.col_t = &columnDefinition{yyDollar[1].string_t, yyDollar[2].data_t, nil}
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:119
		{
			yyVAL.data_t = &charType{yyDollar[3].int64_t}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:120
		{
			yyVAL.data_t = &integerType{}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:121
		{
			yyVAL.data_t = &booleanType{}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:122
		{
			yyVAL.data_t = &datetimeType{}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:123
		{
			yyVAL.data_t = &floatType{}
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:126
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[2].obj_t)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:127
		{
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[1].obj_t)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:130
		{
			yyVAL.obj_t = yyDollar[1].obj_t
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:133
		{
			yyVAL.obj_t = &notNullConstraint{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:134
		{
			yyVAL.obj_t = &defaultConstraint{yyDollar[2].obj_t}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:135
		{
			yyVAL.obj_t = &autoincrementConstraint{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:136
		{
			yyVAL.obj_t = yyDollar[1].obj_t
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:139
		{
			yyVAL.obj_t = &primaryKeyConstraint{}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:140
		{
			yyVAL.obj_t = &foreignKeyConstraint{}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:143
		{
			yyVAL.stmt_t = &alterStatement{yyDollar[3].string_t, yyDollar[4].obj_t}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:146
		{
			yyVAL.obj_t = &alterDrop{yyDollar[3].string_t}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:147
		{
			yyVAL.obj_t = &alterAdd{yyDollar[2].string_t, yyDollar[3].data_t, yyDollar[4].obj_list_t}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:148
		{
			yyVAL.obj_t = &alterAdd{yyDollar[2].string_t, yyDollar[3].data_t, nil}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:149
		{
			yyVAL.obj_t = &alterAdd{yyDollar[3].string_t, yyDollar[4].data_t, yyDollar[5].obj_list_t}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:150
		{
			yyVAL.obj_t = &alterAdd{yyDollar[3].string_t, yyDollar[4].data_t, nil}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:151
		{
			yyVAL.obj_t = &alterModify{yyDollar[3].string_t, yyDollar[4].data_t, nil}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:152
		{
			yyVAL.obj_t = &alterModify{yyDollar[3].string_t, yyDollar[4].data_t, yyDollar[5].obj_list_t}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:155
		{
			yyVAL.stmt_t = &dropStatement{yyDollar[3].string_t}
		}
	case 44:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:158
		{
			yyVAL.stmt_t = &selectStatement{yyDollar[2].columnSpec_list_t, yyDollar[4].string_t, yyDollar[5].string_t, yyDollar[6].joinSpec_list_t, yyDollar[7].expr_t}
		}
	case 45:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:159
		{
			yyVAL.stmt_t = &selectStatement{yyDollar[2].columnSpec_list_t, yyDollar[4].string_t, "", nil, yyDollar[5].expr_t}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:162
		{
			yyVAL.columnSpec_list_t = yyDollar[1].columnSpec_list_t
			yyVAL.columnSpec_list_t = append(yyVAL.columnSpec_list_t, yyDollar[3].columnSpec_t)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:163
		{
			yyVAL.columnSpec_list_t = append(yyVAL.columnSpec_list_t, yyDollar[1].columnSpec_t)
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:166
		{
			yyVAL.columnSpec_t = columnSpec{true, "", "", ""}
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:167
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, "", yyDollar[2].string_t}
		}
	case 50:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:168
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, "", ""}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:169
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, yyDollar[2].string_t, yyDollar[3].string_t}
		}
	case 52:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:170
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, yyDollar[2].string_t, " "}
		}
	case 53:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:173
		{
			yyVAL.stmt_t = &insertStatement{yyDollar[3].string_t, yyDollar[5].string_list_t, yyDollar[8].obj_list_t}
		}
	case 54:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:174
		{
			yyVAL.stmt_t = &insertStatement{yyDollar[3].string_t, nil, yyDollar[5].obj_list_t}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:177
		{
			yyVAL.string_list_t = yyDollar[1].string_list_t
			yyVAL.string_list_t = append(yyVAL.string_list_t, yyDollar[3].string_t)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:178
		{
			yyVAL.string_list_t = append(yyVAL.string_list_t, yyDollar[1].string_t)
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:181
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[3].obj_list_t)
		}
	case 58:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:182
		{
			yyVAL.obj_list_t = append(yyDollar[1].obj_list_t)
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:185
		{
			yyVAL.obj_list_t = yyDollar[2].obj_list_t
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:188
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[3].obj_t)
		}
	case 61:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:189
		{
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[1].obj_t)
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:192
		{
			yyVAL.obj_t = yyDollar[1].string_t
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:193
		{
			yyVAL.obj_t = yyDollar[1].int64_t
		}
	case 64:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:194
		{
			yyVAL.obj_t = yyDollar[1].bool_t
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:195
		{
			yyVAL.obj_t = nil
		}
	case 66:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:196
		{
			yyVAL.obj_t = yyDollar[1].float64_t
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:199
		{
			yyVAL.joinSpec_list_t = yyDollar[1].joinSpec_list_t
		}
	case 68:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:200
		{
			yyVAL.joinSpec_list_t = nil
		}
	case 69:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:203
		{
			yyVAL.joinSpec_list_t = yyDollar[1].joinSpec_list_t
			yyVAL.joinSpec_list_t = append(yyVAL.joinSpec_list_t, yyDollar[2].joinSpec_t)
		}
	case 70:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:204
		{
			yyVAL.joinSpec_list_t = append(yyVAL.joinSpec_list_t, yyDollar[1].joinSpec_t)
		}
	case 71:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:207
		{
			yyVAL.joinSpec_t = joinSpec{yyDollar[3].string_t, yyDollar[4].string_t, yyDollar[6].expr_t}
		}
	case 72:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:208
		{
			yyVAL.joinSpec_t = joinSpec{yyDollar[3].string_t, "", yyDollar[5].expr_t}
		}
	case 73:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:211
		{
			yyVAL.stmt_t = &deleteStatement{yyDollar[2].string_t, yyDollar[3].string_t, yyDollar[4].expr_t}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:212
		{
			yyVAL.stmt_t = &deleteStatement{yyDollar[2].string_t, "", yyDollar[3].expr_t}
		}
	case 75:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:215
		{
			yyVAL.string_t = yyDollar[2].string_t
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:216
		{
			yyVAL.string_t = yyDollar[1].string_t
		}
	case 77:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:219
		{
			yyVAL.expr_t = yyDollar[2].expr_t
		}
	case 78:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:220
		{
			yyVAL.expr_t = nil
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:223
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 80:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:226
		{
			yyVAL.stmt_t = &updateStatement{yyDollar[2].string_t, yyDollar[3].assignments_list, yyDollar[4].expr_t}
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:229
		{
			yyVAL.assignments_list = yyDollar[2].assignments_list
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:232
		{
			yyVAL.assignments_list = yyDollar[1].assignments_list
			yyVAL.assignments_list = append(yyDollar[1].assignments_list, yyDollar[3].assignment_t)
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:233
		{
			yyVAL.assignments_list = append(yyVAL.assignments_list, yyDollar[1].assignment_t)
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:236
		{
			yyVAL.assignment_t = assignment{yyDollar[1].string_t, yyDollar[3].expr_t}
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:239
		{
			yyVAL.expr_t = &orExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:240
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 87:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:243
		{
			yyVAL.expr_t = &andExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:244
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:247
		{
			yyVAL.expr_t = &notExpression{yyDollar[2].expr_t}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:248
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:251
		{
			yyVAL.expr_t = &ltExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:252
		{
			yyVAL.expr_t = &gtExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 93:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:253
		{
			yyVAL.expr_t = &eqExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:254
		{
			yyVAL.expr_t = &lteExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:255
		{
			yyVAL.expr_t = &gteExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:256
		{
			yyVAL.expr_t = &neExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:257
		{
			yyVAL.expr_t = &likeExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:258
		{
			yyVAL.expr_t = &betweenExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:259
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 100:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:262
		{
			yyVAL.expr_t = &betweenExpression{&intExpression{yyDollar[1].int64_t}, &intExpression{yyDollar[3].int64_t}}
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:265
		{
			yyVAL.expr_t = &sumExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:266
		{
			yyVAL.expr_t = &subExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:267
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:270
		{
			yyVAL.expr_t = &multExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:271
		{
			yyVAL.expr_t = &divExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:272
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:275
		{
			yyVAL.expr_t = &intExpression{yyDollar[1].int64_t}
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:276
		{
			yyVAL.expr_t = &boolExpression{yyDollar[1].bool_t}
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:277
		{
			yyVAL.expr_t = &stringExpression{yyDollar[1].string_t}
		}
	case 110:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:278
		{
			yyVAL.expr_t = &idExpression{yyDollar[1].string_t, ""}
		}
	case 111:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:279
		{
			yyVAL.expr_t = &idExpression{yyDollar[1].string_t, yyDollar[2].string_t}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:280
		{
			yyVAL.expr_t = yyDollar[2].expr_t
		}
	case 113:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:283
		{
			yyVAL.string_t = yyDollar[2].string_t
		}
	case 114:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:286
		{
			yyVAL.bool_t = true
		}
	case 115:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:287
		{
			yyVAL.bool_t = false
		}
	}
	goto yystack /* stack new state and value */
}

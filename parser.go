//line parser.y:4
package parser

import __yyfmt__ "fmt"

//line parser.y:4
import (
	"io"
	"sync"
)

var statements statementList
var lock sync.Mutex

//line parser.y:16
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
	joinSpec_list_t   []joinSpec
	joinSpec_t        joinSpec
	col_list_t        columnDefinitions
	col_t             *columnDefinition
	assignment_t      assignment
	data_t            dataType
	assignments_list  []assignment
	obj_list_t        []interface{}
	string_list_t     []string
	group_by_t        GroupBySpec
	group_by_list_t   []GroupBySpec
	obj_t             interface{}
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
const KW_BY = 57372
const KW_CREATE = 57373
const KW_TABLE = 57374
const KW_DELETE = 57375
const KW_INSERT = 57376
const KW_INTO = 57377
const KW_SELECT = 57378
const KW_WHERE = 57379
const KW_FROM = 57380
const KW_UPDATE = 57381
const KW_SET = 57382
const KW_ALTER = 57383
const KW_VALUES = 57384
const KW_BETWEEN = 57385
const KW_LIKE = 57386
const KW_INNER = 57387
const KW_HAVING = 57388
const KW_SUM = 57389
const KW_COUNT = 57390
const KW_AVG = 57391
const KW_MIN = 57392
const KW_MAX = 57393
const KW_NULL = 57394
const KW_IN = 57395
const KW_IS = 57396
const KW_AUTO_INCREMENT = 57397
const KW_JOIN = 57398
const KW_ON = 57399
const KW_GROUP = 57400
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
	"KW_BY",
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

//line parser.y:311

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
		lock.Unlock()
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

const yyLast = 276

var yyAct = [...]int{

	99, 206, 41, 40, 205, 109, 142, 95, 131, 168,
	130, 119, 105, 107, 101, 96, 128, 106, 70, 74,
	234, 102, 104, 71, 115, 112, 42, 124, 221, 219,
	123, 43, 103, 113, 114, 44, 121, 112, 144, 146,
	236, 143, 42, 184, 177, 59, 176, 43, 174, 42,
	75, 94, 66, 24, 43, 69, 65, 64, 63, 62,
	61, 81, 60, 115, 57, 84, 86, 88, 90, 92,
	56, 72, 113, 114, 55, 115, 82, 108, 54, 111,
	110, 42, 51, 50, 113, 114, 43, 33, 98, 108,
	230, 111, 110, 198, 26, 27, 28, 29, 30, 80,
	125, 122, 132, 145, 208, 238, 180, 226, 129, 183,
	212, 213, 53, 67, 162, 25, 32, 78, 23, 71,
	39, 31, 79, 36, 35, 34, 149, 182, 224, 17,
	165, 15, 14, 175, 13, 163, 164, 16, 210, 18,
	229, 68, 38, 209, 179, 178, 150, 151, 154, 153,
	152, 155, 228, 203, 148, 147, 185, 19, 58, 91,
	37, 44, 188, 190, 191, 192, 193, 194, 195, 196,
	189, 89, 21, 44, 201, 202, 199, 200, 44, 204,
	215, 157, 156, 87, 237, 44, 216, 217, 140, 218,
	220, 222, 181, 145, 223, 150, 151, 154, 153, 152,
	155, 170, 173, 169, 171, 172, 85, 225, 44, 83,
	137, 44, 186, 187, 116, 145, 227, 225, 166, 167,
	136, 231, 232, 233, 235, 138, 139, 135, 134, 133,
	157, 156, 97, 225, 225, 239, 214, 76, 49, 240,
	48, 47, 46, 45, 117, 160, 161, 158, 159, 3,
	1, 127, 20, 126, 22, 93, 100, 197, 52, 73,
	211, 77, 207, 141, 120, 118, 9, 8, 7, 6,
	12, 11, 10, 5, 4, 2,
}
var yyPact = [...]int{

	98, -1000, 98, 154, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 47, 86, 78, 19, 93, 92, 91,
	142, -1000, 104, -1000, -1000, 18, 229, 228, 227, 226,
	224, 15, 14, 72, 10, 6, 2, -1000, -4, 47,
	-1000, -21, -6, -1000, -8, -9, -10, -11, -12, -16,
	99, -14, 82, -18, 223, 58, -1000, -14, -1000, -1000,
	-1000, -1000, 194, 191, 168, 156, 144, -17, 218, 82,
	-1000, 11, -1000, 198, -1000, 232, -32, -1000, 36, -38,
	35, 63, 44, -1000, 214, -1000, 213, -1000, 212, -1000,
	205, -1000, 195, 210, -1000, 172, -1000, -28, -1000, -1000,
	136, 134, -1000, 23, 187, 243, 239, -1000, -1000, -1000,
	-1000, 161, 23, -1000, -1000, -1000, -18, 23, 203, -1000,
	-1000, 179, -20, 179, -22, -24, 82, 63, -1000, 50,
	176, -1000, 97, -1000, -1000, -1000, -1000, -1000, 67, -25,
	218, 197, -1000, -1000, -1000, -1000, -1000, 11, 11, 187,
	23, 23, 23, 23, 23, 23, 23, 27, 23, 23,
	23, 23, -1000, 138, -1000, 243, -1000, -32, 83, 222,
	-1000, -1000, -1000, -1000, -1000, 83, 179, 179, 44, -1000,
	-39, 44, -40, 218, -1000, -1000, -1000, -28, 134, -1000,
	243, 243, 243, 243, 243, 243, 243, -1000, 108, 239,
	239, -1000, -1000, -1000, -1000, 83, -1000, -1000, 55, -28,
	-1000, -1000, 123, 111, 24, 83, 83, 83, 176, -37,
	-1000, 161, 172, -1000, -26, -1000, -1000, -1000, -1000, -1000,
	169, 83, 83, 48, 11, -1000, -1000, -1000, 11, -1000,
	-1000,
}
var yyPgo = [...]int{

	0, 2, 3, 275, 249, 274, 273, 272, 271, 270,
	269, 268, 267, 266, 265, 11, 264, 4, 263, 15,
	7, 1, 262, 261, 260, 9, 259, 258, 19, 6,
	13, 17, 12, 5, 257, 22, 21, 14, 256, 18,
	0, 255, 118, 254, 253, 251, 16, 8, 10, 250,
}
var yyR1 = [...]int{

	0, 49, 49, 3, 3, 3, 4, 4, 6, 6,
	6, 5, 5, 5, 5, 7, 14, 14, 15, 16,
	16, 25, 25, 25, 25, 25, 17, 17, 21, 22,
	22, 22, 22, 24, 24, 8, 23, 23, 23, 23,
	23, 23, 23, 9, 10, 10, 48, 48, 48, 47,
	47, 43, 43, 42, 42, 42, 42, 42, 42, 42,
	42, 42, 42, 42, 42, 42, 42, 42, 11, 11,
	41, 41, 20, 20, 19, 18, 18, 29, 29, 29,
	29, 44, 44, 45, 45, 46, 46, 12, 12, 2,
	2, 39, 39, 40, 13, 27, 26, 26, 28, 38,
	38, 37, 37, 36, 36, 35, 35, 35, 35, 35,
	35, 35, 35, 35, 34, 32, 32, 32, 31, 31,
	31, 30, 30, 30, 30, 30, 30, 1, 33, 33,
	33,
}
var yyR2 = [...]int{

	0, 1, 0, 3, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 6, 3, 1, 1, 3,
	2, 4, 1, 1, 1, 1, 2, 1, 1, 2,
	2, 1, 1, 2, 2, 4, 3, 4, 3, 5,
	4, 4, 5, 3, 8, 6, 3, 1, 0, 3,
	4, 3, 1, 1, 2, 1, 3, 2, 4, 5,
	4, 5, 4, 5, 4, 5, 4, 5, 8, 5,
	3, 1, 3, 1, 3, 3, 1, 1, 1, 1,
	1, 1, 0, 2, 1, 6, 5, 5, 4, 2,
	1, 2, 0, 1, 4, 2, 3, 1, 3, 3,
	1, 3, 1, 2, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 1, 3, 3, 3, 1, 3, 3,
	1, 1, 1, 1, 1, 2, 3, 2, 1, 1,
	1,
}
var yyChk = [...]int{

	-1000, -49, -3, -4, -5, -6, -10, -11, -12, -13,
	-7, -8, -9, 36, 34, 33, 39, 31, 41, 59,
	-4, 18, -43, -42, 6, 68, 47, 48, 49, 50,
	51, 35, 38, 68, 32, 32, 32, 18, 38, 16,
	-2, -1, 63, 68, 17, 14, 14, 14, 14, 14,
	68, 68, -27, 40, 68, 68, 68, 68, -42, -2,
	68, 68, 68, 68, 68, 68, 68, 14, 42, -2,
	-39, 37, -39, -26, -28, 68, 14, -23, 59, 64,
	41, -2, -39, 15, -1, 15, -1, 15, -1, 15,
	-1, 15, -1, -41, 68, -20, -19, 14, -39, -40,
	-38, -37, -36, 21, -35, -32, -31, -30, 66, -33,
	69, 68, 14, 61, 62, 52, 16, 12, -14, -15,
	-16, 68, 65, 68, 65, 65, -44, -45, -46, 45,
	-48, -47, 58, 15, 15, 15, 15, 15, 15, 16,
	16, -18, -29, 69, 66, -33, 67, 19, 20, -35,
	8, 9, 12, 11, 10, 13, 44, 43, 4, 5,
	6, 7, -1, -35, -28, -32, 15, 16, -25, 24,
	22, 25, 26, 23, 68, -25, 68, 68, -39, -46,
	56, 16, 30, 42, 68, -19, 15, 16, -37, -36,
	-32, -32, -32, -32, -32, -32, -32, -34, 66, -31,
	-31, -30, -30, 15, -15, -17, -21, -22, 21, 60,
	55, -24, 27, 28, 14, -17, -25, -25, -48, 68,
	-47, 68, -20, -29, 20, -21, 52, -29, 29, 29,
	66, -17, -17, -2, 57, -1, 66, 15, 57, -40,
	-40,
}
var yyDef = [...]int{

	2, -2, 1, 5, 6, 7, 11, 12, 13, 14,
	8, 9, 10, 0, 0, 0, 0, 0, 0, 0,
	0, 4, 0, 52, 53, 55, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 3, 0, 0,
	54, 57, 0, 90, 0, 0, 0, 0, 0, 0,
	0, 92, 92, 0, 0, 0, 43, 92, 51, 56,
	89, 127, 0, 0, 0, 0, 0, 0, 0, 92,
	88, 0, 94, 95, 97, 0, 0, 35, 0, 0,
	0, 82, 48, 58, 0, 60, 0, 62, 0, 64,
	0, 66, 0, 0, 71, 69, 73, 0, 87, 91,
	93, 100, 102, 0, 104, 113, 117, 120, 121, 122,
	123, 124, 0, 128, 129, 130, 0, 0, 0, 17,
	18, 0, 0, 0, 0, 0, 92, 81, 84, 0,
	45, 47, 0, 59, 61, 63, 65, 67, 0, 0,
	0, 0, 76, 77, 78, 79, 80, 0, 0, 103,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 125, 0, 96, 98, 15, 0, 20, 0,
	22, 23, 24, 25, 36, 38, 0, 0, 48, 83,
	0, 0, 0, 0, 70, 72, 74, 0, 99, 101,
	105, 106, 107, 108, 109, 110, 111, 112, 0, 115,
	116, 118, 119, 126, 16, 19, 27, 28, 0, 0,
	31, 32, 0, 0, 0, 37, 40, 41, 44, 0,
	46, 49, 68, 75, 0, 26, 29, 30, 33, 34,
	0, 39, 42, 0, 0, 50, 114, 21, 0, 86,
	85,
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
		//line parser.y:83
		{
			lock.Lock()
			statements = yyDollar[1].stmt_list_t
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:84
		{
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:87
		{
			yyVAL.stmt_list_t = yyDollar[1].stmt_list_t
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[2].stmt_t)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:88
		{
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[1].stmt_t)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:89
		{
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[1].stmt_t)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:92
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:93
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:96
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:97
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:98
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:101
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:102
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:103
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:104
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:107
		{
			yyVAL.stmt_t = &createStatement{yyDollar[3].string_t, yyDollar[5].col_list_t}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:110
		{
			yyVAL.col_list_t = yyDollar[1].col_list_t
			yyVAL.col_list_t = append(yyVAL.col_list_t, yyDollar[3].col_t)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:111
		{
			yyVAL.col_list_t = append(yyVAL.col_list_t, yyDollar[1].col_t)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:114
		{
			yyVAL.col_t = yyDollar[1].col_t
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:117
		{
			yyVAL.col_t = &columnDefinition{yyDollar[1].string_t, yyDollar[2].data_t, yyDollar[3].obj_list_t}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:118
		{
			yyVAL.col_t = &columnDefinition{yyDollar[1].string_t, yyDollar[2].data_t, nil}
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:121
		{
			yyVAL.data_t = &charType{yyDollar[3].int64_t}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:122
		{
			yyVAL.data_t = &integerType{}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:123
		{
			yyVAL.data_t = &booleanType{}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:124
		{
			yyVAL.data_t = &datetimeType{}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:125
		{
			yyVAL.data_t = &floatType{}
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:128
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[2].obj_t)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:129
		{
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[1].obj_t)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:132
		{
			yyVAL.obj_t = yyDollar[1].obj_t
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:135
		{
			yyVAL.obj_t = &notNullConstraint{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:136
		{
			yyVAL.obj_t = &defaultConstraint{yyDollar[2].obj_t}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:137
		{
			yyVAL.obj_t = &autoincrementConstraint{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:138
		{
			yyVAL.obj_t = yyDollar[1].obj_t
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:141
		{
			yyVAL.obj_t = &primaryKeyConstraint{}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:142
		{
			yyVAL.obj_t = &foreignKeyConstraint{}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:145
		{
			yyVAL.stmt_t = &alterStatement{yyDollar[3].string_t, yyDollar[4].obj_t}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:148
		{
			yyVAL.obj_t = &alterDrop{yyDollar[3].string_t}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:149
		{
			yyVAL.obj_t = &alterAdd{yyDollar[2].string_t, yyDollar[3].data_t, yyDollar[4].obj_list_t}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:150
		{
			yyVAL.obj_t = &alterAdd{yyDollar[2].string_t, yyDollar[3].data_t, nil}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:151
		{
			yyVAL.obj_t = &alterAdd{yyDollar[3].string_t, yyDollar[4].data_t, yyDollar[5].obj_list_t}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:152
		{
			yyVAL.obj_t = &alterAdd{yyDollar[3].string_t, yyDollar[4].data_t, nil}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:153
		{
			yyVAL.obj_t = &alterModify{yyDollar[3].string_t, yyDollar[4].data_t, nil}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:154
		{
			yyVAL.obj_t = &alterModify{yyDollar[3].string_t, yyDollar[4].data_t, yyDollar[5].obj_list_t}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:157
		{
			yyVAL.stmt_t = &dropStatement{yyDollar[3].string_t}
		}
	case 44:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:160
		{
			yyVAL.stmt_t = &selectStatement{yyDollar[2].columnSpec_list_t, yyDollar[4].string_t, yyDollar[5].string_t, yyDollar[6].joinSpec_list_t, yyDollar[7].expr_t, yyDollar[8].group_by_list_t}
		}
	case 45:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:161
		{
			yyVAL.stmt_t = &selectStatement{yyDollar[2].columnSpec_list_t, yyDollar[4].string_t, "", nil, yyDollar[5].expr_t, yyDollar[6].group_by_list_t}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:164
		{
			yyVAL.group_by_list_t = yyDollar[1].group_by_list_t
			yyVAL.group_by_list_t = append(yyVAL.group_by_list_t, yyDollar[3].group_by_t)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:165
		{
			yyVAL.group_by_list_t = append(yyVAL.group_by_list_t, yyDollar[1].group_by_t)
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:166
		{
			yyVAL.group_by_list_t = nil
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:169
		{
			yyVAL.group_by_t = GroupBySpec{"", yyDollar[3].string_t}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:170
		{
			yyVAL.group_by_t = GroupBySpec{yyDollar[3].string_t, yyDollar[4].string_t}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:174
		{
			yyVAL.columnSpec_list_t = yyDollar[1].columnSpec_list_t
			yyVAL.columnSpec_list_t = append(yyVAL.columnSpec_list_t, yyDollar[3].columnSpec_t)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:175
		{
			yyVAL.columnSpec_list_t = append(yyVAL.columnSpec_list_t, yyDollar[1].columnSpec_t)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:178
		{
			yyVAL.columnSpec_t = columnSpec{true, "", "", "", nil}
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:179
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, "", yyDollar[2].string_t, nil}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:180
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, "", "", nil}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:181
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, yyDollar[2].string_t, yyDollar[3].string_t, nil}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:182
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, yyDollar[2].string_t, "", nil}
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:183
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, "", "", &functionSum{}}
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:184
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, yyDollar[4].string_t, "", &functionSum{}}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:185
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, "", "", &functionCount{}}
		}
	case 61:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:186
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, yyDollar[4].string_t, "", &functionCount{}}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:187
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, "", "", &functionAvg{}}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:188
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, yyDollar[4].string_t, "", &functionAvg{}}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:189
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, "", "", &functionMin{}}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:190
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, yyDollar[4].string_t, "", &functionMin{}}
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:191
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, "", "", &functionMax{}}
		}
	case 67:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:192
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, yyDollar[4].string_t, "", &functionMax{}}
		}
	case 68:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:194
		{
			yyVAL.stmt_t = &insertStatement{yyDollar[3].string_t, yyDollar[5].string_list_t, yyDollar[8].obj_list_t}
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:195
		{
			yyVAL.stmt_t = &insertStatement{yyDollar[3].string_t, nil, yyDollar[5].obj_list_t}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:198
		{
			yyVAL.string_list_t = yyDollar[1].string_list_t
			yyVAL.string_list_t = append(yyVAL.string_list_t, yyDollar[3].string_t)
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:199
		{
			yyVAL.string_list_t = append(yyVAL.string_list_t, yyDollar[1].string_t)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:202
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[3].obj_list_t)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:203
		{
			yyVAL.obj_list_t = append(yyDollar[1].obj_list_t)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:206
		{
			yyVAL.obj_list_t = yyDollar[2].obj_list_t
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:209
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[3].obj_t)
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:210
		{
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[1].obj_t)
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:213
		{
			yyVAL.obj_t = yyDollar[1].string_t
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:214
		{
			yyVAL.obj_t = yyDollar[1].int64_t
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:215
		{
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:216
		{
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:219
		{
			yyVAL.joinSpec_list_t = yyDollar[1].joinSpec_list_t
		}
	case 82:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:220
		{
			yyVAL.joinSpec_list_t = nil
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:223
		{
			yyVAL.joinSpec_list_t = yyDollar[1].joinSpec_list_t
			yyVAL.joinSpec_list_t = append(yyVAL.joinSpec_list_t, yyDollar[2].joinSpec_t)
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:224
		{
			yyVAL.joinSpec_list_t = append(yyVAL.joinSpec_list_t, yyDollar[1].joinSpec_t)
		}
	case 85:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:227
		{
			yyVAL.joinSpec_t = joinSpec{yyDollar[3].string_t, yyDollar[4].string_t, yyDollar[6].expr_t}
		}
	case 86:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:228
		{
			yyVAL.joinSpec_t = joinSpec{yyDollar[3].string_t, "", yyDollar[5].expr_t}
		}
	case 87:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:231
		{
			yyVAL.stmt_t = &deleteStatement{yyDollar[3].string_t, yyDollar[4].string_t, yyDollar[5].expr_t}
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:232
		{
			yyVAL.stmt_t = &deleteStatement{yyDollar[3].string_t, "", yyDollar[4].expr_t}
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:235
		{
			yyVAL.string_t = yyDollar[2].string_t
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:236
		{
			yyVAL.string_t = yyDollar[1].string_t
		}
	case 91:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:239
		{
			yyVAL.expr_t = yyDollar[2].expr_t
		}
	case 92:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:240
		{
			yyVAL.expr_t = nil
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:243
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 94:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:246
		{
			yyVAL.stmt_t = &updateStatement{yyDollar[2].string_t, yyDollar[3].assignments_list, yyDollar[4].expr_t}
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:249
		{
			yyVAL.assignments_list = yyDollar[2].assignments_list
		}
	case 96:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:252
		{
			yyVAL.assignments_list = yyDollar[1].assignments_list
			yyVAL.assignments_list = append(yyDollar[1].assignments_list, yyDollar[3].assignment_t)
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:253
		{
			yyVAL.assignments_list = append(yyVAL.assignments_list, yyDollar[1].assignment_t)
		}
	case 98:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:256
		{
			yyVAL.assignment_t = assignment{yyDollar[1].string_t, yyDollar[3].expr_t}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:259
		{
			yyVAL.expr_t = &orExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:260
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:263
		{
			yyVAL.expr_t = &andExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:264
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 103:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:267
		{
			yyVAL.expr_t = &notExpression{yyDollar[2].expr_t}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:268
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:271
		{
			yyVAL.expr_t = &ltExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 106:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:272
		{
			yyVAL.expr_t = &gtExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:273
		{
			yyVAL.expr_t = &eqExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:274
		{
			yyVAL.expr_t = &lteExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:275
		{
			yyVAL.expr_t = &gteExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:276
		{
			yyVAL.expr_t = &neExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:277
		{
			yyVAL.expr_t = &likeExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:278
		{
			yyVAL.expr_t = &betweenExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 113:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:279
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:282
		{
			yyVAL.expr_t = &betweenExpression{&intExpression{yyDollar[1].int64_t}, &intExpression{yyDollar[3].int64_t}}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:285
		{
			yyVAL.expr_t = &sumExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 116:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:286
		{
			yyVAL.expr_t = &subExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 117:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:287
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:290
		{
			yyVAL.expr_t = &multExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:291
		{
			yyVAL.expr_t = &divExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:292
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 121:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:295
		{
			yyVAL.expr_t = &intExpression{yyDollar[1].int64_t}
		}
	case 122:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:296
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:297
		{
			yyVAL.expr_t = &stringExpression{yyDollar[1].string_t}
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:298
		{
			yyVAL.expr_t = &idExpression{yyDollar[1].string_t, ""}
		}
	case 125:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:299
		{
			yyVAL.expr_t = &idExpression{yyDollar[1].string_t, yyDollar[2].string_t}
		}
	case 126:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:300
		{
			yyVAL.expr_t = yyDollar[2].expr_t
		}
	case 127:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:303
		{
			yyVAL.string_t = yyDollar[2].string_t
		}
	case 128:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:306
		{
			yyVAL.expr_t = &trueExpression{}
		}
	case 129:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:307
		{
			yyVAL.expr_t = &falseExpression{}
		}
	case 130:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:308
		{
			yyVAL.expr_t = &nullExpression{}
		}
	}
	goto yystack /* stack new state and value */
}

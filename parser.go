//line parser.y:3
package parser

import __yyfmt__ "fmt"

//line parser.y:5
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

//line parser.y:316

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
		} else {
			lock.Unlock()
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

const yyLast = 282

var yyAct = [...]int{

	42, 74, 210, 200, 84, 162, 211, 118, 188, 199,
	119, 142, 80, 82, 94, 104, 102, 81, 41, 53,
	79, 77, 105, 76, 167, 227, 54, 240, 33, 147,
	87, 43, 146, 88, 89, 45, 44, 78, 164, 166,
	43, 163, 87, 43, 144, 44, 204, 202, 44, 197,
	196, 52, 43, 194, 95, 117, 70, 44, 32, 24,
	69, 63, 68, 67, 66, 65, 64, 107, 109, 111,
	113, 115, 73, 61, 90, 91, 92, 88, 89, 60,
	101, 43, 83, 59, 86, 85, 44, 136, 58, 88,
	89, 55, 51, 34, 83, 236, 86, 85, 230, 123,
	26, 27, 28, 29, 30, 178, 100, 148, 137, 145,
	138, 152, 201, 241, 232, 105, 203, 57, 149, 151,
	54, 25, 150, 31, 98, 165, 213, 23, 37, 99,
	40, 36, 217, 218, 71, 35, 225, 170, 171, 172,
	173, 174, 175, 176, 169, 168, 235, 234, 181, 182,
	179, 180, 39, 185, 184, 195, 17, 208, 15, 14,
	215, 13, 72, 122, 16, 214, 18, 121, 62, 198,
	114, 205, 45, 45, 124, 125, 128, 127, 126, 129,
	112, 183, 45, 110, 19, 45, 160, 124, 125, 128,
	127, 126, 129, 190, 193, 189, 191, 192, 220, 209,
	108, 106, 45, 45, 224, 221, 222, 38, 223, 131,
	130, 228, 165, 229, 21, 206, 207, 231, 139, 165,
	233, 226, 131, 130, 237, 238, 243, 231, 239, 242,
	186, 187, 158, 159, 157, 156, 155, 154, 153, 120,
	219, 244, 96, 245, 231, 231, 50, 49, 48, 47,
	46, 140, 134, 135, 132, 133, 3, 1, 103, 20,
	22, 116, 75, 177, 56, 93, 216, 97, 212, 161,
	143, 141, 9, 8, 7, 6, 12, 11, 10, 5,
	4, 2,
}
var yyPact = [...]int{

	125, -1000, 125, 196, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 53, 88, -10, 25, 103, 99, 96,
	189, -1000, 114, -1000, -1000, 18, 236, 235, 234, 233,
	232, 24, -11, 23, 77, 20, 15, 11, -1000, 5,
	53, -1000, -20, -2, -1000, -3, -4, -5, -6, -8,
	-12, 120, 83, -1000, 16, -11, 83, -14, 228, 65,
	-1000, -23, -1000, -1000, -1000, -1000, 186, 185, 168, 165,
	155, -13, 225, -1000, -1000, 148, 143, -1000, 28, 179,
	250, 246, -1000, -1000, -1000, -1000, 156, 28, -1000, -1000,
	83, -1000, -1000, 202, -1000, 239, -24, -1000, 44, -36,
	42, 70, 83, 70, -1000, 55, -1000, 223, -1000, 222,
	-1000, 221, -1000, 220, -1000, 219, 217, -1000, 170, -1000,
	-28, 16, 16, 179, 28, 28, 28, 28, 28, 28,
	28, 39, 28, 28, 28, 28, -1000, 166, -1000, -14,
	28, 215, -1000, -1000, 171, -15, 171, -18, -19, 83,
	54, -1000, -21, -1000, -1000, -1000, -1000, -1000, 74, -22,
	225, 200, -1000, -1000, -1000, -1000, -1000, -1000, 143, -1000,
	250, 250, 250, 250, 250, 250, 250, -1000, 137, 246,
	246, -1000, -1000, -1000, -1000, 250, -1000, -24, 105, 226,
	-1000, -1000, -1000, -1000, -1000, 105, 171, 171, 54, 188,
	-1000, 106, -32, 225, -1000, -1000, -1000, -28, 32, -1000,
	105, -1000, -1000, 62, -28, -1000, -1000, 118, 117, 29,
	105, 105, 105, 188, 54, -41, 56, 16, 170, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, 211, 105, 105, -1000,
	156, 16, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 0, 18, 281, 256, 280, 279, 278, 277, 276,
	275, 274, 273, 272, 271, 11, 270, 2, 269, 10,
	7, 6, 268, 267, 266, 8, 265, 264, 14, 5,
	13, 17, 12, 263, 20, 21, 23, 262, 19, 1,
	261, 127, 260, 16, 258, 15, 3, 9, 4, 257,
}
var yyR1 = [...]int{

	0, 49, 49, 3, 3, 3, 4, 4, 6, 6,
	6, 5, 5, 5, 5, 7, 14, 14, 15, 16,
	16, 25, 25, 25, 25, 25, 17, 17, 21, 22,
	22, 22, 22, 24, 24, 8, 23, 23, 23, 23,
	23, 23, 23, 9, 10, 10, 47, 47, 47, 46,
	46, 42, 42, 41, 41, 41, 41, 41, 41, 41,
	41, 41, 41, 41, 41, 41, 41, 41, 11, 11,
	40, 40, 20, 20, 19, 18, 18, 29, 29, 29,
	29, 29, 43, 43, 44, 44, 45, 45, 12, 12,
	12, 12, 2, 2, 38, 38, 39, 13, 27, 26,
	26, 28, 37, 37, 36, 36, 35, 35, 34, 34,
	34, 34, 34, 34, 34, 34, 34, 33, 32, 32,
	32, 31, 31, 31, 30, 30, 30, 30, 30, 30,
	1, 48, 48,
}
var yyR2 = [...]int{

	0, 1, 0, 3, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 6, 3, 1, 1, 3,
	2, 4, 1, 1, 1, 1, 2, 1, 1, 2,
	2, 1, 1, 2, 2, 4, 3, 4, 3, 5,
	4, 4, 5, 3, 8, 7, 3, 1, 0, 3,
	4, 3, 1, 1, 2, 1, 3, 2, 4, 5,
	4, 5, 4, 5, 4, 5, 4, 5, 8, 5,
	3, 1, 3, 1, 3, 3, 1, 1, 1, 1,
	1, 1, 1, 0, 2, 1, 6, 5, 4, 5,
	3, 4, 2, 1, 2, 0, 1, 4, 2, 3,
	1, 3, 3, 1, 3, 1, 2, 1, 3, 3,
	3, 3, 3, 3, 3, 3, 1, 3, 3, 3,
	1, 3, 3, 1, 1, 1, 1, 1, 2, 3,
	2, 1, 1,
}
var yyChk = [...]int{

	-1000, -49, -3, -4, -5, -6, -10, -11, -12, -13,
	-7, -8, -9, 36, 34, 33, 39, 31, 41, 59,
	-4, 18, -42, -41, 6, 68, 47, 48, 49, 50,
	51, 35, 68, 38, 68, 32, 32, 32, 18, 38,
	16, -2, -1, 63, 68, 17, 14, 14, 14, 14,
	14, 68, -2, -38, 37, 68, -27, 40, 68, 68,
	68, 68, -41, -2, 68, 68, 68, 68, 68, 68,
	68, 14, 42, -38, -39, -37, -36, -35, 21, -34,
	-32, -31, -30, 66, -48, 69, 68, 14, 61, 62,
	-2, -38, -38, -26, -28, 68, 14, -23, 59, 64,
	41, -2, -43, -44, -45, 45, 15, -1, 15, -1,
	15, -1, 15, -1, 15, -1, -40, 68, -20, -19,
	14, 19, 20, -34, 8, 9, 12, 11, 10, 13,
	44, 43, 4, 5, 6, 7, -1, -34, -38, 16,
	12, -14, -15, -16, 68, 65, 68, 65, 65, -43,
	-38, -45, 56, 15, 15, 15, 15, 15, 15, 16,
	16, -18, -29, 69, 66, -48, 67, 52, -36, -35,
	-32, -32, -32, -32, -32, -32, -32, -33, 66, -31,
	-31, -30, -30, 15, -28, -32, 15, 16, -25, 24,
	22, 25, 26, 23, 68, -25, 68, 68, -38, -47,
	-46, 58, 68, 42, 68, -19, 15, 16, 20, -15,
	-17, -21, -22, 21, 60, 55, -24, 27, 28, 14,
	-17, -25, -25, -47, 16, 30, -2, 57, -20, -29,
	66, -21, 52, -29, 29, 29, 66, -17, -17, -46,
	68, 57, -39, 15, -1, -39,
}
var yyDef = [...]int{

	2, -2, 1, 5, 6, 7, 11, 12, 13, 14,
	8, 9, 10, 0, 0, 0, 0, 0, 0, 0,
	0, 4, 0, 52, 53, 55, 0, 0, 0, 0,
	0, 0, 95, 0, 0, 0, 0, 0, 3, 0,
	0, 54, 57, 0, 93, 0, 0, 0, 0, 0,
	0, 0, 95, 90, 0, 95, 95, 0, 0, 0,
	43, 83, 51, 56, 92, 130, 0, 0, 0, 0,
	0, 0, 0, 88, 94, 96, 103, 105, 0, 107,
	116, 120, 123, 124, 125, 126, 127, 0, 131, 132,
	95, 91, 97, 98, 100, 0, 0, 35, 0, 0,
	0, 83, 95, 82, 85, 0, 58, 0, 60, 0,
	62, 0, 64, 0, 66, 0, 0, 71, 69, 73,
	0, 0, 0, 106, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 128, 0, 89, 0,
	0, 0, 17, 18, 0, 0, 0, 0, 0, 95,
	48, 84, 0, 59, 61, 63, 65, 67, 0, 0,
	0, 0, 76, 77, 78, 79, 80, 81, 102, 104,
	108, 109, 110, 111, 112, 113, 114, 115, 0, 118,
	119, 121, 122, 129, 99, 101, 15, 0, 20, 0,
	22, 23, 24, 25, 36, 38, 0, 0, 48, 45,
	47, 0, 0, 0, 70, 72, 74, 0, 0, 16,
	19, 27, 28, 0, 0, 31, 32, 0, 0, 0,
	37, 40, 41, 44, 0, 0, 0, 0, 68, 75,
	117, 26, 29, 30, 33, 34, 0, 39, 42, 46,
	49, 0, 87, 21, 50, 86,
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
		//line parser.y:86
		{
			lock.Lock()
			statements = yyDollar[1].stmt_list_t
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:87
		{
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:90
		{
			yyVAL.stmt_list_t = yyDollar[1].stmt_list_t
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[2].stmt_t)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:91
		{
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[1].stmt_t)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:92
		{
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[1].stmt_t)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:95
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:96
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:99
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:100
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:101
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:104
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:105
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:106
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:107
		{
			yyVAL.stmt_t = yyDollar[1].stmt_t
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:110
		{
			yyVAL.stmt_t = &createStatement{yyDollar[3].string_t, yyDollar[5].col_list_t}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:113
		{
			yyVAL.col_list_t = yyDollar[1].col_list_t
			yyVAL.col_list_t = append(yyVAL.col_list_t, yyDollar[3].col_t)
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:114
		{
			yyVAL.col_list_t = append(yyVAL.col_list_t, yyDollar[1].col_t)
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:117
		{
			yyVAL.col_t = yyDollar[1].col_t
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:120
		{
			yyVAL.col_t = &columnDefinition{yyDollar[1].string_t, yyDollar[2].data_t, yyDollar[3].obj_list_t}
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:121
		{
			yyVAL.col_t = &columnDefinition{yyDollar[1].string_t, yyDollar[2].data_t, nil}
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:124
		{
			yyVAL.data_t = &charType{yyDollar[3].int64_t}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:125
		{
			yyVAL.data_t = &integerType{}
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:126
		{
			yyVAL.data_t = &booleanType{}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:127
		{
			yyVAL.data_t = &datetimeType{}
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:128
		{
			yyVAL.data_t = &floatType{}
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:131
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[2].obj_t)
		}
	case 27:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:132
		{
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[1].obj_t)
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:135
		{
			yyVAL.obj_t = yyDollar[1].obj_t
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:138
		{
			yyVAL.obj_t = &notNullConstraint{}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:139
		{
			yyVAL.obj_t = &defaultConstraint{yyDollar[2].obj_t}
		}
	case 31:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:140
		{
			yyVAL.obj_t = &autoincrementConstraint{}
		}
	case 32:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:141
		{
			yyVAL.obj_t = yyDollar[1].obj_t
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:144
		{
			yyVAL.obj_t = &primaryKeyConstraint{}
		}
	case 34:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:145
		{
			yyVAL.obj_t = &foreignKeyConstraint{}
		}
	case 35:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:148
		{
			yyVAL.stmt_t = &alterStatement{yyDollar[3].string_t, yyDollar[4].obj_t}
		}
	case 36:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:151
		{
			yyVAL.obj_t = &alterDrop{yyDollar[3].string_t}
		}
	case 37:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:152
		{
			yyVAL.obj_t = &alterAdd{yyDollar[2].string_t, yyDollar[3].data_t, yyDollar[4].obj_list_t}
		}
	case 38:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:153
		{
			yyVAL.obj_t = &alterAdd{yyDollar[2].string_t, yyDollar[3].data_t, nil}
		}
	case 39:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:154
		{
			yyVAL.obj_t = &alterAdd{yyDollar[3].string_t, yyDollar[4].data_t, yyDollar[5].obj_list_t}
		}
	case 40:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:155
		{
			yyVAL.obj_t = &alterAdd{yyDollar[3].string_t, yyDollar[4].data_t, nil}
		}
	case 41:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:156
		{
			yyVAL.obj_t = &alterModify{yyDollar[3].string_t, yyDollar[4].data_t, nil}
		}
	case 42:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:157
		{
			yyVAL.obj_t = &alterModify{yyDollar[3].string_t, yyDollar[4].data_t, yyDollar[5].obj_list_t}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:160
		{
			yyVAL.stmt_t = &dropStatement{yyDollar[3].string_t}
		}
	case 44:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:163
		{
			yyVAL.stmt_t = &selectStatement{yyDollar[2].columnSpec_list_t, yyDollar[4].string_t, yyDollar[5].string_t, yyDollar[6].joinSpec_list_t, yyDollar[7].expr_t, yyDollar[8].group_by_list_t}
		}
	case 45:
		yyDollar = yyS[yypt-7 : yypt+1]
		//line parser.y:164
		{
			yyVAL.stmt_t = &selectStatement{yyDollar[2].columnSpec_list_t, yyDollar[4].string_t, "", yyDollar[5].joinSpec_list_t, yyDollar[6].expr_t, yyDollar[7].group_by_list_t}
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:167
		{
			yyVAL.group_by_list_t = yyDollar[1].group_by_list_t
			yyVAL.group_by_list_t = append(yyVAL.group_by_list_t, yyDollar[3].group_by_t)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:168
		{
			yyVAL.group_by_list_t = append(yyVAL.group_by_list_t, yyDollar[1].group_by_t)
		}
	case 48:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:169
		{
			yyVAL.group_by_list_t = nil
		}
	case 49:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:172
		{
			yyVAL.group_by_t = GroupBySpec{"", yyDollar[3].string_t}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:173
		{
			yyVAL.group_by_t = GroupBySpec{yyDollar[3].string_t, yyDollar[4].string_t}
		}
	case 51:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:177
		{
			yyVAL.columnSpec_list_t = yyDollar[1].columnSpec_list_t
			yyVAL.columnSpec_list_t = append(yyVAL.columnSpec_list_t, yyDollar[3].columnSpec_t)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:178
		{
			yyVAL.columnSpec_list_t = append(yyVAL.columnSpec_list_t, yyDollar[1].columnSpec_t)
		}
	case 53:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:181
		{
			yyVAL.columnSpec_t = columnSpec{true, "", "", "", nil}
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:182
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, "", yyDollar[2].string_t, nil}
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:183
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, "", "", nil}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:184
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, yyDollar[2].string_t, yyDollar[3].string_t, nil}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:185
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[1].string_t, yyDollar[2].string_t, "", nil}
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:186
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, "", "", &functionSum{}}
		}
	case 59:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:187
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, yyDollar[4].string_t, "", &functionSum{}}
		}
	case 60:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:188
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, "", "", &functionCount{}}
		}
	case 61:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:189
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, yyDollar[4].string_t, "", &functionCount{}}
		}
	case 62:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:190
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, "", "", &functionAvg{}}
		}
	case 63:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:191
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, yyDollar[4].string_t, "", &functionAvg{}}
		}
	case 64:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:192
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, "", "", &functionMin{}}
		}
	case 65:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:193
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, yyDollar[4].string_t, "", &functionMin{}}
		}
	case 66:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:194
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, "", "", &functionMax{}}
		}
	case 67:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:195
		{
			yyVAL.columnSpec_t = columnSpec{false, yyDollar[3].string_t, yyDollar[4].string_t, "", &functionMax{}}
		}
	case 68:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:197
		{
			yyVAL.stmt_t = &insertStatement{yyDollar[3].string_t, yyDollar[5].string_list_t, yyDollar[8].obj_list_t}
		}
	case 69:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:198
		{
			yyVAL.stmt_t = &insertStatement{yyDollar[3].string_t, nil, yyDollar[5].obj_list_t}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:201
		{
			yyVAL.string_list_t = yyDollar[1].string_list_t
			yyVAL.string_list_t = append(yyVAL.string_list_t, yyDollar[3].string_t)
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:202
		{
			yyVAL.string_list_t = append(yyVAL.string_list_t, yyDollar[1].string_t)
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:205
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[3].obj_list_t)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:206
		{
			yyVAL.obj_list_t = append(yyDollar[1].obj_list_t)
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:209
		{
			yyVAL.obj_list_t = yyDollar[2].obj_list_t
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:212
		{
			yyVAL.obj_list_t = yyDollar[1].obj_list_t
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[3].obj_t)
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:213
		{
			yyVAL.obj_list_t = append(yyVAL.obj_list_t, yyDollar[1].obj_t)
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:216
		{
			yyVAL.obj_t = yyDollar[1].string_t
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:217
		{
			yyVAL.obj_t = yyDollar[1].int64_t
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:218
		{
			yyVAL.obj_t = yyDollar[1].bool_t
		}
	case 80:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:219
		{
			yyVAL.obj_t = yyDollar[1].float64_t
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:220
		{
			yyVAL.obj_t = nil
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:223
		{
			yyVAL.joinSpec_list_t = yyDollar[1].joinSpec_list_t
		}
	case 83:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:224
		{
			yyVAL.joinSpec_list_t = nil
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:227
		{
			yyVAL.joinSpec_list_t = yyDollar[1].joinSpec_list_t
			yyVAL.joinSpec_list_t = append(yyVAL.joinSpec_list_t, yyDollar[2].joinSpec_t)
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:228
		{
			yyVAL.joinSpec_list_t = append(yyVAL.joinSpec_list_t, yyDollar[1].joinSpec_t)
		}
	case 86:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:231
		{
			yyVAL.joinSpec_t = joinSpec{yyDollar[3].string_t, yyDollar[4].string_t, yyDollar[6].expr_t}
		}
	case 87:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:232
		{
			yyVAL.joinSpec_t = joinSpec{yyDollar[3].string_t, "", yyDollar[5].expr_t}
		}
	case 88:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:235
		{
			yyVAL.stmt_t = &deleteStatement{yyDollar[2].string_t, yyDollar[3].string_t, yyDollar[4].expr_t}
		}
	case 89:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:236
		{
			yyVAL.stmt_t = &deleteStatement{yyDollar[3].string_t, yyDollar[4].string_t, yyDollar[5].expr_t}
		}
	case 90:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:237
		{
			yyVAL.stmt_t = &deleteStatement{yyDollar[2].string_t, "", yyDollar[3].expr_t}
		}
	case 91:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:238
		{
			yyVAL.stmt_t = &deleteStatement{yyDollar[3].string_t, "", yyDollar[4].expr_t}
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:241
		{
			yyVAL.string_t = yyDollar[2].string_t
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:242
		{
			yyVAL.string_t = yyDollar[1].string_t
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:245
		{
			yyVAL.expr_t = yyDollar[2].expr_t
		}
	case 95:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:246
		{
			yyVAL.expr_t = nil
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:249
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 97:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:252
		{
			yyVAL.stmt_t = &updateStatement{yyDollar[2].string_t, yyDollar[3].assignments_list, yyDollar[4].expr_t}
		}
	case 98:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:255
		{
			yyVAL.assignments_list = yyDollar[2].assignments_list
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:258
		{
			yyVAL.assignments_list = yyDollar[1].assignments_list
			yyVAL.assignments_list = append(yyDollar[1].assignments_list, yyDollar[3].assignment_t)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:259
		{
			yyVAL.assignments_list = append(yyVAL.assignments_list, yyDollar[1].assignment_t)
		}
	case 101:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:262
		{
			yyVAL.assignment_t = assignment{yyDollar[1].string_t, yyDollar[3].expr_t}
		}
	case 102:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:265
		{
			yyVAL.expr_t = &orExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:266
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 104:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:269
		{
			yyVAL.expr_t = &andExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 105:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:270
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 106:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:273
		{
			yyVAL.expr_t = &notExpression{yyDollar[2].expr_t}
		}
	case 107:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:274
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 108:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:277
		{
			yyVAL.expr_t = &ltExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 109:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:278
		{
			yyVAL.expr_t = &gtExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 110:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:279
		{
			yyVAL.expr_t = &eqExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 111:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:280
		{
			yyVAL.expr_t = &lteExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 112:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:281
		{
			yyVAL.expr_t = &gteExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 113:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:282
		{
			yyVAL.expr_t = &neExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 114:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:283
		{
			yyVAL.expr_t = &likeExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 115:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:284
		{
			yyVAL.expr_t = &betweenExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 116:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:285
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 117:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:288
		{
			yyVAL.expr_t = &betweenExpression{&intExpression{yyDollar[1].int64_t}, &intExpression{yyDollar[3].int64_t}}
		}
	case 118:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:291
		{
			yyVAL.expr_t = &sumExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 119:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:292
		{
			yyVAL.expr_t = &subExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 120:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:293
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 121:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:296
		{
			yyVAL.expr_t = &multExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 122:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:297
		{
			yyVAL.expr_t = &divExpression{yyDollar[1].expr_t, yyDollar[3].expr_t}
		}
	case 123:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:298
		{
			yyVAL.expr_t = yyDollar[1].expr_t
		}
	case 124:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:301
		{
			yyVAL.expr_t = &intExpression{yyDollar[1].int64_t}
		}
	case 125:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:302
		{
			yyVAL.expr_t = &boolExpression{yyDollar[1].bool_t}
		}
	case 126:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:303
		{
			yyVAL.expr_t = &stringExpression{yyDollar[1].string_t}
		}
	case 127:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:304
		{
			yyVAL.expr_t = &idExpression{yyDollar[1].string_t, ""}
		}
	case 128:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:305
		{
			yyVAL.expr_t = &idExpression{yyDollar[1].string_t, yyDollar[2].string_t}
		}
	case 129:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:306
		{
			yyVAL.expr_t = yyDollar[2].expr_t
		}
	case 130:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:309
		{
			yyVAL.string_t = yyDollar[2].string_t
		}
	case 131:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:312
		{
			yyVAL.bool_t = true
		}
	case 132:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:313
		{
			yyVAL.bool_t = false
		}
	}
	goto yystack /* stack new state and value */
}

//line parser.y:4
package parser

import __yyfmt__ "fmt"

//line parser.y:4
import "io"

//line parser.y:10
type yySymType struct {
	yys      int
	int_t    int
	string_t string
	float_t  float64

	expr_t expression

	stmt_list_t statementList
	stmt_t      statement

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

//line parser.y:224

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

func Parse(in io.Reader) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	yyParse(NewLexer(in))

	return nil
}

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 189

var yyAct = [...]int{

	155, 150, 141, 142, 133, 105, 62, 64, 75, 42,
	61, 63, 69, 36, 107, 59, 43, 58, 24, 60,
	157, 69, 37, 156, 138, 136, 39, 38, 76, 108,
	84, 53, 52, 49, 48, 47, 46, 40, 28, 27,
	158, 152, 123, 37, 109, 80, 72, 154, 38, 137,
	79, 55, 45, 144, 73, 72, 70, 71, 43, 82,
	26, 65, 31, 68, 67, 70, 71, 30, 29, 37,
	65, 87, 68, 67, 38, 34, 17, 25, 15, 14,
	101, 13, 100, 146, 16, 35, 18, 139, 135, 145,
	134, 110, 86, 33, 85, 115, 116, 117, 118, 119,
	120, 121, 114, 113, 23, 19, 126, 127, 124, 125,
	130, 129, 32, 41, 21, 39, 88, 89, 92, 91,
	90, 93, 51, 128, 88, 89, 92, 91, 90, 93,
	165, 166, 131, 132, 160, 81, 111, 112, 140, 50,
	102, 148, 163, 151, 147, 153, 77, 95, 94, 54,
	103, 159, 162, 98, 99, 95, 94, 96, 97, 66,
	122, 3, 164, 153, 20, 74, 44, 167, 57, 56,
	161, 149, 83, 22, 78, 143, 106, 104, 1, 9,
	8, 7, 6, 12, 11, 10, 5, 4, 2,
}
var yyPact = [...]int{

	49, -1000, 49, 96, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 12, 29, -26, -27, 40, 39, 34,
	94, -1000, 59, -1000, -1000, 9, -28, -17, 16, -29,
	-30, -31, -1000, -32, 12, -1000, -38, -33, -1000, -34,
	135, 25, -1000, -2, 25, -37, 132, -11, -1000, -17,
	-1000, -1000, -1000, -1000, -35, -1000, -1000, 75, 72, -1000,
	7, 116, 153, 147, -1000, -1000, -1000, -1000, 98, 7,
	-1000, -1000, -1000, -1000, 124, -1000, 138, -51, -1000, -36,
	-18, 25, -1000, 121, -1000, -2, -2, 116, 7, 7,
	7, 7, 7, 7, 7, -21, 7, 7, 7, 7,
	-1000, 108, -37, 7, 117, -1000, -1000, 66, -1000, -40,
	-1000, 11, -41, 72, -1000, 153, 153, 153, 153, 153,
	153, 153, -1000, 67, 147, 147, -1000, -1000, -1000, -1000,
	153, -1000, -51, 32, 130, -1000, 66, 129, -1000, -22,
	-1000, 32, -1000, -1000, -1, -43, -1000, -23, 32, 118,
	-1000, -43, -1000, -1000, -1000, -1000, -1000, -1000, 127, 32,
	129, 115, -1000, -1000, -1000, -1000, -43, -1000,
}
var yyPgo = [...]int{

	0, 188, 161, 187, 186, 185, 184, 183, 182, 181,
	180, 179, 178, 177, 5, 176, 4, 2, 3, 175,
	0, 174, 173, 85, 9, 104, 13, 172, 171, 1,
	170, 169, 168, 166, 165, 8, 6, 17, 15, 10,
	160, 11, 7, 159,
}
var yyR1 = [...]int{

	0, 12, 12, 1, 1, 1, 2, 2, 4, 4,
	4, 3, 3, 3, 3, 5, 13, 13, 14, 15,
	15, 16, 16, 17, 17, 18, 19, 19, 19, 6,
	21, 21, 7, 8, 8, 22, 22, 25, 25, 25,
	25, 25, 9, 27, 27, 28, 28, 29, 30, 30,
	20, 20, 10, 10, 23, 23, 24, 24, 31, 11,
	33, 34, 34, 35, 32, 32, 37, 37, 38, 38,
	39, 39, 39, 39, 39, 39, 39, 39, 39, 40,
	36, 36, 36, 41, 41, 41, 42, 42, 42, 42,
	42, 42, 26, 43, 43, 43,
}
var yyR2 = [...]int{

	0, 1, 0, 3, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 6, 3, 1, 1, 3,
	2, 4, 1, 2, 1, 1, 2, 2, 1, 4,
	2, 5, 3, 6, 5, 3, 1, 1, 2, 1,
	3, 2, 8, 3, 1, 3, 1, 3, 3, 1,
	1, 1, 4, 3, 2, 1, 2, 0, 1, 4,
	2, 3, 1, 3, 3, 1, 3, 1, 2, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 1, 3,
	3, 3, 1, 3, 3, 1, 1, 1, 1, 1,
	2, 3, 2, 1, 1, 1,
}
var yyChk = [...]int{

	-1000, -12, -1, -2, -3, -4, -8, -9, -10, -11,
	-5, -6, -7, 32, 30, 29, 35, 27, 37, 56,
	-2, 18, -22, -25, 6, 65, 31, 65, 65, 28,
	28, 28, 18, 34, 16, -23, -26, 60, 65, 17,
	65, -23, -24, 33, -33, 36, 65, 65, 65, 65,
	-25, -23, 65, 65, 14, -24, -31, -32, -37, -38,
	21, -39, -36, -41, -42, 63, -43, 66, 65, 14,
	58, 59, 48, -24, -34, -35, 65, 14, -21, 61,
	56, -23, -24, -27, 65, 19, 20, -39, 8, 9,
	12, 11, 10, 13, 40, 39, 4, 5, 6, 7,
	-26, -39, 16, 12, -13, -14, -15, 65, 65, 62,
	-24, 15, 16, -37, -38, -36, -36, -36, -36, -36,
	-36, -36, -40, 63, -41, -41, -42, -42, 15, -35,
	-36, 15, 16, -16, 24, 22, 65, 38, 65, 20,
	-14, -17, -18, -19, 21, 57, 51, 14, -16, -28,
	-29, 14, 63, -18, 48, -20, 66, 63, 63, -17,
	16, -30, -20, 15, -29, 15, 16, -20,
}
var yyDef = [...]int{

	2, -2, 1, 5, 6, 7, 11, 12, 13, 14,
	8, 9, 10, 0, 0, 0, 0, 0, 0, 0,
	0, 4, 0, 36, 37, 39, 0, 57, 0, 0,
	0, 0, 3, 0, 0, 38, 41, 0, 55, 0,
	0, 57, 53, 0, 57, 0, 0, 0, 32, 57,
	35, 40, 54, 92, 0, 52, 56, 58, 65, 67,
	0, 69, 78, 82, 85, 86, 87, 88, 89, 0,
	93, 94, 95, 59, 60, 62, 0, 0, 29, 0,
	0, 57, 34, 0, 44, 0, 0, 68, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	90, 0, 0, 0, 0, 17, 18, 0, 30, 0,
	33, 0, 0, 64, 66, 70, 71, 72, 73, 74,
	75, 76, 77, 0, 80, 81, 83, 84, 91, 61,
	63, 15, 0, 20, 0, 22, 0, 0, 43, 0,
	16, 19, 24, 25, 0, 0, 28, 0, 0, 42,
	46, 0, 79, 23, 26, 27, 50, 51, 0, 31,
	0, 0, 49, 21, 45, 47, 0, 48,
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
		//line parser.y:43
		{
			yyDollar[1].stmt_list_t.execute()
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:44
		{
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:47
		{
			yyVAL.stmt_list_t = yyDollar[1].stmt_list_t
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[2].stmt_t)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:48
		{
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[1].stmt_t)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:49
		{
			yyVAL.stmt_list_t = append(yyVAL.stmt_list_t, yyDollar[1].stmt_t)
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:52
		{
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:53
		{
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:56
		{
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:57
		{
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:58
		{
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:61
		{
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:62
		{
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:63
		{
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:64
		{
		}
	case 15:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:67
		{
			yyVAL.stmt_t = &createStatement{}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:70
		{
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:71
		{
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:74
		{
		}
	case 19:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:77
		{
		}
	case 20:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:78
		{
		}
	case 21:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:81
		{
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:82
		{
		}
	case 23:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:85
		{
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:86
		{
		}
	case 25:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:89
		{
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:92
		{
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:93
		{
		}
	case 28:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:94
		{
		}
	case 29:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:97
		{
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:100
		{
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:101
		{
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:104
		{
			yyVAL.stmt_t = &dropStatement{}
		}
	case 33:
		yyDollar = yyS[yypt-6 : yypt+1]
		//line parser.y:107
		{
			yyVAL.stmt_t = &selectStatement{}
		}
	case 34:
		yyDollar = yyS[yypt-5 : yypt+1]
		//line parser.y:108
		{
			yyVAL.stmt_t = &selectStatement{}
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:111
		{
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:112
		{
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:115
		{
		}
	case 38:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:116
		{
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:117
		{
		}
	case 40:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:118
		{
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:119
		{
		}
	case 42:
		yyDollar = yyS[yypt-8 : yypt+1]
		//line parser.y:122
		{
			yyVAL.stmt_t = &insertStatement{}
		}
	case 43:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:125
		{
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:126
		{
		}
	case 45:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:129
		{
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:130
		{
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:144
		{
			yyVAL.stmt_t = &deleteStatement{}
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:145
		{
			yyVAL.stmt_t = &deleteStatement{}
		}
	case 54:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:148
		{
		}
	case 55:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:149
		{
		}
	case 56:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:152
		{
		}
	case 57:
		yyDollar = yyS[yypt-0 : yypt+1]
		//line parser.y:153
		{
		}
	case 59:
		yyDollar = yyS[yypt-4 : yypt+1]
		//line parser.y:159
		{
			yyVAL.stmt_t = &updateStatement{}
		}
	case 60:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:162
		{
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:165
		{
		}
	case 62:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:166
		{
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:169
		{
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:172
		{
		}
	case 65:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:173
		{
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:176
		{
		}
	case 67:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:177
		{
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:180
		{
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:181
		{
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:184
		{
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:185
		{
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:186
		{
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:187
		{
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:188
		{
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:189
		{
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:190
		{
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:191
		{
		}
	case 78:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:192
		{
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:195
		{
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:198
		{
		}
	case 81:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:199
		{
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:200
		{
		}
	case 83:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:203
		{
		}
	case 84:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:204
		{
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:205
		{
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:208
		{
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:209
		{
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:210
		{
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:211
		{
		}
	case 90:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:212
		{
		}
	case 91:
		yyDollar = yyS[yypt-3 : yypt+1]
		//line parser.y:213
		{
		}
	case 92:
		yyDollar = yyS[yypt-2 : yypt+1]
		//line parser.y:216
		{
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:219
		{
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
		//line parser.y:220
		{
		}
	}
	goto yystack /* stack new state and value */
}

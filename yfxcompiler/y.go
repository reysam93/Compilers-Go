//line yfxcompiler.y:4
package yfxcompiler

import __yyfmt__ "fmt"

//line yfxcompiler.y:4
import (
//"strings"
)

var (
	level       int
	DebugParser bool
)

//line yfxcompiler.y:16
type FXSymType struct {
	yys   int
	token *Token

	// AST nodes
	expr       *Expr
	subExpr    *SubExpr
	args       []*Expr
	sentence   *Sentence
	declar     *Declaration
	func_call  *FunctionCall
	if_stm     *IfStatement
	body       Body
	loop       *Loop
	asign      *Asignation
	parameters []*Declaration
	function   *Function
	funcs      []*Function
}

const TYPE = 57346
const RECORD = 57347
const FUNC = 57348
const ITER = 57349
const IF = 57350
const ELSE = 57351
const LOOP_EQ = 57352
const NONE = 57353
const DATA_TYPE = 57354
const BOOL_LIT = 57355
const INT_LIT = 57356
const ID = 57357
const LESS_EQ = 57358
const BIG_EQ = 57359
const POW = 57360

var FXToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TYPE",
	"RECORD",
	"FUNC",
	"ITER",
	"IF",
	"ELSE",
	"LOOP_EQ",
	"NONE",
	"'.'",
	"','",
	"'('",
	"')'",
	"'{'",
	"'}'",
	"'['",
	"']'",
	"';'",
	"DATA_TYPE",
	"BOOL_LIT",
	"INT_LIT",
	"ID",
	"'<'",
	"LESS_EQ",
	"'>'",
	"BIG_EQ",
	"'+'",
	"'-'",
	"'^'",
	"'|'",
	"'*'",
	"'/'",
	"'%'",
	"'&'",
	"POW",
	"'!'",
	"'='",
}
var FXStatenames = [...]string{}

const FXEofCode = 1
const FXErrCode = 2
const FXInitialStackSize = 16

//line yfxcompiler.y:401

func createExpr(op *Token, exp1 *Expr, exp2 *Expr) *Expr {
	var subExpr1, subExpr2 *SubExpr
	if exp1.Op == "" {
		subExpr1 = exp1.SubExpr1
	} else {
		subExpr1 = NewExprSubExpr(exp1)
	}
	if exp2.Op == "" {
		subExpr2 = exp2.SubExpr1
	} else {
		subExpr2 = NewExprSubExpr(exp2)
	}
	return NewExpr(op.lex, subExpr1, subExpr2)
}

//line yacctab:1
var FXExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 13,
	7, 23,
	-2, 0,
}

const FXNprod = 56
const FXPrivate = 57344

var FXTokenNames []string
var FXStates []string

const FXLast = 181

var FXAct = [...]int{

	46, 9, 102, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 59, 60, 61, 62, 63, 64, 65, 66,
	67, 68, 69, 70, 71, 52, 67, 68, 69, 70,
	71, 40, 43, 51, 50, 49, 71, 7, 57, 42,
	32, 54, 55, 16, 17, 76, 37, 39, 72, 47,
	38, 36, 35, 73, 14, 75, 41, 74, 45, 6,
	79, 80, 81, 82, 83, 84, 85, 86, 87, 88,
	89, 90, 91, 97, 77, 44, 93, 94, 95, 31,
	96, 30, 34, 12, 99, 59, 60, 61, 62, 63,
	64, 65, 66, 67, 68, 69, 70, 71, 100, 92,
	33, 5, 101, 103, 104, 78, 22, 8, 2, 59,
	60, 61, 62, 63, 64, 65, 66, 67, 68, 69,
	70, 71, 58, 1, 48, 25, 56, 19, 13, 98,
	21, 26, 59, 60, 61, 62, 63, 64, 65, 66,
	67, 68, 69, 70, 71, 59, 60, 61, 62, 63,
	64, 65, 66, 67, 68, 69, 70, 71, 27, 20,
	11, 24, 53, 15, 23, 4, 3, 0, 0, 0,
	0, 0, 0, 18, 10, 0, 0, 29, 0, 0,
	28,
}
var FXPact = [...]int{

	-1000, -1000, -1000, 95, -1000, 35, -1000, 158, 69, -1000,
	-1000, 37, 23, 156, -1000, 66, -1000, 16, -1000, -1000,
	-1000, -1000, 93, 68, 32, 31, 26, 30, 17, 15,
	158, 23, -1000, 44, 11, -1000, -1000, -1000, -1000, -1000,
	11, 11, -1000, -1000, -1000, 14, 107, 11, -1000, -1000,
	-1000, -1000, 11, 42, 120, 120, 25, 64, -1000, 11,
	11, 11, 11, 11, 11, 11, 11, 11, 11, 11,
	11, 11, -1000, 84, -1000, 11, 11, 11, 158, -26,
	-26, -26, -26, -7, -7, -7, -7, -1, -1, -1,
	-1, -1000, -1000, 120, 60, 120, 75, 11, -1000, -1000,
	-13, 158, 158, -1000, -1000,
}
var FXPgo = [...]int{

	0, 166, 165, 163, 162, 161, 159, 131, 43, 130,
	1, 129, 128, 127, 126, 125, 0, 124, 123, 108,
	107, 106, 105, 102,
}
var FXR1 = [...]int{

	0, 19, 18, 1, 1, 20, 2, 2, 3, 3,
	3, 8, 10, 10, 12, 12, 12, 12, 9, 9,
	9, 9, 9, 21, 13, 14, 22, 6, 23, 11,
	11, 7, 5, 4, 4, 4, 15, 16, 16, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 17, 17, 17, 17,
}
var FXR2 = [...]int{

	0, 0, 2, 2, 0, 0, 7, 3, 1, 3,
	0, 2, 3, 2, 2, 2, 2, 0, 2, 2,
	2, 2, 2, 0, 10, 3, 0, 7, 0, 3,
	0, 2, 4, 1, 3, 0, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 1, 1, 1, 1, 3,
}
var FXChk = [...]int{

	-1000, -18, -19, -1, -2, 6, 24, 2, -20, -10,
	16, 2, 14, -12, 17, -3, -8, 21, 17, -13,
	-6, -9, -21, 8, -5, -15, -7, 2, 24, 21,
	15, 13, 24, 7, 14, 20, 20, 20, 20, 17,
	14, 39, 24, -10, -8, 14, -16, 38, -17, 24,
	23, 22, 14, -4, -16, -16, -14, 24, 15, 25,
	26, 27, 28, 29, 30, 31, 32, 33, 34, 35,
	36, 37, -16, -16, 15, 13, 20, 10, -22, -16,
	-16, -16, -16, -16, -16, -16, -16, -16, -16, -16,
	-16, -16, 15, -16, -16, -16, -10, 13, -11, 9,
	-16, -23, 15, -10, -10,
}
var FXDef = [...]int{

	1, -2, 4, 2, 3, 0, 5, 0, 0, 7,
	17, 0, 10, -2, 13, 0, 8, 0, 12, 14,
	15, 16, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 11, 0, 0, 18, 19, 20, 21, 22,
	35, 0, 31, 6, 9, 0, 0, 0, 51, 52,
	53, 54, 0, 0, 33, 36, 0, 0, 26, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 50, 0, 32, 0, 0, 0, 0, 37,
	38, 39, 40, 41, 42, 43, 44, 45, 46, 47,
	48, 49, 55, 34, 0, 25, 30, 0, 27, 28,
	0, 0, 0, 29, 24,
}
var FXTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 38, 3, 3, 3, 35, 36, 3,
	14, 15, 33, 29, 13, 30, 12, 34, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 20,
	25, 39, 27, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 18, 3, 19, 31, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 16, 32, 17,
}
var FXTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	21, 22, 23, 24, 26, 28, 37,
}
var FXTok3 = [...]int{
	0,
}

var FXErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	FXDebug        = 0
	FXErrorVerbose = false
)

type FXLexer interface {
	Lex(lval *FXSymType) int
	Error(s string)
}

type FXParser interface {
	Parse(FXLexer) int
	Lookahead() int
}

type FXParserImpl struct {
	lval  FXSymType
	stack [FXInitialStackSize]FXSymType
	char  int
}

func (p *FXParserImpl) Lookahead() int {
	return p.char
}

func FXNewParser() FXParser {
	return &FXParserImpl{}
}

const FXFlag = -1000

func FXTokname(c int) string {
	if c >= 1 && c-1 < len(FXToknames) {
		if FXToknames[c-1] != "" {
			return FXToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func FXStatname(s int) string {
	if s >= 0 && s < len(FXStatenames) {
		if FXStatenames[s] != "" {
			return FXStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func FXErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !FXErrorVerbose {
		return "syntax error"
	}

	for _, e := range FXErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + FXTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := FXPact[state]
	for tok := TOKSTART; tok-1 < len(FXToknames); tok++ {
		if n := base + tok; n >= 0 && n < FXLast && FXChk[FXAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if FXDef[state] == -2 {
		i := 0
		for FXExca[i] != -1 || FXExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; FXExca[i] >= 0; i += 2 {
			tok := FXExca[i]
			if tok < TOKSTART || FXExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if FXExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += FXTokname(tok)
	}
	return res
}

func FXlex1(lex FXLexer, lval *FXSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = FXTok1[0]
		goto out
	}
	if char < len(FXTok1) {
		token = FXTok1[char]
		goto out
	}
	if char >= FXPrivate {
		if char < FXPrivate+len(FXTok2) {
			token = FXTok2[char-FXPrivate]
			goto out
		}
	}
	for i := 0; i < len(FXTok3); i += 2 {
		token = FXTok3[i+0]
		if token == char {
			token = FXTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = FXTok2[1] /* unknown char */
	}
	if FXDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", FXTokname(token), uint(char))
	}
	return char, token
}

func FXParse(FXlex FXLexer) int {
	return FXNewParser().Parse(FXlex)
}

func (FXrcvr *FXParserImpl) Parse(FXlex FXLexer) int {
	var FXn int
	var FXVAL FXSymType
	var FXDollar []FXSymType
	_ = FXDollar // silence set and not used
	FXS := FXrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	FXstate := 0
	FXrcvr.char = -1
	FXtoken := -1 // FXrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		FXstate = -1
		FXrcvr.char = -1
		FXtoken = -1
	}()
	FXp := -1
	goto FXstack

ret0:
	return 0

ret1:
	return 1

FXstack:
	/* put a state and value onto the stack */
	if FXDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", FXTokname(FXtoken), FXStatname(FXstate))
	}

	FXp++
	if FXp >= len(FXS) {
		nyys := make([]FXSymType, len(FXS)*2)
		copy(nyys, FXS)
		FXS = nyys
	}
	FXS[FXp] = FXVAL
	FXS[FXp].yys = FXstate

FXnewstate:
	FXn = FXPact[FXstate]
	if FXn <= FXFlag {
		goto FXdefault /* simple state */
	}
	if FXrcvr.char < 0 {
		FXrcvr.char, FXtoken = FXlex1(FXlex, &FXrcvr.lval)
	}
	FXn += FXtoken
	if FXn < 0 || FXn >= FXLast {
		goto FXdefault
	}
	FXn = FXAct[FXn]
	if FXChk[FXn] == FXtoken { /* valid shift */
		FXrcvr.char = -1
		FXtoken = -1
		FXVAL = FXrcvr.lval
		FXstate = FXn
		if Errflag > 0 {
			Errflag--
		}
		goto FXstack
	}

FXdefault:
	/* default state action */
	FXn = FXDef[FXstate]
	if FXn == -2 {
		if FXrcvr.char < 0 {
			FXrcvr.char, FXtoken = FXlex1(FXlex, &FXrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if FXExca[xi+0] == -1 && FXExca[xi+1] == FXstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			FXn = FXExca[xi+0]
			if FXn < 0 || FXn == FXtoken {
				break
			}
		}
		FXn = FXExca[xi+1]
		if FXn < 0 {
			goto ret0
		}
	}
	if FXn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			FXlex.Error(FXErrorMessage(FXstate, FXtoken))
			Nerrs++
			if FXDebug >= 1 {
				__yyfmt__.Printf("%s", FXStatname(FXstate))
				__yyfmt__.Printf(" saw %s\n", FXTokname(FXtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for FXp >= 0 {
				FXn = FXPact[FXS[FXp].yys] + FXErrCode
				if FXn >= 0 && FXn < FXLast {
					FXstate = FXAct[FXn] /* simulate a shift of "error" */
					if FXChk[FXstate] == FXErrCode {
						goto FXstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if FXDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", FXS[FXp].yys)
				}
				FXp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if FXDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", FXTokname(FXtoken))
			}
			if FXtoken == FXEofCode {
				goto ret1
			}
			FXrcvr.char = -1
			FXtoken = -1
			goto FXnewstate /* try again in the same state */
		}
	}

	/* reduction by production FXn */
	if FXDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", FXn, FXStatname(FXstate))
	}

	FXnt := FXn
	FXpt := FXp
	_ = FXpt // guard against "declared and not used"

	FXp -= FXR2[FXn]
	// FXp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if FXp+1 >= len(FXS) {
		nyys := make([]FXSymType, len(FXS)*2)
		copy(nyys, FXS)
		FXS = nyys
	}
	FXVAL = FXS[FXp+1]

	/* consult goto table to find next state */
	FXn = FXR1[FXn]
	FXg := FXPgo[FXn]
	FXj := FXg + FXS[FXp].yys + 1

	if FXj >= FXLast {
		FXstate = FXAct[FXg]
	} else {
		FXstate = FXAct[FXj]
		if FXChk[FXstate] != -FXn {
			FXstate = FXAct[FXg]
		}
	}
	// dummy call; replaced with literal code
	switch FXnt {

	case 1:
		FXDollar = FXS[FXpt-0 : FXpt+1]
		//line yfxcompiler.y:83
		{
			symbEnvs.PushEnv()
		}
	case 2:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:87
		{
			program := NewProgram(FXDollar[2].funcs)
			symbEnvs.PopEnv()
			ast_program = program
		}
	case 3:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:95
		{
			FXVAL.funcs = append(FXDollar[1].funcs, FXDollar[2].function)
		}
	case 4:
		FXDollar = FXS[FXpt-0 : FXpt+1]
		//line yfxcompiler.y:99
		{
			FXVAL.funcs = nil
		}
	case 5:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:106
		{
			if symbEnvs.GetSymb(FXDollar[2].token.lex) != nil {
				Errorf("Function symbol %s already defined", FXDollar[2].token.lex)
			}
			symbEnvs.PutFunction(*FXDollar[2].token)
			symbEnvs.PushEnv()
		}
	case 6:
		FXDollar = FXS[FXpt-7 : FXpt+1]
		//line yfxcompiler.y:114
		{
			FXVAL.function = NewFunc(FXDollar[2].token.lex, FXDollar[5].parameters, FXDollar[7].body)
		}
	case 7:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:118
		{
			Errorf("Wrong function declaration")
			Errflag = 0
			symbEnvs.PushEnv()
			FXVAL.function = NewFunc("", nil, FXDollar[3].body)
		}
	case 8:
		FXDollar = FXS[FXpt-1 : FXpt+1]
		//line yfxcompiler.y:128
		{
			FXVAL.parameters = append(FXVAL.parameters, FXDollar[1].declar)
		}
	case 9:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:132
		{
			FXVAL.parameters = append(FXDollar[1].parameters, FXDollar[3].declar)
		}
	case 10:
		FXDollar = FXS[FXpt-0 : FXpt+1]
		//line yfxcompiler.y:136
		{
			FXVAL.parameters = nil
		}
	case 11:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:143
		{
			if symbEnvs.GetSymb(FXDollar[2].token.lex) != nil {
				Errorf("Variable symbol %s already defined", FXDollar[2].token.lex)
			}
			symbEnvs.PutVar(*FXDollar[2].token)
			FXVAL.declar = NewDeclaration(FXDollar[2].token.lex, FXDollar[1].token.val)
		}
	case 12:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:154
		{
			FXVAL.body = FXDollar[2].body
			symbEnvs.PopEnv()
		}
	case 13:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:159
		{
			Errorf("Wrong body")
			Errflag = 0
			symbEnvs.PopEnv()
			FXVAL.body = nil
		}
	case 14:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:170
		{
			sentence := NewLoopSentence(FXDollar[2].loop)
			FXVAL.body = append(FXDollar[1].body, sentence)
		}
	case 15:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:175
		{
			sentence := NewIfSentence(FXDollar[2].if_stm)
			FXVAL.body = append(FXDollar[1].body, sentence)
		}
	case 16:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:180
		{
			FXVAL.body = append(FXDollar[1].body, FXDollar[2].sentence)
		}
	case 17:
		FXDollar = FXS[FXpt-0 : FXpt+1]
		//line yfxcompiler.y:184
		{
			FXVAL.body = nil
		}
	case 18:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:191
		{
			FXVAL.sentence = NewFuncCallSentence(FXDollar[1].func_call)
		}
	case 19:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:195
		{
			FXVAL.sentence = NewAsginSentence(FXDollar[1].asign)
		}
	case 20:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:199
		{
			FXVAL.sentence = NewDeclSentence(FXDollar[1].declar)
		}
	case 21:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:203
		{
			Errorf("Wrong sentence")
			Errflag = 0
			FXVAL.sentence = NewNullSentence()
		}
	case 22:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:209
		{
			FXVAL.sentence = NewNullSentence()
		}
	case 23:
		FXDollar = FXS[FXpt-0 : FXpt+1]
		//line yfxcompiler.y:216
		{
			symbEnvs.PushEnv()
		}
	case 24:
		FXDollar = FXS[FXpt-10 : FXpt+1]
		//line yfxcompiler.y:220
		{
			FXVAL.loop = NewLoop(FXDollar[4].asign, FXDollar[6].expr, FXDollar[8].expr, FXDollar[10].body)
		}
	case 25:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:227
		{
			if symbEnvs.GetSymb(FXDollar[1].token.lex) == nil {
				Errorf("Symbol %s is not defined", FXDollar[1].token.lex)
			}
			FXVAL.asign = NewAsign(FXDollar[1].token.lex, FXDollar[3].expr)
		}
	case 26:
		FXDollar = FXS[FXpt-4 : FXpt+1]
		//line yfxcompiler.y:237
		{
			symbEnvs.PushEnv()
		}
	case 27:
		FXDollar = FXS[FXpt-7 : FXpt+1]
		//line yfxcompiler.y:241
		{
			FXVAL.if_stm = NewIfStatement(FXDollar[3].expr, FXDollar[6].body, FXDollar[7].body)
		}
	case 28:
		FXDollar = FXS[FXpt-1 : FXpt+1]
		//line yfxcompiler.y:248
		{
			symbEnvs.PushEnv()
		}
	case 29:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:252
		{
			FXVAL.body = FXDollar[3].body
		}
	case 30:
		FXDollar = FXS[FXpt-0 : FXpt+1]
		//line yfxcompiler.y:256
		{
			FXVAL.body = nil
		}
	case 31:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:263
		{
			if symbEnvs.GetSymb(FXDollar[2].token.lex) != nil {
				Errorf("Variable symbol %s already defined", FXDollar[2].token.lex)
			}
			symbEnvs.PutVar(*FXDollar[2].token)
			FXVAL.declar = NewDeclaration(FXDollar[2].token.lex, FXDollar[1].token.val)
		}
	case 32:
		FXDollar = FXS[FXpt-4 : FXpt+1]
		//line yfxcompiler.y:274
		{
			if symbEnvs.GetSymb(FXDollar[1].token.lex) == nil {
				Errorf("Symbol %s is not defined", FXDollar[1].token.lex)
			}
			FXVAL.func_call = NewFuncCall(FXDollar[1].token.lex, FXDollar[3].args)
		}
	case 33:
		FXDollar = FXS[FXpt-1 : FXpt+1]
		//line yfxcompiler.y:284
		{
			FXVAL.args = append(FXVAL.args, FXDollar[1].expr)
		}
	case 34:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:288
		{
			FXVAL.args = append(FXDollar[1].args, FXDollar[3].expr)
		}
	case 35:
		FXDollar = FXS[FXpt-0 : FXpt+1]
		//line yfxcompiler.y:292
		{
			FXVAL.args = nil
		}
	case 36:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:299
		{
			if symbEnvs.GetSymb(FXDollar[1].token.lex) == nil {
				Errorf("Symbol %s is not defined", FXDollar[1].token.lex)
			}
			FXVAL.asign = NewAsign(FXDollar[1].token.lex, FXDollar[3].expr)
		}
	case 37:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:309
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 38:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:313
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 39:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:317
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 40:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:321
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 41:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:326
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 42:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:330
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 43:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:334
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 44:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:338
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 45:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:343
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 46:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:347
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 47:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:351
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 48:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:355
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 49:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:360
		{
			FXVAL.expr = createExpr(FXDollar[2].token, FXDollar[1].expr, FXDollar[3].expr)
		}
	case 50:
		FXDollar = FXS[FXpt-2 : FXpt+1]
		//line yfxcompiler.y:364
		{
			var subExpr *SubExpr
			if FXDollar[2].expr.Op == "" {
				subExpr = FXDollar[2].expr.SubExpr1
			} else {
				subExpr = NewExprSubExpr(FXDollar[2].expr)
			}
			FXVAL.expr = NewExpr(FXDollar[1].token.lex, subExpr, nil)
		}
	case 51:
		FXDollar = FXS[FXpt-1 : FXpt+1]
		//line yfxcompiler.y:374
		{
			FXVAL.expr = FXDollar[1].expr
		}
	case 52:
		FXDollar = FXS[FXpt-1 : FXpt+1]
		//line yfxcompiler.y:381
		{
			if symbEnvs.GetSymb(FXDollar[1].token.lex) == nil {
				Errorf("Symbol %s is not defined", FXDollar[1].token.lex)
			}
			FXVAL.expr = NewExpr("", NewValSubExpr(FXDollar[1].token), nil)
		}
	case 53:
		FXDollar = FXS[FXpt-1 : FXpt+1]
		//line yfxcompiler.y:388
		{
			FXVAL.expr = NewExpr("", NewValSubExpr(FXDollar[1].token), nil)
		}
	case 54:
		FXDollar = FXS[FXpt-1 : FXpt+1]
		//line yfxcompiler.y:392
		{
			FXVAL.expr = NewExpr("", NewValSubExpr(FXDollar[1].token), nil)
		}
	case 55:
		FXDollar = FXS[FXpt-3 : FXpt+1]
		//line yfxcompiler.y:396
		{
			FXVAL.expr = FXDollar[2].expr
		}
	}
	goto FXstack /* stack new state and value */
}

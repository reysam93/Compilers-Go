//line calc_gacc.y:4

// +build ignore
package main

import __yyfmt__ "fmt"

//line calc_gacc.y:5
import (
	"fmt"
)

//line calc_gacc.y:11
type ExprSymType struct {
	yys  int
	num  float64
	name string
}

const NUM = 57346
const PI = 57347
const FN = 57348

var ExprToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"'('",
	"')'",
	"'\\n'",
	"NUM",
	"PI",
	"FN",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
}
var ExprStatenames = [...]string{}

const ExprEofCode = 1
const ExprErrCode = 2
const ExprInitialStackSize = 16

//line calc_gacc.y:49

//line yacctab:1
var ExprExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const ExprNprod = 12
const ExprPrivate = 57344

var ExprTokenNames []string
var ExprStates []string

const ExprLast = 33

var ExprAct = [...]int{

	3, 1, 19, 12, 13, 0, 14, 10, 11, 12,
	13, 15, 16, 17, 18, 9, 2, 0, 8, 10,
	11, 12, 13, 5, 0, 4, 6, 7, 5, 0,
	0, 6, 7,
}
var ExprPact = [...]int{

	19, 19, -1000, 9, -1000, 24, -1000, -1000, -1000, -1000,
	24, 24, 24, 24, -3, -9, -9, -1000, -1000, -1000,
}
var ExprPgo = [...]int{

	0, 0, 1, 16,
}
var ExprR1 = [...]int{

	0, 2, 2, 3, 3, 1, 1, 1, 1, 1,
	1, 1,
}
var ExprR2 = [...]int{

	0, 1, 2, 2, 1, 3, 3, 3, 3, 3,
	1, 1,
}
var ExprChk = [...]int{

	-1000, -2, -3, -1, 6, 4, 7, 8, -3, 6,
	10, 11, 12, 13, -1, -1, -1, -1, -1, 5,
}
var ExprDef = [...]int{

	0, -2, 1, 0, 4, 0, 10, 11, 2, 3,
	0, 0, 0, 0, 0, 5, 6, 7, 8, 9,
}
var ExprTok1 = [...]int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	6, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	4, 5, 12, 10, 3, 11, 3, 13,
}
var ExprTok2 = [...]int{

	2, 3, 7, 8, 9,
}
var ExprTok3 = [...]int{
	0,
}

var ExprErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	ExprDebug        = 0
	ExprErrorVerbose = false
)

type ExprLexer interface {
	Lex(lval *ExprSymType) int
	Error(s string)
}

type ExprParser interface {
	Parse(ExprLexer) int
	Lookahead() int
}

type ExprParserImpl struct {
	lval  ExprSymType
	stack [ExprInitialStackSize]ExprSymType
	char  int
}

func (p *ExprParserImpl) Lookahead() int {
	return p.char
}

func ExprNewParser() ExprParser {
	return &ExprParserImpl{}
}

const ExprFlag = -1000

func ExprTokname(c int) string {
	if c >= 1 && c-1 < len(ExprToknames) {
		if ExprToknames[c-1] != "" {
			return ExprToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func ExprStatname(s int) string {
	if s >= 0 && s < len(ExprStatenames) {
		if ExprStatenames[s] != "" {
			return ExprStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func ExprErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !ExprErrorVerbose {
		return "syntax error"
	}

	for _, e := range ExprErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + ExprTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := ExprPact[state]
	for tok := TOKSTART; tok-1 < len(ExprToknames); tok++ {
		if n := base + tok; n >= 0 && n < ExprLast && ExprChk[ExprAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if ExprDef[state] == -2 {
		i := 0
		for ExprExca[i] != -1 || ExprExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; ExprExca[i] >= 0; i += 2 {
			tok := ExprExca[i]
			if tok < TOKSTART || ExprExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if ExprExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += ExprTokname(tok)
	}
	return res
}

func Exprlex1(lex ExprLexer, lval *ExprSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = ExprTok1[0]
		goto out
	}
	if char < len(ExprTok1) {
		token = ExprTok1[char]
		goto out
	}
	if char >= ExprPrivate {
		if char < ExprPrivate+len(ExprTok2) {
			token = ExprTok2[char-ExprPrivate]
			goto out
		}
	}
	for i := 0; i < len(ExprTok3); i += 2 {
		token = ExprTok3[i+0]
		if token == char {
			token = ExprTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = ExprTok2[1] /* unknown char */
	}
	if ExprDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", ExprTokname(token), uint(char))
	}
	return char, token
}

func ExprParse(Exprlex ExprLexer) int {
	return ExprNewParser().Parse(Exprlex)
}

func (Exprrcvr *ExprParserImpl) Parse(Exprlex ExprLexer) int {
	var Exprn int
	var ExprVAL ExprSymType
	var ExprDollar []ExprSymType
	_ = ExprDollar // silence set and not used
	ExprS := Exprrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Exprstate := 0
	Exprrcvr.char = -1
	Exprtoken := -1 // Exprrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Exprstate = -1
		Exprrcvr.char = -1
		Exprtoken = -1
	}()
	Exprp := -1
	goto Exprstack

ret0:
	return 0

ret1:
	return 1

Exprstack:
	/* put a state and value onto the stack */
	if ExprDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", ExprTokname(Exprtoken), ExprStatname(Exprstate))
	}

	Exprp++
	if Exprp >= len(ExprS) {
		nyys := make([]ExprSymType, len(ExprS)*2)
		copy(nyys, ExprS)
		ExprS = nyys
	}
	ExprS[Exprp] = ExprVAL
	ExprS[Exprp].yys = Exprstate

Exprnewstate:
	Exprn = ExprPact[Exprstate]
	if Exprn <= ExprFlag {
		goto Exprdefault /* simple state */
	}
	if Exprrcvr.char < 0 {
		Exprrcvr.char, Exprtoken = Exprlex1(Exprlex, &Exprrcvr.lval)
	}
	Exprn += Exprtoken
	if Exprn < 0 || Exprn >= ExprLast {
		goto Exprdefault
	}
	Exprn = ExprAct[Exprn]
	if ExprChk[Exprn] == Exprtoken { /* valid shift */
		Exprrcvr.char = -1
		Exprtoken = -1
		ExprVAL = Exprrcvr.lval
		Exprstate = Exprn
		if Errflag > 0 {
			Errflag--
		}
		goto Exprstack
	}

Exprdefault:
	/* default state action */
	Exprn = ExprDef[Exprstate]
	if Exprn == -2 {
		if Exprrcvr.char < 0 {
			Exprrcvr.char, Exprtoken = Exprlex1(Exprlex, &Exprrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if ExprExca[xi+0] == -1 && ExprExca[xi+1] == Exprstate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Exprn = ExprExca[xi+0]
			if Exprn < 0 || Exprn == Exprtoken {
				break
			}
		}
		Exprn = ExprExca[xi+1]
		if Exprn < 0 {
			goto ret0
		}
	}
	if Exprn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Exprlex.Error(ExprErrorMessage(Exprstate, Exprtoken))
			Nerrs++
			if ExprDebug >= 1 {
				__yyfmt__.Printf("%s", ExprStatname(Exprstate))
				__yyfmt__.Printf(" saw %s\n", ExprTokname(Exprtoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Exprp >= 0 {
				Exprn = ExprPact[ExprS[Exprp].yys] + ExprErrCode
				if Exprn >= 0 && Exprn < ExprLast {
					Exprstate = ExprAct[Exprn] /* simulate a shift of "error" */
					if ExprChk[Exprstate] == ExprErrCode {
						goto Exprstack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if ExprDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", ExprS[Exprp].yys)
				}
				Exprp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if ExprDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", ExprTokname(Exprtoken))
			}
			if Exprtoken == ExprEofCode {
				goto ret1
			}
			Exprrcvr.char = -1
			Exprtoken = -1
			goto Exprnewstate /* try again in the same state */
		}
	}

	/* reduction by production Exprn */
	if ExprDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Exprn, ExprStatname(Exprstate))
	}

	Exprnt := Exprn
	Exprpt := Exprp
	_ = Exprpt // guard against "declared and not used"

	Exprp -= ExprR2[Exprn]
	// Exprp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Exprp+1 >= len(ExprS) {
		nyys := make([]ExprSymType, len(ExprS)*2)
		copy(nyys, ExprS)
		ExprS = nyys
	}
	ExprVAL = ExprS[Exprp+1]

	/* consult goto table to find next state */
	Exprn = ExprR1[Exprn]
	Exprg := ExprPgo[Exprn]
	Exprj := Exprg + ExprS[Exprp].yys + 1

	if Exprj >= ExprLast {
		Exprstate = ExprAct[Exprg]
	} else {
		Exprstate = ExprAct[Exprj]
		if ExprChk[Exprstate] != -Exprn {
			Exprstate = ExprAct[Exprg]
		}
	}
	// dummy call; replaced with literal code
	switch Exprnt {

	case 3:
		ExprDollar = ExprS[Exprpt-2 : Exprpt+1]
		//line calc_gacc.y:35
		{
			fmt.Printf("%v\n", ExprDollar[1].num)
		}
	case 5:
		ExprDollar = ExprS[Exprpt-3 : Exprpt+1]
		//line calc_gacc.y:40
		{
			ExprVAL.num = ExprDollar[1].num + ExprDollar[3].num
		}
	case 6:
		ExprDollar = ExprS[Exprpt-3 : Exprpt+1]
		//line calc_gacc.y:41
		{
			ExprVAL.num = ExprDollar[1].num - ExprDollar[3].num
		}
	case 7:
		ExprDollar = ExprS[Exprpt-3 : Exprpt+1]
		//line calc_gacc.y:42
		{
			ExprVAL.num = ExprDollar[1].num * ExprDollar[3].num
		}
	case 8:
		ExprDollar = ExprS[Exprpt-3 : Exprpt+1]
		//line calc_gacc.y:43
		{
			ExprVAL.num = ExprDollar[1].num / ExprDollar[3].num
		}
	case 9:
		ExprDollar = ExprS[Exprpt-3 : Exprpt+1]
		//line calc_gacc.y:44
		{
			ExprVAL.num = ExprDollar[2].num
		}
	}
	goto Exprstack /* stack new state and value */
}

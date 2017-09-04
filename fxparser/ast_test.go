package fxparser

import (
	"testing"
	"fxlex"
	"fxsymbols"
	"math/rand"
	"strings"
)


type TestScanner struct {
	scn fxlex.Scanner
}

func (testScn *TestScanner) Scan() (fxsymbols.Token, error) {
	p := rand.Float32()
	if p >= 0.6 {
		testScn.scn.Scan()
	}
	return testScn.scn.Scan()
}

func (testScn *TestScanner) Peek() (fxsymbols.Token, error) {
	return testScn.scn.Peek()
}

func (testScn *TestScanner) Line() (int) {
	return testScn.scn.Line()
}

func (testScn *TestScanner) File() (string) {
	return testScn.scn.File()
}



func newTestScn(text fxlex.Text) *TestScanner {
	scn := fxlex.NewScanner("TestScanner", text)
	return &TestScanner{scn: scn}
}

func TestDropTokens(t *testing.T) {
	input := `//macro definition
	func line(int x, int y){
					//last number in loop is the step
		iter (i := 0; x, 1){	//declares it, scope is the loop
			line(2, 3, y, 5);
		}
	}
	func line2(int x, int y){
					//last number in loop is the step
		iter (i := 0; x, 1){	//declares it, scope is the loop
			line(2, 3, y, 5);
		}
	}
	func line3(int x, int y){
					//last number in loop is the step
		iter (i := 0; x, 1){	//declares it, scope is the loop
			line(2, 3, y, 5);
		}
	}
	//macro entry
	func main(){
		iter (i := 0; 3, 1){
			rect(i, i, 3, 0xff);
		}
		iter (j := 0; 8, 2){	//loops 0 2 4 6 8
			rect(j, j, 8, 0xff);
			iter (k := kkk; juas, jas){
				// mazo bucles!
				line(4, 5, 2, 0x11000011);
				rect(i, i, 3, 0xff);
			}
		}
		line(4, 5, 2, 0x11000011);
	}
	
	`
	testScn := newTestScn(fxlex.NewText(strings.NewReader(input)))
	p := NewParser(testScn)
	p.Parse()
}

func TestFuncEq(t *testing.T) {
	p1 := NewDeclaration("p1", 1)
	p2 := NewDeclaration("p2", 2)
	p3 := NewDeclaration("p3", 1)

	params1 := []*Declaration{p1, p2, p3}
	params2 := []*Declaration{p1, p2, p3}
	params3 := []*Declaration{p1, p3, p3} 

	v1 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.Id, Lex: "lex1"}), nil)
	v2 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.Id, Lex: "lex2"} ), nil)
	v3 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3}), nil)

	args1 := []*Expr{v1, v3}
	args2 := []*Expr{v1}
	args3 := []*Expr{v1, v3 ,v2}

	s1 := NewFuncCallSentence(NewFuncCall("function1", args1))
	s2 := NewFuncCallSentence(NewFuncCall("function1", args2))
	s3 := NewFuncCallSentence(NewFuncCall("function2", args3))

	body1 := []*Sentence{s1, s2}
	body2 := []*Sentence{s1, s2}
	body3 := []*Sentence{s1, s3}

	f1 := NewFunc("function1", params1, body1)
	f2 := NewFunc("function1", params2, body2)
	// different id
	f3 := NewFunc("function3", params2, body2)
	// diferent params
	f4 := NewFunc("function1", params3, body1)
	// different body
	f5 := NewFunc("function1", params2, body3)

	if !f1.Equal(f2) {
		t.Errorf("Error: should be equal:\n%s\n%s\n", f1, f2)
	}
	if f1.Equal(f3) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", f1, f3)
	}
	if f1.Equal(f4) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", f1, f4)
	}
	if f1.Equal(f5) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", f1, f5)
	}
}

func TestLoopEq(t *testing.T) {
	cVar1 := NewAsign("controlVar1", NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.Id, Lex: "lex1"}), nil))
	cVar2 := NewAsign("controlVar1", NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.Id, Lex: "lex1"}), nil))
	cVar3 := NewAsign("controlVar1", NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3}), nil))

	endCond1 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0xFF", Val: 255}), nil)
	endCond2 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0xFF", Val: 255}), nil)
	endCond4 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3}), nil)

	incr1 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "189", Val: 189}), nil)
	incr2 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "189", Val: 189}), nil)
	incr5 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.Id, Lex: "lex1"}), nil)

	l1 := NewLoop(cVar1, endCond1, incr1, nil)
	l2 := NewLoop(cVar2, endCond2, incr2, nil)
	l3 := NewLoop(cVar3, endCond1, incr1, nil)
	l4 := NewLoop(cVar2, endCond4, incr2, nil)
	l5 := NewLoop(cVar1, endCond2, incr5, nil)

	if !l1.Equal(l2) {
		t.Errorf("Error: should be equal:\n%s\n%s\n", l1, l2)
	}
	if l1.Equal(l3) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", l1, l3)
	}
	if l1.Equal(l4) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", l1, l4)
	}
	if l1.Equal(l5) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", l1, l5)
	}
}

func TestFuncCallEq (t *testing.T) {
	v1 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.Id, Lex: "lex1"}), nil)
	v2 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.Id, Lex: "lex2"} ), nil)
	v3 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3}), nil)
	v4 := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0xFF", Val: 255}), nil)

	args1 := []*Expr{v1, v3, v4}
	args2 := []*Expr{v1, v3, v4}
	args3 := []*Expr{v1, v3 ,v2}
	args4 := []*Expr{v1, v2, v3, v4}

	f1 := NewFuncCall("function1", args1)
	f2 := NewFuncCall("function1", args2)
	f3 := NewFuncCall("function2", args3)
	f4 := NewFuncCall("function2", args4)

	if !f1.Equal(f2) {
		t.Errorf("Error: should be equal:\n%s\n%s\n", f1, f2)
	}
	if f1.Equal(f3) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", f1, f3)
	}
	if f2.Equal(f4) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", f2, f4)
	}
}

func TestAsignEq (t *testing.T) {
	val1 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3, Line: 1})
	val2 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3, Line: 1})
	val3 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "30", Val: 3, Line: 1})

	exp1 := NewExpr(fxsymbols.None, val1, nil)
	exp2 := NewExpr(fxsymbols.None, val2, nil)
	exp3 := NewExpr(fxsymbols.None, val3, nil)

	asign1 := NewAsign("x1", exp1)
	asign2 := NewAsign("x1", exp2)
	asign3 := NewAsign("x3", exp3)

	if !asign1.Equal(asign2) {
		t.Errorf("Error: should be equal:\n%s\n%s\n", asign1, asign2)
	}
	if asign1.Equal(asign3) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", asign1, asign3)
	}
}

func TestIfEq (t *testing.T) {
	val1 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3, Line: 1})
	val2 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3, Line: 1})
	val3 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "30", Val: 3, Line: 1})

	exp1 := NewExpr(fxsymbols.None, val1, nil)
	exp2 := NewExpr(fxsymbols.None, val2, nil)
	exp3 := NewExpr(fxsymbols.None, val3, nil)

	args1 := []*Expr{exp1}
	args2 := []*Expr{exp2}
	args3 := []*Expr{exp1, exp3}

	s1 := NewFuncCallSentence(NewFuncCall("function1", args1))
	s2 := NewFuncCallSentence(NewFuncCall("function1", args2))
	s3 := NewFuncCallSentence(NewFuncCall("function2", args3))

	body1 := []*Sentence{s1, s3}
	body2 := []*Sentence{s2, s3}
	body3 := []*Sentence{s1}

	if1 := NewIfStatement(exp1, body1, body1)
	if2 := NewIfStatement(exp2, body2, body2)
	if3 := NewIfStatement(exp3, body2, body2)
	if4 := NewIfStatement(exp1, body2, nil)
	if5 := NewIfStatement(exp2, body3, nil)
	if6 := NewIfStatement(exp2, body2, body3)

	if !if1.Equal(if2) {
		t.Errorf("Error: should be equal:\n%s\n%s\n", if1, if2)
	}
	if if1.Equal(if3) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", if1, if3)
	}
	if if1.Equal(if4) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", if1, if4)
	}
	if if1.Equal(if5) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", if1, if5)
	}
	if if1.Equal(if6) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", if1, if6)
	}
}

func TestExprEq (t *testing.T) {
	digit1 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "4", Val: 4})
	digit2 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.Id, Lex: "fus"})
	digit3 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3})
	digit4 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "4", Val: 4})
	digit5 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.Id, Lex: "fus"})
	digit6 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3})
	digit7 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "10", Val: 3})

	addSubExpr1 := NewExprSubExpr(NewExpr(fxsymbols.Add, digit1, digit2))
	addSubExpr2 := NewExprSubExpr(NewExpr(fxsymbols.Add, digit4, digit5))
	addSubExpr3 := NewExprSubExpr(NewExpr(fxsymbols.Add, digit4, digit6))
	addSubExpr4 := NewExprSubExpr(NewExpr(fxsymbols.Div, digit4, digit2))

	expr1 := NewExpr(fxsymbols.Mult, digit3, addSubExpr1)
	expr2 := NewExpr(fxsymbols.Mult, digit6, addSubExpr2)
	expr3 := NewExpr(fxsymbols.Mult, digit6, addSubExpr3)
	expr4 := NewExpr(fxsymbols.Mult, digit7, addSubExpr3)
	expr5 := NewExpr(fxsymbols.Mult, digit7, addSubExpr4)

	if !expr1.Equal(expr2) {
		t.Errorf("Error: should be equal:\n%s\n%s\n", expr1, expr2)
	}
	if expr1.Equal(expr3) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", expr1, expr3)
	}
	if expr1.Equal(expr4) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", expr1, expr4)
	}
	if expr1.Equal(expr5) {
		t.Errorf("Error: should not be equal:\n%s\n%s\n", expr1, expr5)
	}
}

func TestPowExpAst (t *testing.T) {
	input := `func line() { line(!3 ** !!0x1); }`

	val1 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3, Line: 1})
	val2 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0x1", Val: 1, Line: 1})
	notSub1 := NewExprSubExpr(NewExpr(fxsymbols.Not, val1, nil))
	powSub1 := NewExprSubExpr(NewExpr(fxsymbols.Pow, notSub1, val2))

	arg := []*Expr{NewExpr(fxsymbols.None, powSub1, nil)}
	body := []*Sentence{NewFuncCallSentence(NewFuncCall("line", arg))}
	line := NewFunc("line", nil, body)
	expected := NewProgram([]*Function{line})


	p := ParserFromReader("Test_AS", strings.NewReader(input))
	program, err := p.Parse()
	if err != nil {
		t.Errorf("Error: should be nil\n%s\n", err)
	}
	if !program.Equal(expected){		
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}

	input = `func line() { line(!3 ** 0x1 ** !True ** 1); }`
	
	val3 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.BoolLit, Lex: "True", Val: 1, Line: 1})
	val4 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "1", Val: 1, Line: 1})
	notSub2 := NewExprSubExpr(NewExpr(fxsymbols.Not, val3, nil))
	powSub3 := NewExprSubExpr(NewExpr(fxsymbols.Pow, notSub2, val4))
	powSub2 := NewExprSubExpr(NewExpr(fxsymbols.Pow, val2, powSub3))
	powSub1 = NewExprSubExpr(NewExpr(fxsymbols.Pow, notSub1, powSub2))

	arg = []*Expr{NewExpr(fxsymbols.None, powSub1, nil)}
	body = []*Sentence{NewFuncCallSentence(NewFuncCall("line", arg))}
	line = NewFunc("line", nil, body)
	expected = NewProgram([]*Function{line})

	p = ParserFromReader("Test_AS", strings.NewReader(input))
	program, err = p.Parse()
	if err != nil {
		t.Errorf("Error: should be nil\n%s\n", err)
	}
	if !program.Equal(expected){		
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}
}

func TestLevel2ExpAst (t *testing.T) {
	input := `func line() { line(3 * (-0x00ff**-03)); }`

	val1 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3, Line: 1})
	val2 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "-0x00ff", Val: -255, Line: 1})
	val3 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "-03", Val: -3, Line: 1})

	powSub1 := NewExprSubExpr(NewExpr(fxsymbols.Pow, val2, val3))
	parSub := NewExprSubExpr(NewExpr(fxsymbols.None, powSub1, nil))
	multSub := NewExprSubExpr(NewExpr(fxsymbols.Mult, val1, parSub))

	arg := []*Expr{NewExpr(fxsymbols.None, multSub, nil)}
	body := []*Sentence{NewFuncCallSentence(NewFuncCall("line", arg))}
	line := NewFunc("line", nil, body)
	expected := NewProgram([]*Function{line})


	p := ParserFromReader("Test_AS", strings.NewReader(input))
	program, err := p.Parse()
	if err != nil {
		t.Errorf("Error: should be nil\n%s\n", err)
		t.FailNow()
	}
	if !program.Equal(expected){		
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}

	input = `func line() { line(3 * (-0x00ff**-03) / 3 % 345 & 3**True ); }`

	val4 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "345", Val: 345, Line: 1})
	val5 := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.BoolLit, Lex: "True", Val: 1, Line: 1})

	powSub2 := NewExprSubExpr(NewExpr(fxsymbols.Pow, val1, val5))
	andSub := NewExprSubExpr(NewExpr(fxsymbols.And, val4, powSub2))
	modSub := NewExprSubExpr(NewExpr(fxsymbols.Mod, val1, andSub))
	divSub := NewExprSubExpr(NewExpr(fxsymbols.Div, parSub, modSub))
	multSub = NewExprSubExpr(NewExpr(fxsymbols.Mult, val1, divSub))

	arg = []*Expr{NewExpr(fxsymbols.None, multSub, nil)}
	body = []*Sentence{NewFuncCallSentence(NewFuncCall("line", arg))}
	line = NewFunc("line", nil, body)
	expected = NewProgram([]*Function{line})

	p = ParserFromReader("Test_AS", strings.NewReader(input))
	program, err = p.Parse()
	if err != nil {
		t.Errorf("Error: should be nil\n%s\n", err)
		t.FailNow()
	}
	if !program.Equal(expected){		
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}
}

func TestLevel3ExpAst (t *testing.T) {
	input := `func line() { line(3 + 3*3 - 3 ^ 3**!3 | 3**(3 % 3)); }`

	val := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Val: 3, Line: 1})

	modSub := NewExprSubExpr(NewExpr(fxsymbols.Mod, val, val))
	parSub := NewExprSubExpr(NewExpr(fxsymbols.None, modSub, nil))
	powSub1 := NewExprSubExpr(NewExpr(fxsymbols.Pow, val, parSub))
	notSub := NewExprSubExpr(NewExpr(fxsymbols.Not, val, nil))
	powSub2 := NewExprSubExpr(NewExpr(fxsymbols.Pow, val, notSub))
	orSub := NewExprSubExpr(NewExpr(fxsymbols.Or, powSub2, powSub1))
	xorSub := NewExprSubExpr(NewExpr(fxsymbols.Xor, val, orSub))
	multSub := NewExprSubExpr(NewExpr(fxsymbols.Mult, val, val))
	subsSub := NewExprSubExpr(NewExpr(fxsymbols.Subs, multSub, xorSub))
	addSub := NewExprSubExpr(NewExpr(fxsymbols.Add, val, subsSub))

	arg := []*Expr{NewExpr(fxsymbols.None, addSub, nil)}
	body := []*Sentence{NewFuncCallSentence(NewFuncCall("line", arg))}
	line := NewFunc("line", nil, body)
	expected := NewProgram([]*Function{line})


	p := ParserFromReader("Test_AS", strings.NewReader(input))
	program, err := p.Parse()
	if err != nil {
		t.Errorf("Error: should be nil\n%s\n", err)
		t.FailNow()
	}
	if !program.Equal(expected){		
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}
}

func TestIneqExpAst (t *testing.T) {
	input := `func line() { line(2 < 2 <= 2*(2 - 2) > 2 >= 2); }`

	val := NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "2", Val: 2, Line: 1})

	bigEqSub := NewExprSubExpr(NewExpr(fxsymbols.BigEq, val, val))
	subsSub := NewExprSubExpr(NewExpr(fxsymbols.Subs, val, val))
	parSub := NewExprSubExpr(NewExpr(fxsymbols.None, subsSub, nil))
	multSub := NewExprSubExpr(NewExpr(fxsymbols.Mult, val, parSub))
	bigSub := NewExprSubExpr(NewExpr(fxsymbols.Big, multSub, bigEqSub))
	lessEqSub := NewExprSubExpr(NewExpr(fxsymbols.LessEq, val, bigSub))
	lessSub := NewExpr(fxsymbols.Less, val, lessEqSub)

	arg := []*Expr{lessSub}
	body := []*Sentence{NewFuncCallSentence(NewFuncCall("line", arg))}
	line := NewFunc("line", nil, body)
	expected := NewProgram([]*Function{line})


	p := ParserFromReader("Test_AS", strings.NewReader(input))
	program, err := p.Parse()
	if err != nil {
		t.Errorf("Error: should be nil\n%s\n", err)
		t.FailNow()
	}
	if !program.Equal(expected){		
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}
}

func TestAst1 (t *testing.T) {
	input := `func line(int hello){
		line(2, hello);
	}
	func line2(int x, int y){
					//last number in loop is the step
		iter (y := 0; x, 1){	//declares it, scope is the loop
			line(2);
		}
	}`

	cvar := NewAsign("y", NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0", Val: 0, Line: 6}),nil))
	endCond := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.Id, Lex: "x", Line: 6}), nil)
	incr := NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Val: 1, Lex: "1", Line: 6}), nil)

	args2 := []*Expr{NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Val: 2, Lex: "2", Line: 7}),nil)}
	loopBody := []*Sentence{NewFuncCallSentence(NewFuncCall("line", args2))}

	body2 := []*Sentence{NewLoopSentence(NewLoop(cvar, endCond, incr, loopBody))}
	params2 := []*Declaration{NewDeclaration("x", 1), NewDeclaration("y", 1)}
	line2 := NewFunc("line2", params2, body2)

	args1 := []*Expr{NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "2", Val: 2, Line: 2}), nil),
					NewExpr(fxsymbols.None, NewValSubExpr(&fxsymbols.Token{Id: fxsymbols.Id, Lex: "hello", Line: 2}), nil)}

	body := []*Sentence{NewFuncCallSentence(NewFuncCall("line", args1))}
	params := []*Declaration{NewDeclaration("hello", 1)}
	line := NewFunc("line", params, body)

	expected := NewProgram([]*Function{line, line2})

	p := ParserFromReader("Test_AS", strings.NewReader(input))
	program, err := p.Parse()
	if err != nil {
		t.Errorf("Error: should be nil\n%s\n", err)
	}
	if !program.Equal(expected){
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}
}
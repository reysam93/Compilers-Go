package yfxcompiler

import (
	"testing"
	"math/rand"
	"strings"
)

type TestScanner struct {
	scn FXLexer
}

func (testScn *TestScanner) Lex (lval *FXSymType) int {
	p := rand.Float32()
	if p >= 0.6 {
		testScn.scn.Lex(lval)
	}
	return testScn.scn.Lex(lval)
}

func (testScn *TestScanner) Error(s string) {
	testScn.scn.Error(s)
}

func newTestScn(text Text) *TestScanner {
	scn := NewScanner(text, "Test_Syms_Env")
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
	testScn := newTestScn(NewText(strings.NewReader(input)))
	NewEnvStack()
	FXParse(testScn)
}

func TestFuncEq(t *testing.T) {
	p1 := NewDeclaration("p1", 1)
	p2 := NewDeclaration("p2", 2)
	p3 := NewDeclaration("p3", 1)

	params1 := []*Declaration{p1, p2, p3}
	params2 := []*Declaration{p1, p2, p3}
	params3 := []*Declaration{p1, p3, p3} 

	v1 := NewExpr("", NewValSubExpr(&Token{lex: "lex1"}), nil)
	v2 := NewExpr("", NewValSubExpr(&Token{lex: "lex2"}), nil)
	v3 := NewExpr("", NewValSubExpr(&Token{lex: "3", val: 3}), nil)

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
	cVar1 := NewAsign("controlVar1", NewExpr("", NewValSubExpr(&Token{lex: "lex1"}), nil))
	cVar2 := NewAsign("controlVar1", NewExpr("", NewValSubExpr(&Token{lex: "lex1"}), nil))
	cVar3 := NewAsign("controlVar1", NewExpr("", NewValSubExpr(&Token{lex: "3", val: 3}), nil))

	endCond1 := NewExpr("", NewValSubExpr(&Token{lex: "255", val: 255}), nil)
	endCond2 := NewExpr("", NewValSubExpr(&Token{lex: "255", val: 255}), nil)
	endCond4 := NewExpr("", NewValSubExpr(&Token{lex: "3", val: 3}), nil)

	incr1 := NewExpr("", NewValSubExpr(&Token{lex: "189", val: 189}), nil)
	incr2 := NewExpr("", NewValSubExpr(&Token{lex: "189", val: 189}), nil)
	incr5 := NewExpr("", NewValSubExpr(&Token{lex: "lex1"}), nil)

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
	v1 := NewExpr("", NewValSubExpr(&Token{lex: "lex1"}), nil)
	v2 := NewExpr("", NewValSubExpr(&Token{lex: "lex2"}), nil)
	v3 := NewExpr("", NewValSubExpr(&Token{lex: "3", val: 3}), nil)
	v4 := NewExpr("", NewValSubExpr(&Token{lex: "255", val: 255}), nil)

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
	val1 := NewValSubExpr(&Token{lex: "3", val: 3})
	val2 := NewValSubExpr(&Token{lex: "3", val: 3})
	val3 := NewValSubExpr(&Token{lex: "3", val: 3})

	exp1 := NewExpr("", val1, nil)
	exp2 := NewExpr("", val2, nil)
	exp3 := NewExpr("", val3, nil)

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
	val1 := NewValSubExpr(&Token{lex: "3", val: 3})
	val2 := NewValSubExpr(&Token{lex: "3", val: 3})
	val3 := NewValSubExpr(&Token{lex: "30", val: 30})

	exp1 := NewExpr("", val1, nil)
	exp2 := NewExpr("", val2, nil)
	exp3 := NewExpr("", val3, nil)

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
	digit1 := NewValSubExpr(&Token{lex: "4", val: 4})
	digit2 := NewValSubExpr(&Token{lex: "fus"})
	digit3 := NewValSubExpr(&Token{lex: "3", val: 3})
	digit4 := NewValSubExpr(&Token{lex: "4", val: 4})
	digit5 := NewValSubExpr(&Token{lex: "fus"})
	digit6 := NewValSubExpr(&Token{lex: "3", val: 3})
	digit7 := NewValSubExpr(&Token{lex: "3", val: 3})

	addSubExpr1 := NewExprSubExpr(NewExpr("+", digit1, digit2))
	addSubExpr2 := NewExprSubExpr(NewExpr("+", digit4, digit5))
	addSubExpr3 := NewExprSubExpr(NewExpr("+", digit4, digit6))
	addSubExpr4 := NewExprSubExpr(NewExpr("/", digit4, digit2))

	expr1 := NewExpr("*", digit3, addSubExpr1)
	expr2 := NewExpr("*", digit6, addSubExpr2)
	expr3 := NewExpr("*", digit6, addSubExpr3)
	expr4 := NewExpr("*", digit7, addSubExpr3)
	expr5 := NewExpr("*", digit7, addSubExpr4)

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
	input := `func line() { circle(!3 ** 0x1); }`

	val1 := NewValSubExpr(&Token{id: INT_LIT, lex: "3", val: 3})
	val2 := NewValSubExpr(&Token{id: INT_LIT, lex: "0x1", val: 1})
	notSub1 := NewExprSubExpr(NewExpr("!", val1, nil))
	powSub1 := NewExpr("**", notSub1, val2)

	arg := []*Expr{powSub1}
	body := []*Sentence{NewFuncCallSentence(NewFuncCall("circle", arg))}
	line := NewFunc("line", nil, body)
	expected := NewProgram([]*Function{line})

	scn := NewScanner(NewText(strings.NewReader(input)), "Test_AST")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n")
	}
	program := GetAST()

	if !program.Equal(expected){		
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}

	input = `func line() { line(!3 ** 0x1 ** !True ** 1); }`
	
	// Should use NewBoolVaSubExpr
	val3 := NewValSubExpr(&Token{id: BOOL_LIT, lex: "True", val: 1})
	val4 := NewValSubExpr(&Token{id: INT_LIT, lex: "1", val: 1})
	powSub3 := NewExprSubExpr(NewExpr("**", notSub1, val2))
	notSub2 := NewExprSubExpr(NewExpr("!", val3, nil))
	powSub2 := NewExprSubExpr(NewExpr("**", powSub3, notSub2))
	powSub1 = NewExpr("**", powSub2, val4)
	
	arg = []*Expr{powSub1}
	body = []*Sentence{NewFuncCallSentence(NewFuncCall("line", arg))}
	line = NewFunc("line", nil, body)
	expected = NewProgram([]*Function{line})

	scn = NewScanner(NewText(strings.NewReader(input)), "Test_AST")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n")
	}
	program = GetAST()
	if !program.Equal(expected){		
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}
}

// Ojo a las precedencias!!!
func TestLevel2ExpAst (t *testing.T) {
	input := `func line() { line(3 * (-0x00ff**-03)); }`
	
	val1 := NewValSubExpr(&Token{id: INT_LIT, lex: "3", val: 3})
	val2 := NewValSubExpr(&Token{id: INT_LIT, lex: "-0x00ff", val: -255})
	val3 := NewValSubExpr(&Token{id: INT_LIT, lex: "-03", val: -3})

	powSub1 := NewExprSubExpr(NewExpr("**", val2, val3))
	// Tipo par sub??
	//parSub := NewExprSubExpr(NewExpr("", powSub1, nil))
	multExpr := NewExpr("*", val1, powSub1)

	arg := []*Expr{multExpr}
	body := []*Sentence{NewFuncCallSentence(NewFuncCall("line", arg))}
	line := NewFunc("line", nil, body)
	expected := NewProgram([]*Function{line})

	scn := NewScanner(NewText(strings.NewReader(input)), "Test_AST")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n")
	}
	program := GetAST()
	if !program.Equal(expected){		
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}
	input = `func line() { line(3 * (-0x00ff**-03) / 3 % 345 & 3**True ); }`

	val4 := NewValSubExpr(&Token{id: INT_LIT, lex: "345", val: 345})
	val5 := NewValSubExpr(&Token{id: BOOL_LIT, lex: "True", val: 1})

	multSub := NewExprSubExpr(NewExpr("*", val1, powSub1))
	divSub := NewExprSubExpr(NewExpr("/", multSub, val1))
	modSub := NewExprSubExpr(NewExpr("%", divSub, val4))
	powSub2 := NewExprSubExpr(NewExpr("**", val1, val5))
	andSub := NewExpr("&", modSub, powSub2)
	
	
	

	arg = []*Expr{andSub}
	body = []*Sentence{NewFuncCallSentence(NewFuncCall("line", arg))}
	line = NewFunc("line", nil, body)
	expected = NewProgram([]*Function{line})

	scn = NewScanner(NewText(strings.NewReader(input)), "Test_AST")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n")
	}
	program = GetAST()
	if !program.Equal(expected){		
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}
}

func TestLevel3ExpAst (t *testing.T) {
	input := `func line() { line(3 + 3*3 - 3 ^ 3**!3 | 3**(3 % 3)); }`

	val := NewValSubExpr(&Token{id: INT_LIT, lex: "3", val: 3})

	multSub := NewExprSubExpr(NewExpr("*", val, val))
	addSub := NewExprSubExpr(NewExpr("+", val, multSub))
	subsSub := NewExprSubExpr(NewExpr("-", addSub, val))
	notSub := NewExprSubExpr(NewExpr("!", val, nil))
	powSub2 := NewExprSubExpr(NewExpr("**", val, notSub))
	xorSub := NewExprSubExpr(NewExpr("^", subsSub, powSub2))
	modSub := NewExprSubExpr(NewExpr("%", val, val))
	//parSub := NewExpr("", modSub, nil)
	powSub1 := NewExprSubExpr(NewExpr("**", val, modSub))
	orSub := NewExpr("|", xorSub, powSub1)

	arg := []*Expr{orSub}
	body := []*Sentence{NewFuncCallSentence(NewFuncCall("line", arg))}
	line := NewFunc("line", nil, body)
	expected := NewProgram([]*Function{line})


	scn := NewScanner(NewText(strings.NewReader(input)), "Test_AST")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n")
	}
	program := GetAST()
	if !program.Equal(expected){		
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}
}

func TestIneqExpAst (t *testing.T) {
	input := `func line() { line(2 < 2 <= 2*(2 - 2) > 2 >= 2); }`

	val := NewValSubExpr(&Token{id: INT_LIT, lex: "2", val: 2})
	lessSub := NewExprSubExpr(NewExpr("<", val, val))
	subsSub := NewExprSubExpr(NewExpr("-", val, val))
	//parSub := NewExprSubExpr(NewExpr("", subsSub, nil))
	multSub := NewExprSubExpr(NewExpr("*", val, subsSub))
	lessEqSub := NewExprSubExpr(NewExpr("<=", lessSub, multSub))
	bigSub := NewExprSubExpr(NewExpr(">", lessEqSub, val))
	bigEqSub := NewExpr(">=", bigSub, val)
	
	arg := []*Expr{bigEqSub}
	body := []*Sentence{NewFuncCallSentence(NewFuncCall("line", arg))}
	line := NewFunc("line", nil, body)
	expected := NewProgram([]*Function{line})


	scn := NewScanner(NewText(strings.NewReader(input)), "Test_AST")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n")
	}
	program := GetAST()
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

	cvar := NewAsign("y", NewExpr("", NewValSubExpr(&Token{id: INT_LIT, lex: "0", val: 0}), nil))
	endCond := NewExpr("", NewValSubExpr(&Token{id: ID, lex: "x"}), nil)
	incr := NewExpr("", NewValSubExpr(&Token{id: INT_LIT, lex: "1", val: 1}), nil)

	args2 := []*Expr{NewExpr("", NewValSubExpr(&Token{id: INT_LIT, lex: "2", val: 2}) ,nil)}
	loopBody := []*Sentence{NewFuncCallSentence(NewFuncCall("line", args2))}

	body2 := []*Sentence{NewLoopSentence(NewLoop(cvar, endCond, incr, loopBody))}
	params2 := []*Declaration{NewDeclaration("x", 1), NewDeclaration("y", 1)}
	line2 := NewFunc("line2", params2, body2)

	args1 := []*Expr{NewExpr("", NewValSubExpr(&Token{id: INT_LIT, lex: "2", val: 2}), nil),
					NewExpr("", NewValSubExpr(&Token{id: ID, lex: "hello"}), nil)}

	body := []*Sentence{NewFuncCallSentence(NewFuncCall("line", args1))}
	params := []*Declaration{NewDeclaration("hello", 1)}
	line := NewFunc("line", params, body)

	expected := NewProgram([]*Function{line, line2})

	scn := NewScanner(NewText(strings.NewReader(input)), "Test_AST")
	NewEnvStack()
	FXParse(scn)
	if nerrors != 0 {
		t.Errorf("Found %d errors\n")
	}
	program := GetAST()
	if !program.Equal(expected){
		t.Errorf("Error: Expected:\n%s\nFound:\n%s\n", expected, program)
	}
}
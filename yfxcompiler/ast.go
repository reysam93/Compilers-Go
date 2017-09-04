package yfxcompiler

import (
	"fmt"
	"strings"
)

var debugAST bool
var ast_program *Program

// Sentence Types
const (
	NullSentence SentenceType = iota
	LoopSent
	FuncCallSent
	VarDeclarSent
	AsignSent
	IfSentence
)

type SentenceType int

func (s SentenceType) String() string {
	switch s {
	case NullSentence:
		return "Null Sentence"
	case LoopSent:
		return "Loop"
	case FuncCallSent:
		return "Function Call"
	default:
		return "Unknown Sentence Type"
	}
}

const (
	// SubExpr Type
	NullSubExpr int = iota
	ExprSubExpr
	ValSubExpr
)

func DebugAST(debug bool) {
	debugAST = debug
}

func GetAST() *Program {
	return ast_program
}

type Program struct {
	Funcs []*Function
	level int 	// debug
}

func NewProgram(funcs []*Function) *Program {
	if debugAST {
		fmt.Println("New node: Program")
	}
	return &Program{Funcs: funcs}
}

func (program Program) String() string {
	str := "PROGRAM\n"
	for _, function := range program.Funcs {
		function.level = program.level + 1 
		str += function.String()
	}
	return str
}

func (p1 Program) Equal(p2 *Program) bool {
	if len(p1.Funcs) != len(p2.Funcs) {
		return false
	}
	for i, function := range p1.Funcs {
		if !function.Equal(p2.Funcs[i]) {
			return false
		}
	}
	return true
}

type Body []*Sentence

func (body Body) String(level int) string {
	str := ""
	for _, sentence := range body {
		sentence.level = level + 1
		str += sentence.String()
	}
	return str
}

func (body1 Body) Equal(body2 Body) bool {
	if len(body1) != len(body2) {
		return false
	}
	for i, sentence := range body1 {
		if !sentence.Equal(body2[i]) {
			return false
		}
	}
	return true
}

type Function struct {
	Id     string
	Params []*Declaration
	Body   Body
	level  int // debug
}

func NewFunc(id string, params []*Declaration, body Body) *Function {
	if debugAST {
		fmt.Println("New node: function_declar")
	}
	return &Function{Id: id, Params: params, Body: body}
}

func (function Function) String() string {
	params := ""
	tab := strings.Repeat("   ", function.level)

	for i, param := range function.Params {
		params += param.StringAsParam()
		if i < len(function.Params)-1 {
			params += ", "
		}
	}
	return fmt.Sprintf("%s%s (%s):\n%s\n", tab, function.Id, params,
											function.Body.String(function.level))	
}

func (f1 Function) Equal(f2 *Function) bool {
	if f1.Id != f2.Id || len(f1.Params) != len(f2.Params) {
		return false
	}

	for i, param := range f1.Params {
		if *param != *(f2.Params[i]) {
			return false
		}
	}
	return f1.Body.Equal(f2.Body)
}

type Sentence struct {
	Type     SentenceType
	FuncCall *FunctionCall
	Loop     *Loop
	Decl 	 *Declaration
	Asign 	 *Asignation
	If 		 *IfStatement
	level    int // debug
}

func NewNullSentence() *Sentence {
	if debugAST {
		fmt.Println("New node: null_sentence")
	}
	return &Sentence{Type: NullSentence}
}

func NewLoopSentence(loop *Loop) *Sentence {
	if debugAST {
		fmt.Println("New node: loop_sentence")
	}
	return &Sentence{Type: LoopSent, Loop: loop}
}

func NewFuncCallSentence(funcCall *FunctionCall) *Sentence {
	if debugAST {
		fmt.Println("New node: func_call_sentence")
	}
	return &Sentence{Type: FuncCallSent, FuncCall: funcCall}
}

func NewDeclSentence(decl *Declaration) *Sentence {
	if debugAST {
		fmt.Println("New node: decl_sentence")
	}
	return &Sentence{Type: VarDeclarSent, Decl: decl}
}

func NewAsginSentence(asign *Asignation) *Sentence {
	if debugAST {
		fmt.Println("New node: asign_sentence")
	}
	return &Sentence{Type: AsignSent, Asign: asign}
}

func NewIfSentence(if_sent *IfStatement) *Sentence {
	if debugAST {
		fmt.Println("New node: if sentence")
	}
	return &Sentence{Type: IfSentence, If: if_sent}
}

func (sentence Sentence) String() string {
	switch sentence.Type {
	case LoopSent:
		sentence.Loop.level = sentence.level + 1
		return fmt.Sprintf("%s", sentence.Loop)
	case FuncCallSent:
		sentence.FuncCall.level = sentence.level + 1
		return sentence.FuncCall.String()
	case VarDeclarSent:
		sentence.Decl.level = sentence.level + 1
		return sentence.Decl.String()
	case AsignSent:
		sentence.Asign.level = sentence.level + 1
		return sentence.Asign.String()
	case IfSentence:
		sentence.If.level = sentence.level + 1
		return sentence.If.String()
	case NullSentence:
		return "Null Sentence"
	default:
		tab := strings.Repeat("   ", sentence.level)
		return tab + "Unknown sentence"
	}
}

func (s1 Sentence) Equal(s2 *Sentence) bool {
	if s1.Type != s2.Type {
		return false
	}
	switch s1.Type {
	case LoopSent:
		return s1.Loop.Equal(s2.Loop)
	case FuncCallSent:
		return s1.FuncCall.Equal(s2.FuncCall)
	case VarDeclarSent:
		return s1.Decl == s2.Decl
	case AsignSent:
		return s1.Asign.Equal(s2.Asign)
	default:
		return false
	}
	return true
}

type FunctionCall struct {
	Id    string
	Args  []*Expr
	level int // debug
}

func NewFuncCall(id string, args []*Expr) *FunctionCall {
	if debugAST {
		fmt.Println("new node: func_call")
	}
	return &FunctionCall{Id: id, Args: args}
}

func (function FunctionCall) String() string {
	args := ""
	tab := strings.Repeat("   ", function.level)

	for i, arg := range function.Args {
		args += arg.String()
		if i < len(function.Args) -1 {
			args += ", "
		} 
	}
	return fmt.Sprintf("%s%s (%s)\n", tab, function.Id, args)
}

func (f1 FunctionCall) Equal(f2 *FunctionCall) bool{
	if f1.Id != f2.Id || len(f1.Args) != len(f2.Args) {
		return false
	}
	for i, arg := range f1.Args {
		if !arg.Equal(f2.Args[i]) {
			return false
		} 
	}
	return true
}

type Loop struct {
	ControlVar *Asignation
	EndCond    *Expr
	Increment  *Expr
	Body       Body
	level      int // debug
}

func NewLoop(control *Asignation, endCond *Expr, incr *Expr, body Body) *Loop {
	if debugAST {
		fmt.Println("New node: loop")
	}
	return &Loop{ControlVar: control, EndCond: endCond, Increment: incr, Body: body}
}

func (loop Loop) String() string {
	tab := strings.Repeat("   ", loop.level)
	str := fmt.Sprintf("%siter (%s; %s, %s):\n", tab, loop.ControlVar.StringAsLoopAsign(),
									loop.EndCond, loop.Increment)
	str += loop.Body.String(loop.level)
	return str

}

func (l1 Loop) Equal(l2 *Loop) bool {
	if !l1.ControlVar.Equal(l2.ControlVar) || len(l1.Body) != len(l2.Body) {
		return false
	}
	if !l1.EndCond.Equal(l2.EndCond) || !l1.Increment.Equal(l2.Increment) {
		return false
	}
	return l1.Body.Equal(l2.Body)
	return true
}

type IfStatement struct {
	Expr *Expr
	Body Body
	ElseBody Body
	level int // debug
}

func NewIfStatement(expr *Expr, body Body, else_body Body) *IfStatement {
	if debugAST {
		fmt.Println("New node: if_statement")
	}
	return &IfStatement{Expr: expr, Body: body, ElseBody: else_body}
}

func (if_stm IfStatement) String() string {
	tab := strings.Repeat("   ", if_stm.level)
	str := fmt.Sprintf("%sif %s:\n", tab, if_stm.Expr)
	str += if_stm.Body.String(if_stm.level)
	if if_stm.ElseBody != nil {
		str += fmt.Sprintf("%selse:\n%s", tab, if_stm.ElseBody.String(if_stm.level))
	}
	return str
}

func (if1 IfStatement) Equal(if2 *IfStatement) bool {
	if !if1.Expr.Equal(if2.Expr) {
		return false
	}
	if !if1.Body.Equal(if2.Body) {
		return false
	}
	if if1.ElseBody != nil && !if1.ElseBody.Equal(if2.ElseBody) {
		return false
	}
	return true
} 

type Asignation struct {
	Id string
	Val *Expr
	level int
}

func NewAsign(id string, expr *Expr) *Asignation {
	if debugAST {
		fmt.Println("New node: asignation")
	}
	return &Asignation{Id: id, Val: expr}
}

func (asign Asignation) String() string {
	tab := strings.Repeat("   ", asign.level)
	return fmt.Sprintf("%s%s = %s;\n", tab, asign.Id, asign.Val)
}

func (asign Asignation) StringAsLoopAsign() string {
	return fmt.Sprintf("%s = %s", asign.Id, asign.Val)
}

func (asign1 Asignation) Equal(asign2 *Asignation) bool {
	if asign1.Id != asign2.Id {
		return false
	}
	return asign1.Val.Equal(asign2.Val)
}

type Declaration struct {
	Id string
	DataType DataTypeVal
	level int
}

func NewDeclaration(id string, dataType int64) *Declaration {
	if debugAST {
		fmt.Println("New node: declaration")
	}
	return &Declaration{Id: id, DataType: DataTypeVal(dataType)}
}

func (decl Declaration) String() string {
	tab := strings.Repeat("   ", decl.level)
	return fmt.Sprintf("%s%s %s;\n", tab, decl.DataType, decl.Id)
}

func (decl Declaration) StringAsParam() string {
	return fmt.Sprintf("%s %s", decl.DataType, decl.Id)
}

type Expr struct {
	Op string	
	SubExpr1 *SubExpr 		
 	SubExpr2 *SubExpr
}

func NewExpr(op string, sub1 *SubExpr, sub2 *SubExpr) *Expr {
	e := &Expr{Op: op, SubExpr1: sub1, SubExpr2: sub2}
	if debugAST {
		fmt.Printf("New node: expr %s\n", e)
	}
	return e
}

func (expr Expr) String() string {
	switch (expr.Op){
	case "!":
		return fmt.Sprintf("!%s", expr.SubExpr1)
	case "":
		return expr.SubExpr1.String()
	default:
		return fmt.Sprintf("%s %s %s", expr.SubExpr1, expr.Op, expr.SubExpr2)
	}
}

func (exp1 Expr) Equal(exp2 *Expr) bool {
	if exp1.Op != exp2.Op {
		return false
	}
	if exp1.SubExpr1 != nil && exp2.SubExpr1 != nil && !exp1.SubExpr1.Equal(exp2.SubExpr1) {
		return false
	}
	if exp1.SubExpr2 != nil && exp2.SubExpr2 != nil && !exp1.SubExpr2.Equal(exp2.SubExpr2) {
		return false
	}
	return true
}

func NewExprSubExpr(expr *Expr) *SubExpr {
	if debugAST {
		fmt.Printf("New node: subexpr %s\n", expr)
	}
	return &SubExpr{Expr: expr}
}

type SubExpr struct {
	Val *Token
	Expr *Expr
}

func NewValSubExpr(val *Token) *SubExpr {
	if debugAST {
		fmt.Println("New node: Val_Sub_Expr")
	}
	return &SubExpr{Val: val}
}

func (sub SubExpr) String() string {
	if sub.Val != nil {
		return sub.Val.lex
	}
	return sub.Expr.String()
}

func (sub1 SubExpr) Equal(sub2 *SubExpr) bool {
	if sub1.Val != nil && sub2.Val != nil {
		if *sub1.Val != *sub2.Val {
			return false
		}
	}
	if sub1.Expr != nil && sub2.Expr != nil && !sub1.Expr.Equal(sub2.Expr) {
		return false
	}
	return true
}
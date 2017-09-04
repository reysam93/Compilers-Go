// go tool yacc -p FX yfxcompiler.y

%{
	package yfxcompiler

	import (
		//"strings"
	)

	var (
		level int
		DebugParser bool
	)
%}

%union {
	token *Token

	// AST nodes
	expr *Expr
	subExpr *SubExpr
	args []*Expr
	sentence *Sentence
	declar *Declaration
	func_call *FunctionCall
	if_stm *IfStatement
	body Body
	loop *Loop
	asign *Asignation
	parameters []*Declaration
	function *Function
	funcs []*Function
}

// Reserved words
%token 
	TYPE RECORD FUNC ITER IF ELSE

%token LOOP_EQ NONE

// Punctuation
%token
	'.' ',' '(' ')' '{' '}' '[' ']' ';'

// Data types
%token <token> DATA_TYPE BOOL_LIT INT_LIT

%token <token> ID

// Operands
%token <token>
	'<' LESS_EQ '>' BIG_EQ '+' '-' '^' '|'
	'*' '/' '%' '&' POW '!'


// Production Types -> AST nodes
%type <funcs> funcs_decls
%type <function> func_declar
%type <parameters> parameters 
%type <args> args 
%type <func_call> call_func
%type <if_stm> if_statement
%type <declar> var_declar parameter
%type <sentence> line_sentence
%type <body> body else_stm sentences
%type <loop> loop
%type <asign> loop_asign asign
%type <expr> expr val
//%type <subExpr> val


// Operator Precedence
%left '<' LESS_EQ '>' BIG_EQ
%left '+' '-' '^' '|'
%left '*' '/' '%' '&'
%left POW
%left '!'


%%
program
	: 
	{
		symbEnvs.PushEnv()
	}
	funcs_decls
	{
		program := NewProgram($2)
		symbEnvs.PopEnv()
		ast_program = program
	}

funcs_decls
	: funcs_decls func_declar
	{
		$$ = append($1, $2)
	}
	|
	{
		$$ = nil
	}
	;

func_declar
	: FUNC ID 
	{
		if symbEnvs.GetSymb($2.lex) != nil {
			Errorf("Function symbol %s already defined", $2.lex)
		}
		symbEnvs.PutFunction(*$2)
		symbEnvs.PushEnv()
	}	
	'(' parameters ')' body
	{
		$$ = NewFunc($2.lex, $5, $7)
	}	
	| FUNC error body
	{
		Errorf("Wrong function declaration")
		Errflag = 0
		symbEnvs.PushEnv()
		$$ = NewFunc("", nil, $3)
	}
	;

parameters
	: parameter
	{
		$$ = append($$, $1)
	}
	| parameters ',' parameter
	{
		$$ = append($1, $3)
	}
	|
	{
		$$ = nil
	}
	;

parameter
	: DATA_TYPE ID
	{
		if symbEnvs.GetSymb($2.lex) != nil {
			Errorf("Variable symbol %s already defined", $2.lex)
		}
		symbEnvs.PutVar(*$2)
		$$ = NewDeclaration($2.lex, $1.val)
	}
	;

body
	: '{' sentences '}'
	{
		$$ = $2
		symbEnvs.PopEnv()
	}
	| error '}'
	{
		Errorf("Wrong body")
		Errflag = 0
		symbEnvs.PopEnv()
		$$ = nil
	}
	;


sentences
	: sentences loop
	{
		sentence := NewLoopSentence($2)
		$$ = append($1, sentence)
	}
	| sentences if_statement
	{
		sentence := NewIfSentence($2)
		$$ = append($1, sentence)
	}
	| sentences line_sentence
	{
		$$ = append($1, $2)
	}
	|
	{
		$$ = nil
	}
	;

line_sentence
	: call_func ';'
	{
		$$ = NewFuncCallSentence($1)
	}
	| asign ';'
	{
		$$ = NewAsginSentence($1)
	}
	| var_declar ';'
	{
		$$ = NewDeclSentence($1)
	}
	| error ';'
	{
		Errorf("Wrong sentence")
		Errflag = 0
		$$ = NewNullSentence()
	}
	| error '}'
	{
		$$ = NewNullSentence()
	}
	;

loop
	: 
	{ 
		symbEnvs.PushEnv()
	}
	ITER '(' loop_asign ';' expr ',' expr ')' body
	{
		$$ = NewLoop($4, $6, $8, $10)
	}
	;

loop_asign
	: ID LOOP_EQ expr
	{
		if symbEnvs.GetSymb($1.lex) == nil {
			Errorf("Symbol %s is not defined", $1.lex)
		}
		$$ = NewAsign($1.lex, $3)
	}
	;

if_statement
	: IF '(' expr ')' 
	{ 
		symbEnvs.PushEnv()
	}
	body else_stm
	{
		$$ = NewIfStatement($3, $6, $7)
	}
	;

else_stm
	: ELSE 
	{ 
		symbEnvs.PushEnv()
	}
	body
	{
		$$ = $3
	}
	|
	{
		$$ = nil
	}
	;

var_declar
	: DATA_TYPE ID
	{
		if symbEnvs.GetSymb($2.lex) != nil {
			Errorf("Variable symbol %s already defined", $2.lex)
		}
		symbEnvs.PutVar(*$2)
		$$ = NewDeclaration($2.lex, $1.val)
	}
	;

call_func
	: ID '(' args ')'
	{
		if symbEnvs.GetSymb($1.lex) == nil {
			Errorf("Symbol %s is not defined", $1.lex)
		}
		$$ = NewFuncCall($1.lex, $3)
	}
	;

args
	: expr
	{
		$$ = append($$, $1)
	}
	| args ',' expr
	{
		$$ = append($1, $3)
	}
	|
	{
		$$ = nil
	}
	;

asign
	: ID '=' expr
	{
		if symbEnvs.GetSymb($1.lex) == nil {
			Errorf("Symbol %s is not defined", $1.lex)
		}
		$$ = NewAsign($1.lex, $3)
	}
	;

expr
	: expr '<' expr
	{
		$$ = createExpr($2, $1, $3)
	}
	| expr LESS_EQ expr
	{
		$$ = createExpr($2, $1, $3)
	}
	| expr '>' expr
	{
		$$ = createExpr($2, $1, $3)
	}
	| expr BIG_EQ expr
	{
		$$ = createExpr($2, $1, $3)
	}

	| expr '+' expr
	{
		$$ = createExpr($2, $1, $3)
	}
	| expr '-' expr
	{
		$$ = createExpr($2, $1, $3)
	}
	| expr '^' expr
	{
		$$ = createExpr($2, $1, $3)
	}
	| expr '|' expr
	{
		$$ = createExpr($2, $1, $3)
	}

	| expr '*' expr
	{
		$$ = createExpr($2, $1, $3)
	}
	| expr '/' expr
	{
		$$ = createExpr($2, $1, $3)
	}
	| expr '%' expr
	{
		$$ = createExpr($2, $1, $3)
	}
	| expr '&' expr
	{
		$$ = createExpr($2, $1, $3)
	}

	| expr POW expr
	{
		$$ = createExpr($2, $1, $3)
	}
	| '!' expr
	{
		var subExpr *SubExpr
		if $2.Op == "" {
			subExpr = $2.SubExpr1
		} else {
			subExpr = NewExprSubExpr($2)
		}
		$$ = NewExpr($1.lex, subExpr, nil)
	}
	| val
	{
		$$ = $1
	}
	;

val
	: ID
	{
		if symbEnvs.GetSymb($1.lex) == nil {
			Errorf("Symbol %s is not defined", $1.lex)
		}
		$$ = NewExpr("", NewValSubExpr($1), nil)
	}
	| INT_LIT
	{
		$$ = NewExpr("", NewValSubExpr($1), nil)
	}
	| BOOL_LIT
	{
		$$ = NewExpr("", NewValSubExpr($1), nil)
	}
	| '(' expr ')'
	{
		$$ = $2
	}
	;

%%

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
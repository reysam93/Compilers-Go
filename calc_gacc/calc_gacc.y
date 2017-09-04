// go tool yacc -p Expr calc_gacc.y

%{
	// +build ignore
	package main
	import (
		"fmt"
	)
%}

%union {
 num float64
 name string
}

/* tokens */
%token '('
%token ')'
%token '\n'
%token <num> NUM PI 
%token <name> FN 

%left '+' '-'
%left '*' '/'

%type <num> expr
%%

lines
	 : line
	 | lines line
	 ;

line
	 : expr '\n' { fmt.Printf("%v\n", $1) }
	 | '\n'
	 ;

expr
	 : expr '+' expr  { $$ = $1 + $3 }
	 | expr '-' expr  { $$ = $1 - $3 }
	 | expr '*' expr  { $$ = $1 * $3 }
	 | expr '/' expr  { $$ = $1 / $3 }
	 | '(' expr ')' { $$ = $2 }
	 | NUM 
	 | PI 
	 ;

%%


	

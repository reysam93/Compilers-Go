# Compilers-Go

This repository contain a compiler made on diferrent stages:
	* fxlex: the lexer for the compiler
	* fxparser: the parser for the compiler
	* fxsymbols: the table of symbols

The packet yfxcompiler is the same compiler but developed using yacc.

For creating the code with the yacc tool of go use:
	
	go tool yacc -p FX yfxcompiler.y

After the code is copied in the go working path. 
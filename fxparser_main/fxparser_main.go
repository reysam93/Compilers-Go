package main

import (
	"fmt"
	"fxparser"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: fxlex [filepath]")
	}
	file, err := os.Open(os.Args[1])
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	parser := fxparser.ParserFromReader(os.Args[1], io.Reader(file))
	//parser.DebugLex = true
	//fxparser.DebugAST = true
	//parser.DebugEnv(true)
	//parser.DebugParser = true
	prog, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Parser Done")
		fmt.Println(prog)
	}
}

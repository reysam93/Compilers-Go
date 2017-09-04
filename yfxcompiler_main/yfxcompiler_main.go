package main

import (
	"yfxcompiler"
	"log"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: fxlex [filepath]")
	}
	file := os.Args[1]
	fileDesc, err := os.Open(file)
	defer fileDesc.Close()
	if err != nil {
		log.Fatal(err)
	}
	text := yfxcompiler.NewText(fileDesc)
	scn := yfxcompiler.NewScanner(text, file)
	yfxcompiler.NewEnvStack()
	//scn.Debug = true
	//yfxcompiler.DebugAST(true)
	//yfxcompiler.EnvDebug(true)
	yfxcompiler.FXParse(scn)
	ast := yfxcompiler.GetAST()
	fmt.Println(ast)
	fmt.Printf("Errors found: %d\n", yfxcompiler.NErrors())
	os.Exit(yfxcompiler.NErrors())
}
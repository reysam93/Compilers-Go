package main

import (
	"fmt"
	"fxlex"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var input io.Reader

	input = strings.NewReader(` + - > =  < = {}.,; bool
		  =:=//COmentariO!type //otro dentro! func typefunc  cosarandom :=
	jajaja_a2  4245 3 lasuper_funcionci43ka		nada[]	
	0x01  0xA  0x00f 0xFf	0x123  -78 0holaloko00423 403 FUNCIONMAyus  -0x0002 0x2kMmm int Coord`)

	if len(os.Args) > 2 {
		log.Fatal("Usage: fxlex [filepath]\nNote: If no file is provided, a default string will be used as input.")
	}

	if len(os.Args) == 2 {
		file, err := os.Open(os.Args[1])
		defer file.Close()
		input = io.Reader(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	text := fxlex.NewText(input)
	scn := fxlex.NewScanner(text)
	for {
		t, err := scn.Scan()
		if t.Id == fxlex.EOF {
			fmt.Println("Scan done")
			break
		}
		if err != nil {
			fmt.Printf("got err %s\n", err)
			break
		}
		fmt.Printf("line: %d tok %s lex: '%v' val: %d\n", t.Line, t.Id, t.Lex, t.Val)
	}
}

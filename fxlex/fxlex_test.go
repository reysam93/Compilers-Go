package fxlex

import (
	"strings"
	"testing"
	"fxsymbols"
)

func TestPunctuation(t *testing.T) {
	input := `.,(){}[];`
	expected := []fxsymbols.TokenId{fxsymbols.Dot, fxsymbols.Coma, fxsymbols.LeftPar,
									fxsymbols.RightPar, fxsymbols.LeftBra, fxsymbols.RightBra,
									fxsymbols.LeftSqBrack, fxsymbols.RigthSqBrack, fxsymbols.Scol}

	cont := 0
	scn := NewScanner("TestPunctuation", NewText(strings.NewReader(input)))
	for {
		token, err := scn.Scan()
		if token.Id == fxsymbols.EOF {
			break
		}
		if err != nil {
			t.Errorf("Got err %s\n", err)
			break
		}
		if token.Id != expected[cont] {
			t.Errorf("Expected token %s\nFound token %s\n", expected[cont], token.Id)
		}
		cont++
	}
}

func SimpleTestOperators(t *testing.T) {
	// strings.Replace(".=", ".", ":", -1)	This has been used because the strings `:=` or ":=" broke the pairs of "" and `` in
	// sublime
	input := "+-*/**>>==<<=%|&!^:==/"
	expected := []fxsymbols.Token{fxsymbols.Token{Id: fxsymbols.Add, Lex: "+"}, fxsymbols.Token{Id: fxsymbols.Subs, Lex: "-"}, fxsymbols.Token{Id: fxsymbols.Mult, Lex: "*"},
		fxsymbols.Token{Id: fxsymbols.Div, Lex: "/"}, fxsymbols.Token{Id: fxsymbols.Pow, Lex: "**"}, fxsymbols.Token{Id: fxsymbols.Big, Lex: ">"}, fxsymbols.Token{Id: fxsymbols.BigEq, Lex: ">="}, fxsymbols.Token{Id: fxsymbols.Eq, Lex: "="},
		fxsymbols.Token{Id: fxsymbols.Less, Lex: "<"}, fxsymbols.Token{Id: fxsymbols.LessEq, Lex: "<="}, fxsymbols.Token{Id: fxsymbols.Mod, Lex: "%"}, fxsymbols.Token{Id: fxsymbols.Or, Lex: "|"}, fxsymbols.Token{Id: fxsymbols.And, Lex: "&"},
		fxsymbols.Token{Id: fxsymbols.Not, Lex: "!"}, fxsymbols.Token{Id: fxsymbols.Xor, Lex: "^"}, fxsymbols.Token{Id: fxsymbols.LoopEq, Lex: strings.Replace(".=", ".", ":", -1)}, fxsymbols.Token{Id: fxsymbols.Eq, Lex: "="},
		fxsymbols.Token{Id: fxsymbols.Div, Lex: "/"}}

	cont := 0
	scn := NewScanner("SimpleTestOperators", NewText(strings.NewReader(input)))
	for {
		token, err := scn.Scan()
		if token.Id == fxsymbols.EOF {
			break
		}
		if err != nil {
			t.Errorf("Got err %s\n", err)
			break
		}
		if token.Id != expected[cont].Id || token.Lex != expected[cont].Lex {
			t.Errorf("\nExpected token %s\nFound token %s\n", expected[cont], token)
		}
		cont++
	}
}

func TestWrongOperators(t *testing.T) {
	input := "\t\t\n\t\t* * *** :"
	expected := []fxsymbols.Token{fxsymbols.Token{Id: fxsymbols.Mult, Lex: "*", Line: 2}, fxsymbols.Token{Id: fxsymbols.Mult, Lex: "*", Line: 2}, fxsymbols.Token{Id: fxsymbols.Pow, Lex: "**", Line: 2}, fxsymbols.Token{Id: fxsymbols.Mult, Lex: "*", Line: 2}}

	cont := 0
	scn := NewScanner("TestWrongOperators", NewText(strings.NewReader(input)))
	for {
		token, err := scn.Scan()
		if token.Id == fxsymbols.EOF {
			t.Errorf("EOF founded. Should fin error")
		}
		if err != nil {
			break
		}
		if token != expected[cont] {
			t.Errorf("\nExpected %s\nFound %s\n", expected[cont], token)
		}
		cont++
	}
}

func TestDecimalNums(t *testing.T) {
	input := "3245 0 -545 \n\n-000002 000  - 3"
	expected := []fxsymbols.Token{fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3245", Line: 1, Val: 3245}, fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0", Line: 1, Val: 0},
		fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "-545", Line: 1, Val: -545}, fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "-000002", Line: 3, Val: -2},
		fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "000", Line: 3, Val: 0}, fxsymbols.Token{Id: fxsymbols.Subs, Lex: "-", Val: 0, Line: 3},
		fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Line: 3, Val: 3}}
	cont := 0
	scn := NewScanner("TestDecimalNums", NewText(strings.NewReader(input)))
	for {
		token, err := scn.Scan()
		if token.Id == fxsymbols.EOF {
			break
		}
		if err != nil {
			t.Errorf("Got err %s\n", err)
			break
		}
		if token != expected[cont] {
			t.Errorf("\nExpected %s\nFound %s\n", expected[cont], token)
		}
		cont++
	}
}

func TestHexNums(t *testing.T) {
	input := "0xFF -0x00f1\n\n 0X00 -  0x4"
	expected := []fxsymbols.Token{fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0xFF", Line: 1, Val: 255},
		fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "-0x00f1", Line: 1, Val: -241}, fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0X00", Line: 3, Val: 0},
		fxsymbols.Token{Id: fxsymbols.Subs, Lex: "-", Val: 0, Line: 3}, fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0x4", Line: 3, Val: 4}}

	cont := 0
	scn := NewScanner("TestHexNums", NewText(strings.NewReader(input)))
	for {
		token, err := scn.Scan()
		if token.Id == fxsymbols.EOF {
			break
		}
		if err != nil {
			t.Errorf("Got err %s\n", err)
			break
		}
		if token != expected[cont] {
			t.Errorf("\nExpected %s\nFound %s\n", expected[cont], token)
		}
		cont++
	}
}

func TestWrongHexNums(t *testing.T) {
	input := "0XFf -0x00f1\n\n 0x0not_num 0Xk"
	expected := []fxsymbols.Token{fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0XFf", Line: 1, Val: 255}, fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "-0x00f1", Line: 1, Val: -241},
		fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0x0", Line: 3}, fxsymbols.Token{Id: fxsymbols.Id, Lex: "not_num", Line: 3}}

	cont := 0
	scn := NewScanner("TestWrongHexNums", NewText(strings.NewReader(input)))
	for {
		token, err := scn.Scan()
		if token.Id == fxsymbols.EOF {
			t.Errorf("EOF founded. Should fin error")
			break
		}
		if err != nil {
			break
		}
		if token != expected[cont] {
			t.Errorf("\nExpected %s\nFound %s\n", expected[cont], token)
		}
		cont++
	}
}

func tTestAllNums(t *testing.T) {
	input := "3245 0 -545       0xFF -0x00f1\n\n-000002 0x4-not_num\t\t 0X00 - 3 0x4"
	expected := []fxsymbols.Token{fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3245", Line: 1, Val: 3245}, fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0", Line: 1, Val: 0},
		fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "-545", Line: 1, Val: -545}, fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0xFF", Line: 1, Val: 255},
		fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "-0x00f1", Line: 1, Val: -241}, fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "-000002", Line: 3, Val: -2},
		fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0x4", Line: 3, Val: 4}, fxsymbols.Token{Id: fxsymbols.Subs, Lex: "-", Line: 3},
		fxsymbols.Token{Id: fxsymbols.Id, Lex: "not_num", Line: 3}, fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0X00", Line: 3, Val: 0},
		fxsymbols.Token{Id: fxsymbols.Subs, Lex: "-", Val: 0, Line: 3}, fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "3", Line: 3, Val: 3},
		fxsymbols.Token{Id: fxsymbols.IntLit, Lex: "0x4", Line: 3, Val: 4}}

	cont := 0
	scn := NewScanner("tTestAllNums", NewText(strings.NewReader(input)))
	for {
		token, err := scn.Scan()
		if token.Id == fxsymbols.EOF {
			break
		}
		if err != nil {
			t.Errorf("Got err %s\n", err)
			break
		}
		if token != expected[cont] {
			t.Errorf("\nExpected %s\nFound %s\n", expected[cont], token)
		}
		cont++
	}
}

func TestSpacesAndComments(t *testing.T) {
	input := `// Com 1
	/// 324 +-* & Com 2 \t\t
	type // The only TOken!!!
	    	 bool    	 	     Coord
	// % Com 3   type int bool
	the_last34 // end comments`

	expected := []fxsymbols.Token{fxsymbols.Token{Id: fxsymbols.Type, Lex: "type", Line: 3}, fxsymbols.Token{Id: fxsymbols.DataType, Lex: "bool", Line: 4, Val: int64(Bool)},
		fxsymbols.Token{Id: fxsymbols.DataType, Lex: "Coord", Line: 4, Val: int64(Coord)}, fxsymbols.Token{Id: fxsymbols.Id, Lex: "the_last34", Line: 6}}

	cont := 0
	scn := NewScanner("TestSpacesAndComments", NewText(strings.NewReader(input)))
	for {
		token, err := scn.Scan()
		if token.Id == fxsymbols.EOF {
			break
		}
		if err != nil {
			t.Errorf("Got err %s\n", err)
			break
		}
		if token != expected[cont] {
			t.Errorf("\nExpected %s\nFound %s\n", expected[cont], token)
		}
		cont++
	}
}

func TestKeyWords(t *testing.T) {
	input := "type\nrecord\ncircle\nrect\nfunc\niter\nif\nTrue\nFalse\nint\nbool\nCoord\t\nthese_are_not reserved_words"
	expected := []fxsymbols.Token{fxsymbols.Token{Id: fxsymbols.Type, Lex: "type", Line: 1}, fxsymbols.Token{Id: fxsymbols.Record, Lex: "record", Line: 2}, fxsymbols.Token{Id: fxsymbols.Id, Lex: "circle", Line: 3}, fxsymbols.Token{Id: fxsymbols.Id, Lex: "rect", Line: 4},
		fxsymbols.Token{Id: fxsymbols.Func, Lex: "func", Line: 5}, fxsymbols.Token{Id: fxsymbols.Iter, Lex: "iter", Line: 6}, fxsymbols.Token{Id: fxsymbols.If, Lex: "if", Line: 7}, fxsymbols.Token{Id: fxsymbols.BoolLit, Lex: "True", Line: 8, Val: 1},
		fxsymbols.Token{Id: fxsymbols.BoolLit, Lex: "False", Line: 9}, fxsymbols.Token{Id: fxsymbols.DataType, Lex: "int", Line: 10, Val: int64(Int)}, fxsymbols.Token{Id: fxsymbols.DataType, Lex: "bool", Line: 11, Val: int64(Bool)},
		fxsymbols.Token{Id: fxsymbols.DataType, Lex: "Coord", Line: 12, Val: int64(Coord)}, fxsymbols.Token{Id: fxsymbols.Id, Lex: "these_are_not", Line: 13}, fxsymbols.Token{Id: fxsymbols.Id, Lex: "reserved_words", Line: 13}}

	cont := 0
	scn := NewScanner("TestKeyWords", NewText(strings.NewReader(input)))
	for {
		token, err := scn.Scan()
		if token.Id == fxsymbols.EOF {
			break
		}
		if err != nil {
			t.Errorf("Got err %s\n", err)
			break
		}
		if token != expected[cont] {
			t.Errorf("\nExpected %s\nFound %s\n", expected[cont], token)
		}
		cont++
	}
}

func TestPeek(t *testing.T) {
	input := "/ * / -"
	text := NewText(strings.NewReader(input))
	scn := NewScanner("TestPeek", text)

	tok1, _ := scn.Peek()
	tok2, _ := scn.Scan()
	if tok1 != tok2 {
		t.Error("Tok1 != Tok2")
	}
	tok2, _ = scn.Scan()
	if tok1 == tok2 {
		t.Error("Tok1 == Tok2")
	}
	tok3, _ := scn.Peek()
	if tok3 == tok2 {
		t.Error("Tok3 shoudn't equal tok2")
	}
}

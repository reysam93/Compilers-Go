package yfxcompiler

import (
	"strings"
	"testing"
	"fmt"
)

func TestPunctuation(t *testing.T) {
	input := `.,(){}[];`
	expected := []int{'.', ',', '(', ')', '{', '}', '[', ']', ';'}

	cont := 0
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Lexer")
	//scn.Debug = true
	for {
		lval := &FXSymType{}
		token := scn.Lex(lval)
		if token == 0 {
			break
		}
		if token < 0 {
			t.Errorf("Got error\n")
			break
		}
		if token != expected[cont] {
			t.Errorf("Expected token %c\nFound token %c\n", expected[cont], token)
		}
		cont++
	}
	if NErrors() != 0 {
		t.Errorf("Found %d errors\n", NErrors())
	}
}

func SimpleTestOperators(t *testing.T) {
	// strings.Replace(".=", ".", ":", -1)	This has been used because the strings `:=` or ":=" broke the pairs of "" and `` in
	// sublime
	input := "+-*/**>>==<<=%|&!^:==/"
	expected := []int{'+', '-', '*', '/', POW, '>', BIG_EQ, '=', '<', LESS_EQ,
						'%', '|', '&', '!', '^', LOOP_EQ, '=', '/'}

	cont := 0
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Lexer")
	//scn.Debug = true
	for {
		lval := &FXSymType{}
		token := scn.Lex(lval)
		if token == 0 {
			break
		}
		if token < 0 {
			t.Errorf("Got error\n")
			break
		}
		if token != expected[cont] {
			t.Errorf("Expected token %c\nFound token %c\n", expected[cont], token)
		}
		cont++
	}
	if NErrors() != 0 {
		t.Errorf("Found %d errors\n", NErrors())
	}
}

func TestWrongOperators(t *testing.T) {
	input := "\t\t\n\t\t* * *** : *"
	expected := []int{'*', '*', POW, '*', LOOP_EQ, '*'}

	cont := 0
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Lexer")
	//scn.Debug = true
	for {
		lval := &FXSymType{}
		token := scn.Lex(lval)
		if token == 0 {
			break
		}
		if token != -1 && token != expected[cont] {
			t.Errorf("Expected token %c\nFound token %c\n", expected[cont], token)
		}
		cont++
	}
	if NErrors() == 0 {
		t.Errorf("No errors found.\n", NErrors())
	}
}

func TestDecimalNums(t *testing.T) {
	input := "3245 0 -545 \n\n-000002 000  - 3"
	expected := []int64{3245, 0, -545, -000002, 000, 0, 3}

	cont := 0
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Lexer")
	// scn.Debug = true
	for {
		lval := &FXSymType{}
		token := scn.Lex(lval)
		if token == 0 {
			break
		}
		if token < 0 {
			t.Errorf("Got error\n")
			break
		}
		fmt.Println(lval.token)
		if token == INT_LIT && lval.token.val != expected[cont] {
			t.Errorf("Expected value %d\nFound value %d\n", expected[cont], lval.token.val)
		}
		cont++
	}
	if NErrors() != 0 {
		t.Errorf("Found %d errors\n", NErrors())
	}
}

func TestHexNums(t *testing.T) {
	input := "0xFF -0x00f1\n\n 0X00 -  0x4"
	expected := []int64{0xFF, -0x00f1, 0X00, '-', 0x4}

	cont := 0
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Lexer")
	//scn.Debug = true
	for {
		lval := &FXSymType{}
		token := scn.Lex(lval)
		if token == 0 {
			break
		}
		if token < 0 {
			t.Errorf("Got error\n")
			break
		}
		if token == INT_LIT && lval.token.val != expected[cont] {
			t.Errorf("Expected value %d\nFound value %d\n", expected[cont], lval.token.val)
		}
		cont++
	}
	if NErrors() != 0 {
		t.Errorf("Found %d errors\n", NErrors())
	}
}

func TestWrongHexNums(t *testing.T) {
	input := "0XFf -0x00f1\n\n 0x0not_num 0xk"
	expected := []int64{0XFf, -0x00f1, 0x0, ID, 0x0}

	cont := 0
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Lexer")
	// scn.Debug = true
	for {
		lval := &FXSymType{}
		token := scn.Lex(lval)
		if token == 0 {
			break
		}
		if token == INT_LIT && lval.token.val != expected[cont] {
			t.Errorf("Expected value %d\nFound value %d\n", expected[cont], lval.token.val)
		}
		cont++
	}
	if NErrors() == 0 {
		t.Errorf("No errors found.\n", NErrors())
	}
}

func TestAllNums(t *testing.T) {
	input := "3245 0 -545       0xFF -0x00f1\n\n-000002 0x4-not_num\t\t 0X00 - 3 0x4"
	expected := []int64{3245, 0, -545, 0xFF, -0x00f1, -000002, 0x4, '-', ID, 0X00, '-', 3, 0x4}

	cont := 0
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Lexer")
	//scn.Debug = true
	for {
		lval := &FXSymType{}
		token := scn.Lex(lval)
		if token == 0 {
			break
		}
		if token < 0 {
			t.Errorf("Got error\n")
			break
		}
		if token == INT_LIT && lval.token.val != expected[cont] {
			t.Errorf("Expected value %d\nFound value %d\n", expected[cont], lval.token.val)
		}
		cont++
	}
	if NErrors() != 0 {
		t.Errorf("Found %d errors\n", NErrors())
	}
}

func TestSpacesAndComments(t *testing.T) {
	input := `// Com 1
	/// 324 +-* & Com 2 \t\t
	type // The only TOken!!!
	    	 bool    	 	     Coord
	// % Com 3   type int bool
	the_last34 // end comments`

	expected := []int{TYPE, DATA_TYPE, DATA_TYPE, ID}
	cont := 0
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Lexer")
	// scn.Debug = true
	for {
		lval := &FXSymType{}
		token := scn.Lex(lval)
		if token == 0 {
			break
		}
		if token < 0 {
			t.Errorf("Got error\n")
			break
		}
		if token != expected[cont] {
			t.Errorf("Expected token %d\nFound token %d\n", expected[cont], token)
		}
		cont++
	}
	if NErrors() != 0 {
		t.Errorf("Found %d errors\n", NErrors())
	}
}

func TestKeyWords(t *testing.T) {
	input := "type\nrecord\ncircle\nrect\nfunc\niter\nif\nTrue\nFalse\nint\nbool\nCoord\t\nthese_are_not reserved_words"
	expected := []int{TYPE, RECORD, ID, ID, FUNC, ITER, IF, BOOL_LIT, BOOL_LIT, DATA_TYPE,
						DATA_TYPE, DATA_TYPE, ID, ID}

	cont := 0
	scn := NewScanner(NewText(strings.NewReader(input)), "Test_Lexer")
	// scn.Debug = true
	for {
		lval := &FXSymType{}
		token := scn.Lex(lval)
		if token == 0 {
			break
		}
		if token < 0 {
			t.Errorf("Got error\n")
			break
		}
		if token != expected[cont] {
			t.Errorf("Expected value %d\nFound value %d\n", expected[cont], token)
		}
		cont++
	}
	if NErrors() != 0 {
		t.Errorf("Found %d errors\n", NErrors())
	}
}
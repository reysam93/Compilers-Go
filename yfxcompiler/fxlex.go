// This is a simple lexer for the simplified fx lenguage
// Buldins are treated as "normal" function so they receive the token Id
// Samuel Rey Escdero

package yfxcompiler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

const (
	MaxErrors = 10
)

var (
	nerrors int
	file string
	line int
)

// DataTypes
const (
	NullType DataTypeVal = iota
	Int
	Bool
	Coord
)

type DataTypeVal int64

func (t DataTypeVal) String() string {
	switch t {
	case NullType:
		return "Null Type"
	case Int:
		return "int"
	case Bool:
		return "bool"
	case Coord:
		return "coord"
	default:
		return "Unknown Data Type"
	}
}

type Text interface {
	ReadRune() (rune, error)
	UnreadRune() error
}

type textBuf struct {
	in io.RuneScanner
}

func NewText(reader io.Reader) Text {
	return &textBuf{in: bufio.NewReader(reader)}
}

func (text *textBuf) ReadRune() (rune, error) {
	r, _, err := text.in.ReadRune()
	return r, err
}

func (text *textBuf) UnreadRune() error {
	return text.in.UnreadRune()
}

type scanner struct {
	in      Text
	lex     []rune
	val     int64
	Debug   bool
	lastLex string //debug
}

func NewScanner(text Text, fileName string) *scanner {
	file = fileName
	line = 1
	nerrors = 0
	debugAST = false
	return &scanner{in: text}
}

func Errorf(s string, v ...interface{}) {
	fmt.Printf("%s:%d: ", file, line)
	fmt.Printf(s, v...)
	fmt.Printf("\n")
	nerrors++
	if nerrors > MaxErrors {
		fmt.Printf("too many errors\n")
		os.Exit(1)
	}
}

func NErrors() int {
	return nerrors
}

func (scn scanner) Error(s string) {
	Errorf("%s near '%s'", s, scn.getLex())
}

func (scn *scanner) skipSpaces() error {
	for {
		r, err := scn.in.ReadRune()
		if err != nil {
			return err
		}
		if r != '\n' && r != ' ' && r != '\t' {
			scn.in.UnreadRune()
			return nil
		}
		if r == '\n' {
			line++
		}
	}
}

func (scn *scanner) addToLex(r rune) {
	scn.lex = append(scn.lex, r)
}

func (scn *scanner) putVal(num int64) {
	scn.val = num
}

func (scn scanner) getLex() string {
	return string(scn.lex)
}

func (scn *scanner) skipComments() error {
	var r rune
	var err error

	for r != '\n' {
		r, err = scn.in.ReadRune()
		if err != nil {
			return err
		}
	}
	line++
	return nil
}

func (scn *scanner) scanWord(token *Token) (int, error) {
	for {
		r, err := scn.in.ReadRune()
		if err == io.EOF || !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '_' {
			scn.in.UnreadRune()
			token.lex = scn.getLex()
			if typeVal, ok := typesWords[scn.getLex()]; ok {
				token.val = typeVal
				return DATA_TYPE, nil
			}
			if id, ok := keyWords[scn.getLex()]; ok {
				if scn.getLex() == "True" {
					token.val = 1
				}
				return id, nil
			}
			return ID, nil
		}
		if err != nil {
			return -1, err
		}
		scn.addToLex(r)
	}
}

func (scn *scanner) scanDecNum(token *Token) (int, error) {
	num := 0
	for {
		r, err := scn.in.ReadRune()
		if err == io.EOF || !unicode.IsNumber(r) {
			scn.in.UnreadRune()
			scn.putVal(int64(num))
			token.lex = scn.getLex()
			token.val = scn.val
			return INT_LIT, nil
		}
		if err != nil {
			return -1, err
		}
		scn.addToLex(r)
		num *= 10
		num += int(r - '0')
	}
}

func (scn *scanner) scanNum(token *Token) (int, error) {
	r, _ := scn.in.ReadRune()
	if r == '0' {
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		if nextRune == 'x' || nextRune == 'X' {
			scn.addToLex(nextRune)
			return scn.scanHexNum(token)
		}
	}
	scn.in.UnreadRune()
	return scn.scanDecNum(token)
}

func (scn *scanner) scanHexNum(token *Token) (int, error) {
	some := false
	val := 0
	num := 0
	for {
		r, err := scn.in.ReadRune()
		if err != nil && err != io.EOF {
			return -1, err
		}
		switch {
		case unicode.IsNumber(r):
			val = int(r - '0')
		case r >= 'a' && r <= 'f':
			val = int(r - 'a' + 10)
		case r >= 'A' && r <= 'F':
			val = int(r - 'A' + 10)
		default:
			if !some {
				return -1, fmt.Errorf("line %d: wrong input: wrong hex value\n", line)
			}
			scn.in.UnreadRune()
			scn.putVal(int64(num))
			token.val = scn.val
			token.lex = scn.getLex()
			return INT_LIT, nil
		}
		scn.addToLex(r)
		num *= 16
		num += val
		some = true
	}
	return -1, nil
}

func (scn *scanner) scanToken(lval *FXSymType) (int, error) {
	scn.lex = nil
	scn.val = 0

	r, err := scn.in.ReadRune()
	if err != nil {
		return -1, err
	}
	switch {
	case r == '.':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == ',':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '(':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == ')':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '{':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '}':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '[':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == ']':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == ';':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '+':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return '+', nil
	case r == '-':
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		scn.in.UnreadRune()
		if unicode.IsNumber(nextRune) {
			lval.token = &Token{}
			tokenId, err := scn.scanNum(lval.token)
			lval.token.val *= -1
			return tokenId, err
		}
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '*':
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		if nextRune == '*' {
			scn.addToLex(nextRune)
			lval.token = &Token{lex: scn.getLex()}
			return POW, nil
		}
		scn.in.UnreadRune()
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '/':
		nextRune, _ := scn.in.ReadRune()
		if nextRune == '/' {
			if err := scn.skipComments(); err != nil {
				return -1, err
			}
			return scn.Lex(lval), nil
		}
		scn.in.UnreadRune()
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '>':
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		if nextRune == '=' {
			scn.addToLex(nextRune)
			lval.token = &Token{lex: scn.getLex()}
			return BIG_EQ, nil
		}
		scn.in.UnreadRune()
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '<':
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		if nextRune == '=' {
			scn.addToLex(nextRune)
			lval.token = &Token{lex: scn.getLex()}
			return LESS_EQ, nil
		}
		scn.in.UnreadRune()
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '%':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '|':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '&':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '!':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '^':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == '=':
		scn.addToLex(r)
		lval.token = &Token{lex: scn.getLex()}
		return int(r), nil
	case r == ':':
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		if nextRune == '=' {
			scn.addToLex(nextRune)
			lval.token = &Token{lex: scn.getLex()}
			return LOOP_EQ, nil
		}
		scn.in.UnreadRune()
		return -1, fmt.Errorf("wrong input: '=' expected after %c. Got %c\n",
			r, nextRune)
	case unicode.IsLetter(r):
		scn.in.UnreadRune()
		lval.token = &Token{}
		return scn.scanWord(lval.token)
	case unicode.IsNumber(r):
		scn.in.UnreadRune()
		lval.token = &Token{}
		return scn.scanNum(lval.token)
	}
	return -1, fmt.Errorf("line %d: unknown input: char %c\n", line, r)
}

func (scn *scanner) Lex (lval *FXSymType) int {
	if err := scn.skipSpaces(); err != nil {
		if err != io.EOF {
			Errorf("%s", err)
			return -1
		}
		return 0
	}
	token, err := scn.scanToken(lval)
	if scn.Debug && scn.lastLex != scn.getLex(){
		fmt.Printf("token %s\n", scn.getLex())
		scn.lastLex = scn.getLex()
	}
	if err != nil {
		if err != io.EOF {
			Errorf("%s", err)
			return -1
		}
		return 0
	}
	lval.token.id = token
	return token
}

var keyWords = map[string]int{
	"type": TYPE, "record": RECORD, "func": FUNC, "iter": ITER,
	"if": IF, "else": ELSE, "True": BOOL_LIT, "False": BOOL_LIT,
}

var typesWords = map[string]int64{
	"bool": int64(Bool), "int": int64(Int),
	"Coord": int64(Coord),
}
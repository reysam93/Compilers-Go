// This is a simple lexer for the simplified fx lenguage
// Buldins are treated as "normal" function so they receive the token Id
// Samuel Rey Escdero

// Echar vistazo a : Advance data strucure

package fxlex

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
	"fxsymbols"
)

// DataTypes
const (
	NullType DataTypeConst = iota
	Int
	Bool
	Coord
)

type DataTypeConst int64

func (t DataTypeConst) String() string {
	switch(t){
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

type Scanner interface {
	Scan() (fxsymbols.Token, error)
	Peek() (fxsymbols.Token, error)
	File() (string)
	Line() (int)
}

type scanner struct {
	in    Text
	saved fxsymbols.Token
	file string
	line  int
	lex   []rune
	val   int64
}

func NewScanner(file string, text Text) Scanner {
	return &scanner{in: text, file: file, line: 1}
}

func (scn *scanner) File() (string) {
	return scn.file
}

func (scn *scanner) Line() (int) {
	return scn.line
}

func (scn *scanner) Peek() (fxsymbols.Token, error) {
	tok, err := scn.Scan()
	scn.saved = tok
	return tok, err
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
			scn.line++
		}
	}
}

func (scn *scanner) addToLex(r rune) {
	scn.lex = append(scn.lex, r)
}

func (scn *scanner) putVal(num int64) {
	scn.val = num
}

func (scn *scanner) gotTok(id fxsymbols.TokenId) fxsymbols.Token {
	t := fxsymbols.Token{
		Id:   id,
		Line: scn.line,
		Lex:  string(scn.lex),
		Val:  scn.val,
	}
	scn.lex = nil
	scn.val = 0
	return t
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
	scn.line++
	return nil
}

func (scn *scanner) scanWord() (fxsymbols.Token, error) {
	for {
		r, err := scn.in.ReadRune()
		if err == io.EOF || !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '_' {
			scn.in.UnreadRune()
			token := scn.gotTok(fxsymbols.Id)
			if typeVal, ok := typesWords[token.Lex]; ok {
				token.Id = fxsymbols.DataType
				token.Val = typeVal
			}
			if id, ok := keyWords[token.Lex]; ok {
				if token.Lex == "True" {
					token.Val = 1
				}
				token.Id = id
			}
			return token, nil
		}
		if err != nil {
			return fxsymbols.Token{}, err
		}
		scn.addToLex(r)
	}
}

func (scn *scanner) scanDecNum() (fxsymbols.Token, error) {
	num := 0
	for {
		r, err := scn.in.ReadRune()
		if err == io.EOF || !unicode.IsNumber(r) {
			scn.in.UnreadRune()
			scn.putVal(int64(num))
			return scn.gotTok(fxsymbols.IntLit), nil
		}
		if err != nil {
			return fxsymbols.Token{}, err
		}
		scn.addToLex(r)
		num *= 10
		num += int(r - '0')
	}
}

func (scn *scanner) scanNum() (fxsymbols.Token, error) {
	r, _ := scn.in.ReadRune()
	if r == '0' {
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		if nextRune == 'x' || nextRune == 'X' {
			scn.addToLex(nextRune)
			return scn.scanHexNum()
		}
	}
	scn.in.UnreadRune()
	return scn.scanDecNum()
}

func (scn *scanner) scanHexNum() (fxsymbols.Token, error) {
	some := false
	val := 0
	num := 0
	for {
		r, err := scn.in.ReadRune()
		if err != nil && err != io.EOF {
			return fxsymbols.Token{}, err
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
				return fxsymbols.Token{}, fmt.Errorf("line %d: wrong input: wrong hex value\n", scn.line)
			}
			scn.in.UnreadRune()
			scn.putVal(int64(num))
			return scn.gotTok(fxsymbols.IntLit), nil
		}
		scn.addToLex(r)
		num *= 16
		num += val
		some = true
	}
	return fxsymbols.Token{}, nil
}

func (scn *scanner) scanToken() (fxsymbols.Token, error) {
	r, err := scn.in.ReadRune()
	if err != nil {
		return fxsymbols.Token{}, err
	}
	switch {
	case r == '.':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.Dot), nil
	case r == ',':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.Coma), nil
	case r == '(':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.LeftPar), nil
	case r == ')':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.RightPar), nil
	case r == '{':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.LeftBra), nil
	case r == '}':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.RightBra), nil
	case r == '[':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.LeftSqBrack), nil
	case r == ']':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.RigthSqBrack), nil
	case r == ';':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.Scol), nil
	case r == '+':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.Add), nil
	case r == '-':
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		scn.in.UnreadRune()
		if unicode.IsNumber(nextRune) {
			token, err := scn.scanNum()
			token.Val *= -1
			return token, err
		}
		return scn.gotTok(fxsymbols.Subs), nil
	case r == '*':
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		if nextRune == '*' {
			scn.addToLex(nextRune)
			return scn.gotTok(fxsymbols.Pow), nil
		}
		scn.in.UnreadRune()
		return scn.gotTok(fxsymbols.Mult), nil
	case r == '/':
		nextRune, _ := scn.in.ReadRune()
		if nextRune == '/' {
			if err := scn.skipComments(); err != nil {
				return fxsymbols.Token{}, err
			}
			return scn.Scan()
		}
		scn.in.UnreadRune()
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.Div), nil
	case r == '>':
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		if nextRune == '=' {
			scn.addToLex(nextRune)
			return scn.gotTok(fxsymbols.BigEq), nil
		}
		scn.in.UnreadRune()
		return scn.gotTok(fxsymbols.Big), nil
	case r == '<':
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		if nextRune == '=' {
			scn.addToLex(nextRune)
			return scn.gotTok(fxsymbols.LessEq), nil
		}
		scn.in.UnreadRune()
		return scn.gotTok(fxsymbols.Less), nil
	case r == '%':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.Mod), nil
	case r == '|':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.Or), nil
	case r == '&':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.And), nil
	case r == '!':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.Not), nil
	case r == '^':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.Xor), nil
	case r == '=':
		scn.addToLex(r)
		return scn.gotTok(fxsymbols.Eq), nil
	case r == ':':
		scn.addToLex(r)
		nextRune, _ := scn.in.ReadRune()
		if nextRune == '=' {
			scn.addToLex(nextRune)
			return scn.gotTok(fxsymbols.LoopEq), nil
		}
		scn.in.UnreadRune()
		return fxsymbols.Token{}, fmt.Errorf("line %d: wrong input: '=' expected after %c. Got %c\n",
			scn.line, r, nextRune)
	case unicode.IsLetter(r):
		scn.in.UnreadRune()
		return scn.scanWord()
	case unicode.IsNumber(r):
		scn.in.UnreadRune()
		return scn.scanNum()
	}
	return fxsymbols.Token{}, fmt.Errorf("line %d: unknown input: char %c\n", scn.line, r)
}

func (scn *scanner) Scan() (fxsymbols.Token, error) {
	if scn.saved.Id != fxsymbols.None {
		tokAux := scn.saved
		scn.saved = fxsymbols.Token{}
		return tokAux, nil
	}
	err := scn.skipSpaces()
	if err == io.EOF {
		return fxsymbols.Token{Id: fxsymbols.EOF, Line: scn.line}, nil
	}
	if err != nil {
		return fxsymbols.Token{}, err
	}
	token, err := scn.scanToken()
	if err == io.EOF {
		return fxsymbols.Token{Id: fxsymbols.EOF, Line: scn.line}, nil
	}
	return token, err
}

var keyWords = map[string]fxsymbols.TokenId{
	"type": fxsymbols.Type, "record": fxsymbols.Record, "func": fxsymbols.Func, "iter": fxsymbols.Iter,
	"if": fxsymbols.If, "else": fxsymbols.Else, "True": fxsymbols.BoolLit, "False": fxsymbols.BoolLit,
}

var typesWords = map[string]int64{
	"bool": int64(Bool), "int": int64(Int), 
	"Coord": int64(Coord),
}
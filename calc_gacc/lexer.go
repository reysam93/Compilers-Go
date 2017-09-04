package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
	"bufio"
)

type Text interface {
	Get() (rune, error)
	Unget() error
}

type bufsrc struct {
	in io.RuneScanner
}

func (s *bufsrc) Get() (rune, error) {
	r, _, err := s.in.ReadRune()
	return r, err
}

func (s *bufsrc) Unget() error {
	return s.in.UnreadRune()
}

type builtin struct {
	tok int
	num float64
}

var builtins = map[string]builtin{
	"pi":  builtin{tok: PI, num: 3.1415926},
	"abs": builtin{tok: FN},
}

var file string
var line int
var debuglex bool
var nerrors int

type ExprLex interface {
	Lex(lval *ExprSymType) int
	Error(e string)
}

type lex struct {
	in  Text
	val []rune
}

func NewLex(t Text, fname string) *lex {
	file = fname
	line = 1
	return &lex{in: t}
}

func (l *lex) got(r rune) {
	l.val = append(l.val, r)
}
func (l *lex) getval() string {
	return string(l.val)
}
func (l *lex) skipBlanks() error {
	for {
		c, err := l.in.Get()
		if err != nil {
			return err
		}
		if c == '#' {
			for c != '\n' {
				if c, err = l.in.Get(); err != nil {
					return err
				}
			}
			if c == '\n' {
				line++
			}
		}
		if c == '\n' {
			line++
		}
		if c != ' ' && c != '\t' {
			l.in.Unget()
			return nil
		}
	}
}

func Errorf(s string, v ...interface{}) {
	fmt.Printf("%s:%d: ", file, line)
	fmt.Printf(s, v...)
	fmt.Printf("\n")
	nerrors++
	if nerrors > 5 {
		fmt.Printf("too many errors\n")
		os.Exit(1)
	}
}
func (l *lex) Error(s string) {
	Errorf("%s near '%s'", s, l.getval())
}

func (l *lex) Lex(lval *ExprSymType) (tid int) {
	if debuglex {
		defer func() {
			fmt.Printf("tok\n")
		}()
	}
	l.val = nil
	if err := l.skipBlanks(); err != nil {
		if err != io.EOF {
			Errorf("%s", err)
		}
		return 0
	}
	c, err := l.in.Get()
	if err != nil {
		Errorf("%s", err)
		return 0
	}
	l.got(c)
	switch {
	case c == '\n' || c == '+' || c == '*' || c == '/' || c == '(' || c == ')':
		lval.name = l.getval()
		return int(c)
	case c == '-':
		n, _ := l.in.Get()
		l.in.Unget()
		if n < '0' || n > '9' {
			return '-'
		}
		fallthrough
	case c >= '0' && c <= '9':
		for {
			c, err := l.in.Get()
			if err != nil {
				Errorf("%s", err)
				return 0
			}
			if c != '-' && c != 'e' && c != '+' && c != '.' &&
				!unicode.IsNumber(c) {
				l.in.Unget()
				break
			}
			l.got(c)
		} // id,kw
		lval.name = l.getval()
		n, err := strconv.ParseFloat(lval.name, 64)
		if err != nil {
			Errorf("%s", err)
			return 0
		}
		lval.num = n
		return NUM
	case unicode.IsLetter(c):
		for {
			c, err := l.in.Get()
			if err != nil {
				Errorf("%s", err)
				return 0
			}
			if !unicode.IsLetter(c) && !unicode.IsNumber(c) {
				l.in.Unget()
				break
			}
			l.got(c)
		} // id,kw
		lval.name = l.getval()
		b, ok := builtins[lval.name]
		if !ok {
			Errorf("unknown name '%s'", lval.name)
			return 0
		}
		lval.num = b.num
		return b.tok
	}
	Errorf("wrong input at char %c", c)
	return 0
}

func main() {
	debuglex = false
	txt := &bufsrc{in: bufio.NewReader(os.Stdin)}
	l := NewLex(txt, "stdin")
	ExprParse(l)
	os.Exit(nerrors)
}

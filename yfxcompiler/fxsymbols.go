package yfxcompiler

import (
	"fmt"
)

var (
	builtins = []Token{Token{lex: "circle"}, Token{lex: "rect"}}
	symbEnvs *EnvStack
)

type Token struct {
	id int
	lex  string
	val  int64
}

func (t Token) String() string {
	return fmt.Sprintf("Lex: %s\tVal: %d", t.lex, t.val)
}

type SymbolId int

const (
	SNone SymbolId = iota
	Svar
	Sfunc
)

type Symbol struct{
	id SymbolId
	token Token
	name string
}

type EnvStack struct {
	envs []map[string]*Symbol
	Debug bool
} 

func NewEnvStack() {
	envs := EnvStack{}
	envs.PushEnv()
	for _, builtin := range builtins {
		symbol := &Symbol{id: Sfunc, name: builtin.lex, token: builtin}
		envs.putSymb(symbol)
	}
	symbEnvs = &envs
}

func EnvDebug(debug bool) {
	symbEnvs.Debug = debug
}

func (envs *EnvStack) PushEnv() {
	envs.envs = append(envs.envs, map[string]*Symbol{})
	if envs.Debug {
		fmt.Printf("New env created. Total: %d\n", len(envs.envs))
	}
}

func (envs *EnvStack) PopEnv() {
	if len(envs.envs) == 1 {
		panic("error: trying to pop builtin env")
	}
	envs.envs = envs.envs[:len(envs.envs)-1]
	if envs.Debug {
		fmt.Printf("Env removed. Total: %d\n", len(envs.envs))
	}
}

func (envs *EnvStack) GetSymb(name string) *Symbol {
	for i := len(envs.envs)-1; i >= 0; i-- {
		if symbol, ok := envs.envs[i][name]; ok {
			return symbol
		}
	}
	if envs.Debug {
		fmt.Printf("Symbol %s not defined\n", name)
	}
	return nil
}

func (envs *EnvStack) putSymb(s *Symbol){
	envs.envs[len(envs.envs)-1][s.name] = s
}

func (envs *EnvStack) PutFunction(token Token) {
	envs.putSymb(&Symbol{name: token.lex, token: token, id: Sfunc})
	if envs.Debug {
		fmt.Printf("Function %s created\n", token.lex)
	}
}

func (envs *EnvStack) PutVar(token Token) {
	envs.putSymb(&Symbol{name: token.lex, token: token, id: Svar})
	if envs.Debug {
		fmt.Printf("Var %s created\n", token.lex)
	}
}
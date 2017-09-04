package fxsymbols

import (
	"fmt"
)

// ERRORS
var (
	ErrFuncExists = fmt.Errorf("Function already defined")/*func (name string){
		fmt.Errorf("Function %s already defined", name)
	}*/
	ErrVarExists = fmt.Errorf("Variable already defined")
)

type TokenId int

// Tokens
const (
	None TokenId = iota
	IntLit
	BoolLit
	Id
	DataType
	EOF

	// Reserved Words
	Type
	Record
	Func
	Iter
	If
	Else

	// Punctuation
	Dot
	Coma
	LeftPar
	RightPar
	LeftBra
	RightBra
	LeftSqBrack
	RigthSqBrack
	Scol

	// Operators
	Add
	Subs
	Mult
	Div
	Pow
	Big
	BigEq
	Less
	LessEq
	Mod
	Or
	And
	Not
	Xor
	LoopEq
	Eq
)

func (id TokenId) String() string {
	switch id {
	case None:
		return "None"
	case Dot:
		return "."
	case Coma:
		return ","
	case LeftPar:
		return "("
	case RightPar:
		return ")"
	case LeftBra:
		return "{"
	case RightBra:
		return "}"
	case LeftSqBrack:
		return "["
	case RigthSqBrack:
		return "]"
	case Scol:
		return ";"
	case Add:
		return "+"
	case Subs:
		return "-"
	case Mult:
		return "*"
	case Div:
		return "/"
	case Pow:
		return "**"
	case Big:
		return ">"
	case BigEq:
		return ">="
	case Less:
		return "<"
	case LessEq:
		return "<="
	case Mod:
		return "%"
	case Or:
		return "|"
	case And:
		return "&"
	case Not:
		return "!"
	case Xor:
		return "^"
	case LoopEq:
		return ":="
	case Eq:
		return "="
	case Id:
		return "Id"
	case Type:
		return "Type"
	case Record:
		return "Record"
	case Func:
		return "Func"
	case Iter:
		return "Iter"
	case If:
		return "If"
	case Else:
		return "Else"
	case IntLit:
		return "IntLit"
	case BoolLit:
		return "BoolLit"
	case DataType:
		return "DataType"
	case EOF:
		return "EOF"
	default:
		return "Unknown Token"
	}
}

type Token struct {
	Id   TokenId
	Line int
	Lex  string
	Val  int64
}

func (t Token) String() string {
	return fmt.Sprintf("%s\tLex: %s\tVal: %d\tLine: %d",
		t.Id, t.Lex, t.Line, t.Val)
}

type SymbolId int

const (
	SNone SymbolId = iota
	Svar
	Sfunc
)

type Symbol struct{
	id SymbolId
	name string
	token Token
}

type EnvStack struct {
	envs []map[string]*Symbol
	Debug bool
} 

func NewEnvStack(builtins map[string]*Token) *EnvStack {
	envs := EnvStack{}
	envs.PushEnv()
	for _, builtin := range builtins {
		symbol := &Symbol{id: Sfunc, name: builtin.Lex, token: *builtin}
		envs.putSymb(symbol)
	}
	return &envs
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

func (envs *EnvStack) PutFunction(t Token) {
	envs.putSymb(&Symbol{name: t.Lex, id: Sfunc, token: t})
	if envs.Debug {
		fmt.Printf("Function %s created\n", t.Lex)
	}
}

func (envs *EnvStack) PutVar(t Token) {
	envs.putSymb(&Symbol{name: t.Lex, id: Svar, token: t})
	if envs.Debug {
		fmt.Printf("Var %s created\n", t.Lex)
	}
}
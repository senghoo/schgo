package typ

import (
	"fmt"
	"strconv"
	"strings"
)

// type Typ int

// func (t Typ) Type() Typ {
// 	return t
// }

// const (
// 	TypIdent Typ = iota
// 	TypString
// 	TypNumber
// 	TypCons
// 	TypVector
// )

type Val interface {
	String() string
}

type Symbol string

const Nil = Symbol("<nil>")
const T = Symbol("T")

func NewSymbol(s string) Symbol {
	return Symbol(fmt.Sprintf("<%s>", s))
}

func (s Symbol) String() string {
	return string(s)
}

type String string

func NewString(s string) String {
	return String(s)
}

func (s String) String() string {
	return string(s)
}

type Int int

func NewInt(s string) Int {
	v, _ := strconv.Atoi(s)
	return NewIntFromInt(v)
}
func NewIntFromInt(i int) Int {
	return Int(i)
}

func (s Int) String() string {
	return strconv.Itoa(int(s))
}

func (s Int) Int() int {
	return int(s)
}

type Cons struct {
	Car Val
	Cdr Val
}

func MakeCons(car Val, cdr Val) *Cons {
	return &Cons{car, cdr}
}

func NewCons(val []Val) Val {
	if len(val) == 0 {
		return Nil
	}
	return &Cons{val[0], NewCons(val[1:])}
}

func (s Cons) String() string {
	return fmt.Sprintf("(%s . %s)", s.Car.String(), s.Cdr.String())
}

type Vect struct {
	Val []Val
}

func NewVect(val []Val) *Vect {
	return &Vect{val}
}

func (s Vect) String() string {
	res := make([]string, len(s.Val))
	for i, v := range s.Val {
		res[i] = v.String()
	}
	return fmt.Sprintf("[%s]", strings.Join(res, ", "))
}

type Func struct {
	name     string
	extract  bool
	function func(Val) (Val, error)
}

var lastLambda = 1

func NewFunc(name string, extract bool, f func(Val) (Val, error)) *Func {
	if name == "" {
		name = fmt.Sprintf("lambda#%d", lastLambda)
		lastLambda++
	}
	return &Func{name, extract, f}
}

func (s Func) String() string {
	return fmt.Sprint("func: %s", s.name)
}

func (s *Func) Extract() bool {
	return s.extract
}

func (s *Func) Call(v Val) (Val, error) {
	return s.function(v)
}

func (s *Func) SetName(fname string) {
	s.name = fname
}

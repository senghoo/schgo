package vm

import (
	"ligo/typ"
	"ligo/utils"
)

type env struct {
	parent *env
	vars   map[typ.Symbol]typ.Val
}

func newEnv() *env {
	env := &env{}
	env.parent = env
	env.vars = make(map[typ.Symbol]typ.Val)
	return env
}

func (e *env) newEnv(vars ...map[typ.Symbol]typ.Val) *env {
	n := &env{
		parent: e,
		vars:   make(map[typ.Symbol]typ.Val),
	}
	mergeVars(n.vars, vars...)
	return n
}

func (e *env) find(name typ.Symbol) (typ.Val, bool) {
	if v, ok := e.vars[name]; ok {
		return v, ok
	}
	if e.parent == e {
		return nil, false
	}

	return e.parent.parent.find(name)
}

func (e *env) set(name typ.Symbol, val typ.Val) {
	e.vars[name] = val
}
func (e *env) Get(n typ.Symbol) (typ.Val, bool) {
	return e.find(n)
}
func (e *env) Set(n typ.Symbol, v typ.Val) {
	e.set(n, v)
}

func (e *env) Print() string {
	var prefix string
	if e != e.parent {
		prefix = e.parent.Print()
	}
	for k, v := range e.vars {
		utils.Debugf("%s|%s: %s\n", prefix, k.String(), v.String())
	}
	return prefix + "\t"
}

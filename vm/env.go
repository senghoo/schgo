package vm

import "ligo/typ"

type env struct {
	parent *env
	vars   map[typ.Symbol]typ.Val
	dny    *env
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
	if e.dny != nil {
		e.dny.set(name, val)
		return
	}
	e.vars[name] = val
}
func (e *env) Get(n typ.Symbol) (typ.Val, bool) {
	if e.dny != nil {
		if v, ok := e.dny.find(n); ok {
			return v, ok
		}
	}
	return e.find(n)
}
func (e *env) Set(n typ.Symbol, v typ.Val) {
	e.set(n, v)
}

func (e *env) NewWith(n typ.ENV) typ.ENV {
	ne := e.newEnv()
	ne.dny = n.(*env)
	return ne
}

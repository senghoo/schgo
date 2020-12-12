package vm

import "ligo/typ"

type env struct {
	parent *env
	vars   map[string]typ.Val
}

func newEnv() *env {
	env := &env{}
	env.parent = env
	env.vars = make(map[string]typ.Val)
	return env
}

func (e *env) newEnv(vars ...map[string]typ.Val) *env {
	n := &env{
		parent: e,
		vars:   make(map[string]typ.Val),
	}
	mergeVars(n.vars, vars...)
	return n
}

func (e *env) find(name string) (typ.Val, bool) {
	if v, ok := e.vars[name]; ok {
		return v, ok
	}
	if e.parent == e {
		return nil, false
	}

	return e.parent.parent.find(name)
}

func (e *env) set(name string, val typ.Val) {
	e.vars[name] = val
}

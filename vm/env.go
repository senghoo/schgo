package vm

type env struct {
	parent *env
	vars   map[string]interface{}
}

func newEnv() *env {
	env := &env{}
	env.parent = env
	env.vars = make(map[string]interface{})
	return env
}

func (e *env) newEnv(vars ...map[string]interface{}) *env {
	n := &env{
		parent: e,
		vars:   make(map[string]interface{}),
	}
	mergeVars(n.vars, vars...)
	return n
}

func (e *env) find(name string) (interface{}, bool) {
	if v, ok := e.vars[name]; ok {
		return v, ok
	}
	if e.parent == e {
		return nil, false
	}

	return e.parent.parent.find(name)
}

func (e *env) set(name string, val interface{}) {
	e.vars[name] = val
}

package vm

import (
	"fmt"
	"ligo/parser"
	"ligo/typ"
	"ligo/utils"
)

type VM struct {
	env *env
}

func NewVM() *VM {
	vm := &VM{
		env: newEnv(),
	}
	for k, v := range basicfuncs(vm, vm.env) {
		v.SetName(k)
		vm.env.vars[typ.NewSymbol(k)] = v
	}
	return vm
}

type buildins struct {
}

func (v *VM) EvalNodes(nodes []parser.Node) (*typ.Cons, error) {
	vars := make([]typ.Val, len(nodes))
	for idx, node := range nodes {
		vars[idx] = node.Val()
	}
	cons := typ.NewCons(vars).(*typ.Cons)
	return v.EvalList(v.env, cons)
}

func (v *VM) EvalNodesL(nodes []parser.Node) (typ.Val, error) {
	vars := make([]typ.Val, len(nodes))
	for idx, node := range nodes {
		vars[idx] = node.Val()
	}
	cons := typ.NewCons(vars).(*typ.Cons)
	return v.EvalListL(v.env, cons)
}

func (v *VM) EvalList(env *env, cons *typ.Cons) (*typ.Cons, error) {
	car, err := v.Eval(env, cons.Car)
	if err != nil {
		return nil, err
	}
	var cdr typ.Val
	if c, ok := cons.Cdr.(*typ.Cons); ok {
		cdr, err = v.EvalList(env, c)
	} else {
		cdr = typ.Nil
	}
	if err != nil {
		return nil, err
	}
	return typ.MakeCons(car, cdr), nil
}

func (v *VM) EvalListL(env *env, cons *typ.Cons) (typ.Val, error) {
	cons, err := v.EvalList(env, cons)
	if err != nil {
		return typ.Nil, err
	}
	array := cons.ToArray()
	return array[len(array)-1], nil
}

func (v *VM) Eval(env *env, cons typ.Val) (typ.Val, error) {
	utils.Debugf("[VM]Eval: %#v\n", cons)
	switch vv := cons.(type) {
	case typ.Symbol:
		if vv == typ.Nil {
			return typ.Nil, nil
		}
		if vv == typ.T {
			return typ.T, nil
		}
		if val, ok := env.find(vv); ok {
			return val, nil
		} else {
			return typ.Nil, fmt.Errorf("variable %s not defined", vv.String())
		}
	case typ.String, typ.Int, *typ.Vect:
		return vv, nil
	case *typ.Cons:
		return v.call(env, vv.Car, vv.Cdr)
	}
	return nil, fmt.Errorf("unsupported operator %s", cons.String())
}

func (v *VM) call(env *env, callee typ.Val, args typ.Val) (typ.Val, error) {
	utils.Debugf("[vm]callee %#s\n", callee.String())
	utils.Debugf("[vm]args %#s\n", args.String())
	fv, err := v.Eval(env, callee)
	if err != nil {
		return nil, err
	}
	if f, ok := fv.(*typ.Func); ok {
		if f.Extract() && args != typ.Nil {
			args, err = v.EvalList(env, args.(*typ.Cons))
			if err != nil {
				return nil, err
			}
		}
		if f.IsCommand() {
			return f.CallCommand(env, args)
		} else {
			return f.Call(args)
		}
	} else {
		return nil, fmt.Errorf("got non function on call %#v", fv)
	}
	return nil, nil
}

package vm

import (
	"fmt"
	"ligo/parser"
	"ligo/typ"
	"strconv"
)

type VM struct {
	env *env
}

func NewVM() *VM {
	vm := &VM{
		env: newEnv(),
	}
	for k, v := range basicfuncs {
		vm.env.vars[k] = v
	}
	return vm
}

type buildins struct {
}

func (v *VM) Eval(nodes []parser.Node) (interface{}, error) {
	return v.eval(v.env, nodes)
}

func (v *VM) eval(env *env, nodes []parser.Node) (interface{}, error) {
	var res interface{}
	var err error
	for _, node := range nodes {
		res, err = v.evalNode(env, node)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

type function struct {
	extract bool
	env     *env
	f       func(*env, []interface{}) (interface{}, error)
}

func (f *function) call(args []interface{}) (interface{}, error) {
	return f.f(f.env, args)
}

func (v *VM) call(env *env, calleeNode parser.Node, argsNode []parser.Node) (interface{}, error) {
	fmt.Printf("callee %#v\n", calleeNode)
	fmt.Println("args")
	for idx, arg := range argsNode {
		fmt.Printf("arg %d:%#v\n", idx, arg)
	}
	fv, err := v.evalNode(env, calleeNode)
	if err != nil {
		return nil, err
	}
	if f, ok := fv.(*function); ok {
		args := make([]interface{}, len(argsNode))
		for idx, arg := range argsNode {
			var argn interface{}
			if f.extract {
				argn, err = v.evalNode(env, arg)
				if err != nil {
					return nil, err
				}
			} else {
				argn = arg
			}
			args[idx] = argn
		}
		return f.call(args)
	} else {
		return nil, fmt.Errorf("got non function on call %#v", fv)
	}
	return nil, nil
}

func (v *VM) vec(env *env, vec []parser.Node) (interface{}, error) {
	fmt.Println("vecs")
	for i, v := range vec {
		fmt.Printf("[%d]%#v\n", i, v)
	}
	return nil, nil
}

func (v *VM) evalNode(env *env, node parser.Node) (typ.Val, error) {
	switch node.Type() {
	case parser.NodeIdent:
		if val, ok := env.find(node.String()); ok {
			return val, nil
		} else {
			return nil, fmt.Errorf("variable %s not defined", node.String())
		}
	case parser.NodeString:
		return node.String(), nil
	case parser.NodeNumber:
		return strconv.Atoi(node.String())
	case parser.NodeCons:
		// c := node.(*parser.CallNode)
		// return v.call(env, c.Callee, c.Args)
		return nil, nil
	case parser.NodeVector:
		c := node.(*parser.VectorNode)
		return v.vec(env, c.Nodes)
	}
	return nil, fmt.Errorf("unsupported operator %#v", node)
}

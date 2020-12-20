package vm

import (
	"errors"
	"fmt"
	"ligo/typ"
	"ligo/utils"
	"strings"
)

func basicfuncs(vm *VM, e *env) map[string]*typ.Func {
	return map[string]*typ.Func{
		// numberic
		"+": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			var sum int
			for ; v != typ.Nil; v = v.(*typ.Cons).Cdr {
				arg := v.(*typ.Cons)
				num := arg.Car.(typ.Int)
				sum += num.Int()
			}
			return typ.NewIntFromInt(sum), nil
		}),
		"-": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			var sum int
			i := 0
			for ; v != typ.Nil; v = v.(*typ.Cons).Cdr {
				arg := v.(*typ.Cons)
				num := arg.Car.(typ.Int)
				n := num.Int()
				if i == 0 {
					sum += n
				} else {
					sum -= n
				}
				i++
			}
			return typ.NewIntFromInt(sum), nil
		}),

		"*": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			sum := 1
			for ; v != typ.Nil; v = v.(*typ.Cons).Cdr {
				arg := v.(*typ.Cons)
				num := arg.Car.(typ.Int)
				sum *= num.Int()
			}
			return typ.NewIntFromInt(sum), nil
		}),
		"/": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			sum := 1
			i := 0
			for ; v != typ.Nil; v = v.(*typ.Cons).Cdr {
				arg := v.(*typ.Cons)
				num := arg.Car.(typ.Int)
				n := num.Int()
				if i == 0 {
					sum *= n
				} else {
					sum /= n
				}
				i++
			}
			return typ.NewIntFromInt(sum), nil
		}),
		"mod": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				args := c.ToArray()
				x := args[0].(typ.Int)
				y := args[1].(typ.Int)
				return x % y, nil
			}
			return typ.Nil, nil
		}),

		//compare
		">": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			var res typ.Symbol = typ.Nil
			var last int
			if c, ok := v.(*typ.Cons); ok {
				last = c.Car.(typ.Int).Int()
				res = typ.T
				for v = c.Cdr; v != typ.Nil; v = v.(*typ.Cons).Cdr {
					arg := v.(*typ.Cons)
					num := arg.Car.(typ.Int).Int()
					if last > num {
						last = num
					} else {
						return typ.Nil, nil
					}
				}
			}
			return res, nil
		}),
		"<": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			var res typ.Symbol = typ.Nil
			var last int
			if c, ok := v.(*typ.Cons); ok {
				utils.Debugf("compare  %s\n", v.String())
				last = c.Car.(typ.Int).Int()
				res = typ.T
				for v = c.Cdr; v != typ.Nil; v = v.(*typ.Cons).Cdr {
					arg := v.(*typ.Cons)
					num := arg.Car.(typ.Int).Int()
					if last < num {
						last = num
					} else {
						return typ.Nil, nil
					}
				}
			}
			return res, nil
		}),
		">=": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			var res typ.Symbol = typ.Nil
			var last int
			if c, ok := v.(*typ.Cons); ok {
				last = c.Car.(typ.Int).Int()
				res = typ.T
				for v = c.Cdr; v != typ.Nil; v = v.(*typ.Cons).Cdr {
					arg := v.(*typ.Cons)
					num := arg.Car.(typ.Int).Int()
					if last >= num {
						last = num
					} else {
						return typ.Nil, nil
					}
				}
			}
			return res, nil
		}),
		"eq": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			var res typ.Symbol = typ.Nil
			var n int
			if c, ok := v.(*typ.Cons); ok {
				n = c.Car.(typ.Int).Int()
				res = typ.T
				for v = c.Cdr; v != typ.Nil; v = v.(*typ.Cons).Cdr {
					arg := v.(*typ.Cons)
					num := arg.Car.(typ.Int).Int()
					if n != num {
						return typ.Nil, nil
					}
				}
			}
			return res, nil
		}),

		"<=": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			var res typ.Symbol = typ.Nil
			var last int
			if c, ok := v.(*typ.Cons); ok {
				last = c.Car.(typ.Int).Int()
				res = typ.T
				for v = c.Cdr; v != typ.Nil; v = v.(*typ.Cons).Cdr {
					arg := v.(*typ.Cons)
					num := arg.Car.(typ.Int).Int()
					if last <= num {
						last = num
					} else {
						return typ.Nil, nil
					}
				}
			}
			return res, nil
		}),
		// boolean
		"and": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			var last typ.Val
			if v == typ.Nil {
				return typ.Nil, nil
			}
			for ; v != typ.Nil; v = v.(*typ.Cons).Cdr {
				arg := v.(*typ.Cons)
				last = arg.Car
				if arg.Car == typ.Nil {
					return typ.Nil, nil
				}
			}
			return last, nil
		}),
		"not": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if v == typ.Nil {
				return typ.T, nil
			}
			for ; v != typ.Nil; v = v.(*typ.Cons).Cdr {
				arg := v.(*typ.Cons)
				if arg.Car != typ.Nil {
					return typ.Nil, nil
				}
			}
			return typ.T, nil
		}),

		"or": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if v == typ.Nil {
				return typ.Nil, nil
			}
			for ; v != typ.Nil; v = v.(*typ.Cons).Cdr {
				arg := v.(*typ.Cons)
				if arg.Car != typ.Nil {
					return arg.Car, nil
				}
			}
			return typ.Nil, nil
		}),

		"car": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				if c1, ok := c.Car.(*typ.Cons); ok {
					return c1.Car, nil
				}
			}
			return typ.Nil, nil
		}),
		"cdr": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				if c1, ok := c.Car.(*typ.Cons); ok {
					return c1.Cdr, nil
				}
			}
			return typ.Nil, nil
		}),

		"cons": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				args := c.ToArray()
				if len(args) != 2 {
					return typ.Nil, errors.New("cons must have 2 args")
				}
				return typ.MakeCons(args[0], args[1]), nil
			}
			return typ.Nil, nil
		}),
		"quote": typ.NewCommand(e, "", false, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				return c.Car, nil
			}
			return typ.Nil, nil
		}),
		"if": typ.NewCommand(e, "", false, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			le.(*env).Print()
			if c, ok := v.(*typ.Cons); ok {
				args := c.ToArray()
				cond := args[0]
				utils.Debugf("cond %s\n", cond.String())
				condR, err := vm.Eval(le.(*env), cond)
				utils.Debugf("condr %s\n", condR.String())
				if err != nil {
					return nil, err
				}
				if condR != typ.Nil && len(args) >= 2 {
					return vm.Eval(le.(*env), args[1])
				} else if len(args) == 3 {
					return vm.Eval(le.(*env), args[2])
				} else {
					return typ.Nil, nil
				}
			}
			return typ.Nil, nil
		}),
		"cond": typ.NewCommand(e, "", false, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				conds := c.ToArray()
				for _, cond := range conds {
					if clause, ok := cond.(*typ.Cons); ok {
						utils.Debugf("[cond] checking %s\n", clause.String())
						check := false
						if clause.Car == typ.Symbol("<else>") {
							check = true
						} else if r, err := vm.Eval(e, clause.Car); r != typ.Nil && err == nil {
							check = true
						}
						if check {
							res, err := vm.EvalList(e, clause.Cdr.(*typ.Cons))
							if err != nil {
								return typ.Nil, err
							}
							l := res.ToArray()
							return l[len(l)-1], nil

						}
					} else {
						return typ.Nil, errors.New("unexpected clause")
					}
				}
			}
			return typ.Nil, nil
		}),
		"lambda": typ.NewCommand(e, "", false, func(lambdaRunEnv typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				ld := c.ToArray()
				if len(ld) != 2 {
					return typ.Nil, errors.New("incorrect lambda")
				}
				argsc, ok := ld[0].(*typ.Cons)
				if !ok {
					return typ.Nil, errors.New("incorrect lambda")
				}
				argN := argsc.ToArray()
				body, ok := ld[1].(*typ.Cons)
				if !ok {
					return typ.Nil, errors.New("incorrect lambda")
				}
				return typ.NewFunc(lambdaRunEnv, "", true, func(lambdaRunEnv typ.ENV, v typ.Val) (typ.Val, error) {
					utils.Debugf("[LAMBDA] lambda running\n")
					utils.Debugf("[LAMBDA] args %s\n", v.String())
					if ci, ok := v.(*typ.Cons); ok {
						args := ci.ToArray()
						utils.Debugf("[LAMBDA] args %s\n", ci.String())
						if len(argN) != len(args) {
							return typ.Nil, fmt.Errorf("require %d, but %d args %#v %#v", len(argN), len(args), argN, args)
						}

						arg := make(map[typ.Symbol]typ.Val)
						for i := range argN {
							arg[argN[i].(typ.Symbol)] = args[i]
						}
						funRunEnv := lambdaRunEnv.(*env).newEnv(arg)
						funRunEnv.Print()
						return vm.Eval(funRunEnv, body)
					}
					return typ.Nil, nil
				}), nil
			}
			return typ.Nil, nil
		}),
		"setq": typ.NewCommand(e, "", false, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				ld := c.ToArray()
				if len(ld) != 2 {
					return typ.Nil, errors.New("incorrect setq")
				}
				s, ok := ld[0].(typ.Symbol)
				if !ok {
					return typ.Nil, errors.New("setq require symbol as first arg")
				}
				val, err := vm.Eval(le.(*env), ld[1])
				if err != nil {
					return typ.Nil, err
				}
				le.Set(s, val)
				return ld[1], nil
			}
			return typ.Nil, nil
		}),

		"display": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				ld := c.ToArray()
				res := make([]string, len(ld))
				for l, v := range ld {
					res[l] = v.String()
				}
				fmt.Printf("%s\n", strings.Join(res, ", "))
			}
			return typ.Nil, nil
		}),
		"go": typ.NewCommand(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				ld := c.ToArray()
				if len(ld) < 2 {
					return typ.Nil, errors.New("incorrect go")
				}
				f := ld[0].(*typ.Func)
				args := c.Cdr
				go func() {
					_, err := f.Call(args)
					if err != nil {
						fmt.Println(err.Error())
					}
				}()
			}
			return typ.Nil, nil
		}),
		"begin": typ.NewCommand(e, "", false, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				return vm.EvalListL(le.(*env), c)
			}
			return typ.Nil, nil
		}),
		"ch": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			l := 0
			if c, ok := v.(*typ.Cons); ok {
				args := c.ToArray()
				l = int(args[0].(typ.Int))
			}
			return typ.NewChan(l), nil
		}),
		"<-": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				args := c.ToArray()
				ch := args[0].(*typ.Chan)
				return ch.Fetch(), nil

			}
			return nil, errors.New("no channel to fetch")
		}),
		"->": typ.NewFunc(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				args := c.ToArray()
				ch := args[0].(*typ.Chan)
				i := args[1].(typ.Int)
				ch.Send(i)
				return i, nil
			}
			return nil, errors.New("no channel to send")
		}),
	}

}

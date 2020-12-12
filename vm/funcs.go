package vm

import (
	"errors"
	"fmt"
	"ligo/typ"
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
		"quote": typ.NewFunc(e, "", false, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				return c.Car, nil
			}
			return typ.Nil, nil
		}),
		"if": typ.NewFunc(e, "", false, func(le typ.ENV, v typ.Val) (typ.Val, error) {

			if c, ok := v.(*typ.Cons); ok {
				args := c.ToArray()
				cond := args[0]
				condR, err := vm.Eval(e, cond)
				if err != nil {
					return nil, err
				}
				if condR != typ.Nil && len(args) >= 2 {
					return vm.Eval(e, args[1])
				} else if len(args) == 3 {
					return vm.Eval(e, args[2])
				} else {
					return typ.Nil, nil
				}
			}
			return typ.Nil, nil
		}),
		"cond": typ.NewFunc(e, "", false, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				conds := c.ToArray()
				for _, cond := range conds {
					if clause, ok := cond.(*typ.Cons); ok {
						fmt.Printf("[cond] checking %s\n", clause.String())
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
					fmt.Printf("[LAMBDA] lambda running\n")
					fmt.Printf("[LAMBDA] args %s\n", v.String())
					if ci, ok := v.(*typ.Cons); ok {
						args := ci.ToArray()
						fmt.Printf("[LAMBDA] args %s\n", ci.String())
						if len(argN) != len(args) {
							return typ.Nil, fmt.Errorf("require %d, but %d args", len(argN), len(args))
						}

						arg := make(map[typ.Symbol]typ.Val)
						for i := range argN {
							arg[argN[i].(typ.Symbol)] = args[i]
						}
						funRunEnv := lambdaRunEnv.(*env).newEnv(arg)
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
				fmt.Printf("STDOUT>>>>%s\n", strings.Join(res, ", "))
			}
			return typ.Nil, nil
		}),
		"go": typ.NewCommand(e, "", true, func(le typ.ENV, v typ.Val) (typ.Val, error) {
			if c, ok := v.(*typ.Cons); ok {
				ld := c.ToArray()
				if len(ld) != 2 {
					return typ.Nil, errors.New("incorrect go")
				}
				f := ld[0].(*typ.Func)
				args := ld[1]
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
	}

}

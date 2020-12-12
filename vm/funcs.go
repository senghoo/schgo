package vm

import (
	"fmt"
	"ligo/typ"
)

var basicfuncs = map[string]*typ.Func{
	// numberic
	"+": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
		var sum int
		for ; v != typ.Nil; v = v.(*typ.Cons).Cdr {
			arg := v.(*typ.Cons)
			num := arg.Car.(typ.Int)
			sum += num.Int()
		}
		return typ.NewIntFromInt(sum), nil
	}),
	"-": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
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

	"*": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
		sum := 1
		for ; v != typ.Nil; v = v.(*typ.Cons).Cdr {
			arg := v.(*typ.Cons)
			num := arg.Car.(typ.Int)
			sum *= num.Int()
		}
		return typ.NewIntFromInt(sum), nil
	}),
	"/": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
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
	">": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
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
	"<": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
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
	">=": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
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

	"<=": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
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
	"and": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
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

	"or": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
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

	"car": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
		if c, ok := v.(*typ.Cons); ok {
			if c1, ok := c.Car.(*typ.Cons); ok {
				return c1.Car, nil
			}
		}
		return typ.Nil, nil
	}),
	"cdr": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
		if c, ok := v.(*typ.Cons); ok {
			if c1, ok := c.Car.(*typ.Cons); ok {
				return c1.Cdr, nil
			}
		}
		return typ.Nil, nil
	}),

	"cons": typ.NewFunc("", true, func(v typ.Val) (typ.Val, error) {
		if c, ok := v.(*typ.Cons); ok {
			arg1 := c.Car
			if c2, ok := c.Car.(*typ.Cons); ok {
				nc := typ.MakeCons(arg1, c2.Cdr)
				fmt.Println("xxxxx")
				fmt.Println(nc.String())
				return nc, nil
			}
		}
		return typ.Nil, nil
	}),
	// "quote": typ.NewFunc("",false, nil, func(v typ.Val) (typ.Val, error) {
	// 	for idx, arg := range args {
	// 		fmt.Printf("arg %d%#v\n", idx, arg)
	// 	}
	// 	return nil, nil
	// }),
	// "cond": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	return nil, nil
	// }),
	// "lambda": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	return nil, nil
	// }),
	// "setq": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	return nil, nil
	// }),
	// "go": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	return nil, nil
	// }),
}

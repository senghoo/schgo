package vm

import "ligo/typ"

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
	// "-": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	var sum int
	// 	for i, arg := range args {
	// 		if i == 0 {
	// 			sum += arg.(int)
	// 		} else {
	// 			sum -= arg.(int)
	// 		}
	// 	}
	// 	return sum, nil
	// }),
	// "*": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	var sum int = 1
	// 	for _, arg := range args {
	// 		sum *= arg.(int)
	// 	}
	// 	return sum, nil
	// }),
	// "/": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	var sum int = 1
	// 	for i, arg := range args {
	// 		if i == 0 {
	// 			sum += arg.(int)
	// 		} else {
	// 			sum /= arg.(int)
	// 		}
	// 	}
	// 	return sum, nil
	// }),

	// //compare
	// ">": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	var res bool = t
	// 	var last int = args[0].(int)
	// 	for _, arg := range args[1:] {
	// 		if last > arg.(int) {
	// 			last = arg.(int)
	// 		} else {
	// 			return n, nil
	// 		}
	// 	}
	// 	return res, nil
	// }),
	// "<": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	var res bool = t
	// 	var last int = args[0].(int)
	// 	for _, arg := range args[1:] {
	// 		if last < arg.(int) {
	// 			last = arg.(int)
	// 		} else {
	// 			return n, nil
	// 		}
	// 	}
	// 	return res, nil
	// }),
	// ">=": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	var res bool = t
	// 	var last int = args[0].(int)
	// 	for _, arg := range args[1:] {
	// 		if last > arg.(int) {
	// 			last = arg.(int)
	// 		} else {
	// 			return n, nil
	// 		}
	// 	}
	// 	return res, nil
	// }),
	// "<=": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	var res bool = t
	// 	var last int = args[0].(int)
	// 	for _, arg := range args[1:] {
	// 		if last < arg.(int) {
	// 			last = arg.(int)
	// 		} else {
	// 			return n, nil
	// 		}
	// 	}
	// 	return res, nil
	// }),

	// //sim func
	// "abs": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {

	// 	var res int = args[0].(int)
	// 	if res > 0 {
	// 		return res, nil
	// 	} else {
	// 		return -res, nil
	// 	}
	// }),

	// "eq": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	var res bool = t
	// 	last := args[0]
	// 	for _, arg := range args[1:] {
	// 		if last == arg {
	// 			last = arg
	// 		} else {
	// 			return n, nil
	// 		}
	// 	}
	// 	return res, nil
	// }),

	// "car": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	c, ok := args[0].(*cons)
	// 	if ok {
	// 		return c.car, nil
	// 	} else {
	// 		return n, nil
	// 	}

	// }),
	// "cdr": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	c, ok := args[0].(*cons)
	// 	if ok {
	// 		return c.cdr, nil
	// 	} else {
	// 		return n, nil
	// 	}
	// }),
	// "cons": typ.NewFunc("",true, func(v typ.Val) (typ.Val, error) {
	// 	return &cons{args[0], args[1]}, nil
	// }),
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

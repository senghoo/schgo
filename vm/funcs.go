package vm

import "fmt"

var basicfuncs = map[string]*function{
	// numberic
	"+": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		var sum int
		for _, arg := range args {
			sum += arg.(int)
		}
		return sum, nil
	}},
	"-": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		var sum int
		for i, arg := range args {
			if i == 0 {
				sum += arg.(int)
			} else {
				sum -= arg.(int)
			}
		}
		return sum, nil
	}},
	"*": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		var sum int = 1
		for _, arg := range args {
			sum *= arg.(int)
		}
		return sum, nil
	}},
	"/": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		var sum int = 1
		for i, arg := range args {
			if i == 0 {
				sum += arg.(int)
			} else {
				sum /= arg.(int)
			}
		}
		return sum, nil
	}},

	//compare
	">": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		var res bool = t
		var last int = args[0].(int)
		for _, arg := range args[1:] {
			if last > arg.(int) {
				last = arg.(int)
			} else {
				return n, nil
			}
		}
		return res, nil
	}},
	"<": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		var res bool = t
		var last int = args[0].(int)
		for _, arg := range args[1:] {
			if last < arg.(int) {
				last = arg.(int)
			} else {
				return n, nil
			}
		}
		return res, nil
	}},
	">=": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		var res bool = t
		var last int = args[0].(int)
		for _, arg := range args[1:] {
			if last > arg.(int) {
				last = arg.(int)
			} else {
				return n, nil
			}
		}
		return res, nil
	}},
	"<=": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		var res bool = t
		var last int = args[0].(int)
		for _, arg := range args[1:] {
			if last < arg.(int) {
				last = arg.(int)
			} else {
				return n, nil
			}
		}
		return res, nil
	}},

	//sim func
	"abs": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {

		var res int = args[0].(int)
		if res > 0 {
			return res, nil
		} else {
			return -res, nil
		}
	}},

	"eq": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		var res bool = t
		last := args[0]
		for _, arg := range args[1:] {
			if last == arg {
				last = arg
			} else {
				return n, nil
			}
		}
		return res, nil
	}},

	"car": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		c, ok := args[0].(*cons)
		if ok {
			return c.car, nil
		} else {
			return n, nil
		}

	}},
	"cdr": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		c, ok := args[0].(*cons)
		if ok {
			return c.cdr, nil
		} else {
			return n, nil
		}
	}},
	"cons": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		return &cons{args[0], args[1]}, nil
	}},
	"quote": &function{false, nil, func(env *env, args []interface{}) (interface{}, error) {
		for idx, arg := range args {
			fmt.Printf("arg %d%#v\n", idx, arg)
		}
		return nil, nil
	}},
	"cond": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		return nil, nil
	}},
	"lambda": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		return nil, nil
	}},
	"setq": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		return nil, nil
	}},
	"go": &function{true, nil, func(env *env, args []interface{}) (interface{}, error) {
		return nil, nil
	}},
}

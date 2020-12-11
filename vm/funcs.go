package vm

var basicfuncs = map[string]*function{
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
}

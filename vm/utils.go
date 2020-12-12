package vm

import "ligo/typ"

func mergeVars(to map[typ.Symbol]typ.Val, from ...map[typ.Symbol]typ.Val) {
	for _, f := range from {
		for k, v := range f {
			to[k] = v
		}
	}
}

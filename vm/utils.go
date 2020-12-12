package vm

import "ligo/typ"

func mergeVars(to map[string]typ.Val, from ...map[string]typ.Val) {
	for _, f := range from {
		for k, v := range f {
			to[k] = v
		}
	}
}

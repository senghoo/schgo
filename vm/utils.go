package vm

func mergeVars(to map[string]interface{}, from ...map[string]interface{}) {
	for _, f := range from {
		for k, v := range f {
			to[k] = v
		}
	}
}

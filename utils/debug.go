package utils

import "fmt"

var debug = false

func Debugf(s string, args ...interface{}) {
	if debug {
		fmt.Printf("[DEBUG]"+s, args...)
	}
}

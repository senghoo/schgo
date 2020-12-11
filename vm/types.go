package vm

type cons struct {
	car interface{}
	cdr interface{}
}

var n *cons = nil

var t = true

type symbol string

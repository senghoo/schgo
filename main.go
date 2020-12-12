package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"ligo/lexer"
	"ligo/parser"
	"ligo/vm"
	"os"
	"time"
)

func processFile(fname string) string {
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		panic(err)
	}
	return string(b) + "\n"
}

func main() {
	vm := vm.NewVM()
	if len(os.Args) > 1 {
		cmd := processFile(os.Args[1])
		lexer.Lex(cmd)
		l := lexer.Lex(cmd)

		n := parser.Parse(l)
		ret, err := vm.EvalNodesL(n)
		if err != nil {
			fmt.Printf("ERR> %s\n", err.Error())
		} else {
			fmt.Printf("RET> %s\n", ret.String())
		}
		time.Sleep(10 * time.Second)
		return
	}

	r := bufio.NewReader(os.Stdin)
	// l := lexer.Lex("(cons (+ 1 1) (* 3 3) 1)")
	// l := lexer.Lex("x")
	// for i := l.NextItem(); i.Type != lexer.ItemEOF; i = l.NextItem() {
	// 	fmt.Printf("lex %#v\n", i)
	// }

	for {
		fmt.Print(">> ")
		line, _, _ := r.ReadLine()

		cmd := string(line)
		if cmd == "bye" {
			fmt.Println("bye bye!!")
			return
		}
		fmt.Println(cmd)
		l := lexer.Lex(cmd)

		n := parser.Parse(l)
		for i, s := range n {
			fmt.Printf("[PAR]\t%d\t: %s\n", i, s.Val().String())
		}
		ret, err := vm.EvalNodesL(n)
		if err != nil {
			fmt.Printf("ERR> %s\n", err.Error())
		} else {
			fmt.Printf("RET> %s\n", ret.String())
		}
	}
}

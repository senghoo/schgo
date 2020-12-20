package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"ligo/lexer"
	"ligo/parser"
	"ligo/utils"
	"ligo/vm"
	"os"
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
			utils.Debugf("ERR> %s\n", err.Error())
		} else {
			utils.Debugf("RET> %s\n", ret.String())
		}
		return
	}

	r := bufio.NewReader(os.Stdin)
	// l := lexer.Lex("(cons (+ 1 1) (* 3 3) 1)")
	// l := lexer.Lex("x")
	// for i := l.NextItem(); i.Type != lexer.ItemEOF; i = l.NextItem() {
	// 	utils.Debugf("lex %#v\n", i)
	// }

	fmt.Print(banner)
	fmt.Println("Welcom to Schgo!")
	fmt.Println("This is course It was done as part of the BUAA PL course in 2020.")
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
			utils.Debugf("[PAR]\t%d\t: %s\n", i, s.Val().String())
		}
		ret, err := vm.EvalNodesL(n)
		if err != nil {
			utils.Debugf("ERR> %s\n", err.Error())
		} else {
			utils.Debugf("RET> %s\n", ret.String())
		}
	}
}

var banner = " \n .-')               ('-. .-. \n ( OO ).            ( OO )  / \n (_)---\\_)   .-----. ,--. ,--.  ,----.     .-'),-----. \n /    _ |   '  .--./ |  | |  | '  .-./-') ( OO'  .-.  ' \n \\  :` `.   |  |('-. |   .|  | |  |_( O- )/   |  | |  | \n '..`''.) /_) |OO  )|       | |  | .--, \\\\_) |  |\\|  | \n .-._)   \\ ||  |`-'| |  .-.  |(|  | '. (_/  \\ |  | |  | \n \\       /(_'  '--'\\ |  | |  | |  '--'  |    `'  '-'  ' \n `-----'    `-----' `--' `--'  `------'       `-----' \n "

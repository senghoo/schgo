package parser

import (
	"fmt"
	"go/token"
	"ligo/lexer"
	"ligo/typ"
)

type Node interface {
	Type() NodeType
	String() string
	Copy() Node
	Val() typ.Val
}

type NodeType int

func (t NodeType) Type() NodeType {
	return t
}

const (
	NodeIdent NodeType = iota
	NodeString
	NodeNumber
	NodeCons
	NodeVector
)

type IdentNode struct {
	NodeType
	Ident string
}

func (node *IdentNode) Val() typ.Val {
	return typ.NewSymbol(node.Ident)
}

func (node *IdentNode) Copy() Node {
	return NewIdentNode(node.Ident)
}

func (node *IdentNode) String() string {
	if node.Ident == "nil" {
		return "()"
	}

	return node.Ident
}

type StringNode struct {
	NodeType
	Value string
}

func (node *StringNode) Val() typ.Val {
	return typ.NewString(node.Value)
}

func (node *StringNode) Copy() Node {
	return newStringNode(node.Value)
}

func (node *StringNode) String() string {
	return node.Value
}

type NumberNode struct {
	NodeType
	Value      string
	NumberType token.Token
}

func (node *NumberNode) Val() typ.Val {
	return typ.NewInt(node.Value)
}

func (node *NumberNode) Copy() Node {
	return &NumberNode{NodeType: node.Type(), Value: node.Value, NumberType: node.NumberType}
}

func (node *NumberNode) String() string {
	return node.Value
}

type VectorNode struct {
	NodeType
	Nodes []Node
}

func (node *VectorNode) Val() typ.Val {
	vals := make([]typ.Val, len(node.Nodes))
	for i, x := range node.Nodes {
		vals[i] = x.Val()
	}
	return typ.NewVect(vals)
}

func (node *VectorNode) Copy() Node {
	vect := &VectorNode{NodeType: node.Type(), Nodes: make([]Node, len(node.Nodes))}
	for i, v := range node.Nodes {
		vect.Nodes[i] = v.Copy()
	}
	return vect
}

func (node *VectorNode) String() string {
	return fmt.Sprint(node.Nodes)
}

type ConsNode struct {
	NodeType
	Nodes []Node
}

func (node *ConsNode) Copy() Node {
	n := &ConsNode{NodeType: node.Type(), Nodes: make([]Node, len(node.Nodes))}
	for i, v := range node.Nodes {
		n.Nodes[i] = v.Copy()
	}
	return n
}

func (node *ConsNode) Val() typ.Val {
	vals := make([]typ.Val, len(node.Nodes))
	for i, x := range node.Nodes {
		vals[i] = x.Val()
	}
	return typ.NewCons(vals)
}

func (node *ConsNode) Cons() *typ.Cons {
	return node.Val().(*typ.Cons)
}

func (node *ConsNode) String() string {
	return node.Val().String()
}

var nilNode = NewIdentNode("nil")

func Parse(l *lexer.Lexer) []Node {
	return parser(l, make([]Node, 0), ' ')
}

func parser(l *lexer.Lexer, tree []Node, lookingFor rune) []Node {
	for item := l.NextItem(); item.Type != lexer.ItemEOF; {
		if item.Type == lexer.ItemError {
			panic(fmt.Sprintf("Lexer error %s", item.Value))
		}
		switch t := item.Type; t {
		case lexer.ItemIdent:
			tree = append(tree, NewIdentNode(item.Value))
		case lexer.ItemString:
			tree = append(tree, newStringNode(item.Value))
		case lexer.ItemInt:
			tree = append(tree, newIntNode(item.Value))
		case lexer.ItemFloat:
			tree = append(tree, newFloatNode(item.Value))
		case lexer.ItemComplex:
			tree = append(tree, newComplexNode(item.Value))
		case lexer.ItemLeftParen:
			tree = append(tree, newConsNode(parser(l, make([]Node, 0), ')')))
		case lexer.ItemLeftVect:
			tree = append(tree, newVectNode(parser(l, make([]Node, 0), ']')))
		case lexer.ItemRightParen:
			if lookingFor != ')' {
				panic(fmt.Sprintf("unexpected \")\" [%d]", item.Pos))
			}
			return tree
		case lexer.ItemRightVect:
			if lookingFor != ']' {
				panic(fmt.Sprintf("unexpected \"]\" [%d]", item.Pos))
			}
			return tree
		case lexer.ItemError:
			println(item.Value)
		default:
			panic("Bad Item type")
		}
		item = l.NextItem()
	}

	return tree
}

func NewIdentNode(name string) *IdentNode {
	return &IdentNode{NodeType: NodeIdent, Ident: name}
}

func newStringNode(val string) *StringNode {
	return &StringNode{NodeType: NodeString, Value: val}
}

func newIntNode(val string) *NumberNode {
	return &NumberNode{NodeType: NodeNumber, Value: val, NumberType: token.INT}
}

func newFloatNode(val string) *NumberNode {
	return &NumberNode{NodeType: NodeNumber, Value: val, NumberType: token.FLOAT}
}

func newComplexNode(val string) *NumberNode {
	return &NumberNode{NodeType: NodeNumber, Value: val, NumberType: token.IMAG}
}

// We return Node here, because it could be that it's nil
func newConsNode(args []Node) Node {
	if len(args) > 0 {
		return &ConsNode{NodeType: NodeCons, Nodes: args}
	} else {
		return nilNode
	}
}

func newVectNode(content []Node) *VectorNode {
	return &VectorNode{NodeType: NodeVector, Nodes: content}
}

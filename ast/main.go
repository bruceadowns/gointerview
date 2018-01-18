package main

import "log"

type node struct {
	val int
	op1 *node
	op2 *node
}

type operation struct {
	val  int
	prec int
}

type operationStack struct {
	l []*operation
}

func (s *operationStack) push(o *operation) {
	s.l = append(s.l, o)
}

func (s *operationStack) pop() (res *operation) {
	res = s.l[len(s.l)-1]
	s.l = s.l[:len(s.l)-1]
	return res
}

func (s operationStack) top() (res *operation) {
	if len(s.l) == 0 {
		return nil
	}

	return s.l[len(s.l)-1]
}

type expressionStack struct {
	l []*node
}

func (es *expressionStack) push(n *node) {
	es.l = append(es.l, n)
}

func (es *expressionStack) pop() (res *node) {
	if len(es.l) == 0 {
		return nil
	}

	res = es.l[len(es.l)-1]
	es.l = es.l[:len(es.l)-1]
	return res
}

func (es expressionStack) top() (res *node) {
	return es.l[len(es.l)-1]
}

const (
	plusPrecedence     = 1
	minusPrecedence    = 1
	multiplyPrecedence = 2
	dividePrecedence   = 2
	modPrecedence      = 3
	paranPrecedence    = 4
)

func buildOp(r rune) (res *operation) {
	switch r {
	case '(':
		res = &operation{val: int(r), prec: paranPrecedence}
	case '+':
		res = &operation{val: int(r), prec: plusPrecedence}
	case '-':
		res = &operation{val: int(r), prec: minusPrecedence}
	case '*':
		res = &operation{val: int(r), prec: multiplyPrecedence}
	case '/':
		res = &operation{val: int(r), prec: dividePrecedence}
	case '%':
		res = &operation{val: int(r), prec: modPrecedence}
	}

	return
}

func parse(s string) *node {
	opStack := &operationStack{}
	exprStack := &expressionStack{}

	for _, v := range s {
		switch v {
		case '(':
			opStack.push(buildOp(v))
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			exprStack.push(&node{val: int(v) - '0'})
		case '%', '+', '-', '*', '/':
			op := buildOp(v)
			for t := opStack.top(); t != nil && t.prec >= op.prec; t = opStack.top() {
				op := opStack.pop()
				e2 := exprStack.pop()
				e1 := exprStack.pop()
				exprStack.push(&node{val: op.val, op1: e1, op2: e2})
			}

			opStack.push(op)
		case ')':
			for t := opStack.top(); t != nil && t.val != '('; t = opStack.top() {
				op := opStack.pop()
				e2 := exprStack.pop()
				e1 := exprStack.pop()
				exprStack.push(&node{op.val, e1, e2})
			}

			//_ = opStack.pop()
		case ' ':
		default:
		}
	}

	return exprStack.pop()
}

func (n node) print() {
	log.Print(n.val)
	if n.op1 != nil {
		n.op1.print()
	}
	if n.op2 != nil {
		n.op2.print()
	}
}

func main() {
	in := "5 * 3 + (4 + 2 % 2 * 8)"
	log.Print(in)

	n := parse(in)
	n.print()
}

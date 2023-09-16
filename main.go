package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

var version = "v0.0.0-unknown" // set with by compiler `-ldflags`

func main() {
	showVersion := flag.Bool("version", false, "show version")
	right := flag.Bool("r", false, "right associative")
	left := flag.Bool("l", false, "left associative")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	input := "1 * 2 + 3 * 4"
	if len(flag.Args()) > 0 {
		input = flag.Arg(0)
	}

	tokens := newTokens(input)

	var ast any
	if *right {
		ast = parseR(tokens)
	} else if *left {
		ast = parseL(tokens)
	} else {
		ast = parse(tokens, lowest)
	}
	fmt.Println(ast)
}

// tokens is a token iterator
type tokens struct {
	toks   []string
	curIdx int
}

func newTokens(input string) *tokens {
	return &tokens{toks: strings.Split(input, " ")}
}

func (t *tokens) done() bool {
	return t.curIdx >= len(t.toks)
}

func (t *tokens) cur() string {
	if t.done() {
		panic("no more tokens")
	}
	return t.toks[t.curIdx]
}

func (t *tokens) next() string {
	token := t.cur()
	t.curIdx++
	return token
}

func (t *tokens) curPrec() int {
	prec, ok := precedence[t.cur()]
	if !ok {
		panic("no precedence for " + t.cur())
	}
	return prec
}

const (
	lowest = iota
	eq
	sum
	product
)

var precedence = map[string]int{
	"==": eq,
	"+":  sum,
	"-":  sum,
	"*":  product,
	"/":  product,
}

// expr is a node in the AST
type expr struct {
	left  any
	op    string
	right any
}

func parseR(t *tokens) any {
	left := t.next()
	if t.done() {
		return left
	}
	op := t.next()
	right := parseR(t)
	return expr{left, op, right}
}

func parseL(t *tokens) any {
	var left any = t.next()
	for !t.done() {
		op := t.next()
		right := t.next()
		left = expr{left, op, right}
	}
	return left
}

func parse(t *tokens, prec int) any {
	var left any = t.next()
	for !t.done() && prec < t.curPrec() {
		p := t.curPrec()
		op := t.next()
		right := parse(t, p)
		left = expr{left, op, right}
	}
	return left
}

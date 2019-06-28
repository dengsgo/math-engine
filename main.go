package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	loop()
}

// input loop
func loop() {
	for {
		fmt.Print("input /> ")
		f := bufio.NewReader(os.Stdin)
		s, err := f.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		start := time.Now()
		exec(s)
		cost := time.Since(start)
		fmt.Println("time: " + cost.String())
	}
}

// engine
func exec(exp string) {
	// input text -> []token
	toks, err := parse(exp)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	// []token -> AST Tree
	ast := newAST(toks, exp)
	if ast.err != nil {
		fmt.Println("ERROR: " + ast.err.Error())
		return
	}
	// AST builder
	ar := ast.parseExpression()
	if ast.err != nil {
		fmt.Println("ERROR: " + ast.err.Error())
		return
	}
	fmt.Printf("ExprAST: %+v\n", ar)
	// AST traversal -> result
	r := binaryExec(ar)
	fmt.Println("binaryExec:", r)
	fmt.Printf("%s = %v\n", exp, r)
}

// AST traversal
func binaryExec(expr ExprAST) float64 {
	var l, r float64
	switch expr.(type) {
	case BinaryExprAST:
		ast := expr.(BinaryExprAST)
		switch ast.Lhs.(type) {
		case BinaryExprAST:
			l = binaryExec(ast.Lhs.(BinaryExprAST))
		default:
			l = ast.Lhs.(NumberExprAST).Val
		}
		switch ast.Rhs.(type) {
		case BinaryExprAST:
			r = binaryExec(ast.Rhs.(BinaryExprAST))
		default:
			r = ast.Rhs.(NumberExprAST).Val
		}
		switch ast.Op {
		case "+":
			return l + r
		case "-":
			return l - r
		case "*":
			return l * r
		case "/":
			return l / r
		case "%":
			return float64(int(l) % int(r))
		default:

		}
	case NumberExprAST:
		return expr.(NumberExprAST).Val
	}

	return 0.0
}

package main

import (
	"bufio"
	"fmt"
	"math-engine/engine"
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
		if s == "exit" || s == "quit" || s == "q" {
			fmt.Println("bye")
			break
		}
		start := time.Now()
		exec(s)
		cost := time.Since(start)
		fmt.Println("time: " + cost.String())
	}
}

// call engine
func exec(exp string) {
	// input text -> []token
	toks, err := engine.Parse(exp)
	if err != nil {
		fmt.Println("ERROR: " + err.Error())
		return
	}
	// []token -> AST Tree
	ast := engine.NewAST(toks, exp)
	if ast.Err != nil {
		fmt.Println("ERROR: " + ast.Err.Error())
		return
	}
	// AST builder
	ar := ast.ParseExpression()
	if ast.Err != nil {
		fmt.Println("ERROR: " + ast.Err.Error())
		return
	}
	fmt.Printf("ExprAST: %+v\n", ar)
	// AST traversal -> result
	r := engine.ExprASTResult(ar)
	fmt.Println("binaryExec:", r)
	fmt.Printf("%s = %v\n", exp, r)
}

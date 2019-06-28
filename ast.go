package main

import (
	"errors"
	"fmt"
	"strconv"
)

var precedence = map[string]int{"+": 20, "-": 20, "*": 40, "/": 40, "%": 40}

type ExprAST interface {
	toStr() string
}

type NumberExprAST struct {
	Val float64
}

type BinaryExprAST struct {
	Op string
	Lhs,
	Rhs ExprAST
}

func (n NumberExprAST) toStr() string {
	return fmt.Sprintf(
		"NumberExprAST:%s",
		strconv.FormatFloat(n.Val, 'f', 0, 64),
	)
}

func (b BinaryExprAST) toStr() string {
	return fmt.Sprintf(
		"BinaryExprAST: (%s %s %s)",
		b.Op,
		b.Lhs.toStr(),
		b.Rhs.toStr(),
	)
}

type AST struct {
	Tokens    []*Token
	source    string
	currTok   *Token
	currIndex int

	err error
}

func newAST(toks []*Token, s string) *AST {
	a := &AST{
		Tokens: toks,
		source: s,
	}
	if a.Tokens == nil || len(a.Tokens) == 0 {
		a.err = errors.New("empty token")
	} else {
		a.currIndex = 0
		a.currTok = a.Tokens[0]
	}
	return a
}

func (a *AST) parseExpression() ExprAST {
	lhs := a.parsePrimary()
	return a.parseBinOpRHS(0, lhs)
}

func (a *AST) getNextToken() *Token {
	a.currIndex++
	if a.currIndex < len(a.Tokens) {
		a.currTok = a.Tokens[a.currIndex]
		return a.currTok
	}
	return nil
}

func (a *AST) getTokPrecedence() int {
	if p, ok := precedence[a.currTok.Tok]; ok {
		return p
	}
	return -1
}

func (a *AST) parseNumber() NumberExprAST {
	f64, err := strconv.ParseFloat(a.currTok.Tok, 64)
	if err != nil {
		a.err = errors.New(
			fmt.Sprintf("%v\nwant '(' or '0-9' but get '%s'\n%s",
				err.Error(),
				a.currTok.Tok,
				errPos(a.source, a.currTok.offset)))
		return NumberExprAST{}
	}
	n := NumberExprAST{
		Val: f64,
	}
	a.getNextToken()
	return n
}

func (a *AST) parsePrimary() ExprAST {
	switch a.currTok.Type {
	case Literal:
		return a.parseNumber()
	case Operator:
		if a.currTok.Tok == "(" {
			a.getNextToken()
			e := a.parseExpression()
			if e == nil {
				return nil
			}
			if a.currTok.Tok != ")" {
				a.err = errors.New(
					fmt.Sprintf("want ')' but get %s\n%s",
						a.currTok.Tok,
						errPos(a.source, a.currTok.offset)))
				return nil
			}
			a.getNextToken()
			return e
		} else {
			return a.parseNumber()
		}
	default:
		return nil
	}
}

func (a *AST) parseBinOpRHS(execPrec int, lhs ExprAST) ExprAST {
	for {
		tokPrec := a.getTokPrecedence()
		if tokPrec < execPrec {
			return lhs
		}
		binOp := a.currTok.Tok
		a.getNextToken()
		rhs := a.parsePrimary()
		if rhs == nil {
			return nil
		}
		nextPrec := a.getTokPrecedence()
		if tokPrec < nextPrec {
			rhs = a.parseBinOpRHS(tokPrec+1, rhs)
			if rhs == nil {
				return nil
			}
		}
		lhs = BinaryExprAST{
			Op:  binOp,
			Lhs: lhs,
			Rhs: rhs,
		}
	}
}

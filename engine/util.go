package engine

import (
	"math"
	"strings"
)

var definedConst = map[string]float64{
	"pi": math.Pi,
}

var definedFunc map[string]func(expr ExprAST) float64

func initFunc() {
	definedFunc = map[string]func(expr ExprAST) float64{
		"sin": definedSin,
		"cos": definedCos,
		"tan": definedTan,
		"cot": definedCot,
		"sec": definedSec,
		"csc": definedCsc,
	}
}

func ErrPos(s string, pos int) string {
	r := strings.Repeat("-", len(s)) + "\n"
	s += "\n"
	for i := 0; i < pos; i++ {
		s += " "
	}
	s += "^\n"
	return r + s + r
}

// the integer power of a number
func Pow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	r := calPow(x, n)
	if n < 0 {
		r = 1 / r
	}
	return r
}

func calPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	r := calPow(x, n>>1) // move right 1 byte
	r *= r
	if n&1 == 1 {
		r *= x
	}
	return r
}

func angle2Radian(i float64) float64 {
	return i / 180 * math.Pi
}

func radian2Angle(i float64) float64 {
	return i / math.Pi * 180
}

func definedSin(expr ExprAST) float64 {
	return math.Sin(ExprASTResult(expr))
}

func definedCos(expr ExprAST) float64 {
	return math.Cos(ExprASTResult(expr))
}

func definedTan(expr ExprAST) float64 {
	return math.Tan(ExprASTResult(expr))
}

func definedCot(expr ExprAST) float64 {
	return 1 / definedTan(expr)
}

func definedSec(expr ExprAST) float64 {
	return 1 / definedCos(expr)
}

func definedCsc(expr ExprAST) float64 {
	return 1 / definedSin(expr)
}

// AST traversal
func ExprASTResult(expr ExprAST) float64 {
	var l, r float64
	switch expr.(type) {
	case BinaryExprAST:
		ast := expr.(BinaryExprAST)
		l = ExprASTResult(ast.Lhs)
		r = ExprASTResult(ast.Rhs)
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
		case "^":
			return Pow(l, int(r))
		default:

		}
	case NumberExprAST:
		return expr.(NumberExprAST).Val
	case FunCallerExprAST:
		f := expr.(FunCallerExprAST)
		return definedFunc[f.Name](f.Arg)
	}

	return 0.0
}

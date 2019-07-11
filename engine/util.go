package engine

import (
	"math"
	"strings"
)

// Top level function
// Analytical expression and execution
func ParseAndExec(s string) (float64, error) {
	toks, err := Parse(s)
	if err != nil {
		return 0, err
	}
	ast := NewAST(toks, s)
	if ast.Err != nil {
		return 0, ast.Err
	}
	ar := ast.ParseExpression()
	if ast.Err != nil {
		return 0, ast.Err
	}
	return ExprASTResult(ar), nil
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

func expr2Radian(expr ExprAST) float64 {
	r := ExprASTResult(expr)
	if TrigonometricMode == AngleMode {
		r = r / 180 * math.Pi
	}
	return r
}

// Top level function
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
		return defFunc[f.Name](f.Arg)
	}

	return 0.0
}

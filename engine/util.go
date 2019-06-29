package engine

import "strings"

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

// AST traversal
func ExprASTResult(expr ExprAST) float64 {
	var l, r float64
	switch expr.(type) {
	case BinaryExprAST:
		ast := expr.(BinaryExprAST)
		switch ast.Lhs.(type) {
		case BinaryExprAST:
			l = ExprASTResult(ast.Lhs)
		default:
			l = ast.Lhs.(NumberExprAST).Val
		}
		switch ast.Rhs.(type) {
		case BinaryExprAST:
			r = ExprASTResult(ast.Rhs)
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
		case "^":
			return Pow(l, int(r))
		default:

		}
	case NumberExprAST:
		return expr.(NumberExprAST).Val
	}

	return 0.0
}

package engine

import "math"

const (
	RadianMode = iota
	AngleMode
)

// enum "RadianMode", "AngleMode"
var TrigonometricMode = RadianMode

var defConst = map[string]float64{
	"pi": math.Pi,
}

var defFunc map[string]func(expr ExprAST) float64

func init() {
	defFunc = map[string]func(expr ExprAST) float64{
		"sin": defSin,
		"cos": defCos,
		"tan": defTan,
		"cot": defCot,
		"sec": defSec,
		"csc": defCsc,

		"abs":   defAbs,
		"ceil":  defCeil,
		"floor": defFloor,
		"sqrt":  defSqrt,
		"cbrt":  defCbrt,
	}
}

// sin(pi/2) = 1
func defSin(expr ExprAST) float64 {
	return math.Sin(expr2Radian(expr))
}

// cos(0) = 1
func defCos(expr ExprAST) float64 {
	return math.Cos(expr2Radian(expr))
}

// tan(pi/4) = 1
func defTan(expr ExprAST) float64 {
	return math.Tan(expr2Radian(expr))
}

// cot(pi/4) = 1
func defCot(expr ExprAST) float64 {
	return 1 / defTan(expr)
}

// sec(0) = 1
func defSec(expr ExprAST) float64 {
	return 1 / defCos(expr)
}

// csc(pi/2) = 1
func defCsc(expr ExprAST) float64 {
	return 1 / defSin(expr)
}

// abs(-2) = 2
func defAbs(expr ExprAST) float64 {
	return math.Abs(ExprASTResult(expr))
}

// ceil(4.2) = ceil(4.8) = 5
func defCeil(expr ExprAST) float64 {
	return math.Ceil(ExprASTResult(expr))
}

// floor(4.2) = floor(4.8) = 4
func defFloor(expr ExprAST) float64 {
	return math.Floor(ExprASTResult(expr))
}

// sqrt(4) = 2
func defSqrt(expr ExprAST) float64 {
	return math.Sqrt(ExprASTResult(expr))
}

// cbrt(27) = 3
func defCbrt(expr ExprAST) float64 {
	return math.Cbrt(ExprASTResult(expr))
}

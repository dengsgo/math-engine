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

var defFunC = map[string]int{
	"sin": 1,
	"cos": 1,
	"tan": 1,
	"cot": 1,
	"sec": 1,
	"csc": 1,

	"abs":   1,
	"ceil":  1,
	"floor": 1,
	"round": 1,
	"sqrt":  1,
	"cbrt":  1,

	"max": 2,
	"min": 2,
}

var defFunc map[string]func(expr []ExprAST) float64

func init() {
	defFunc = map[string]func(expr []ExprAST) float64{
		"sin": defSin,
		"cos": defCos,
		"tan": defTan,
		"cot": defCot,
		"sec": defSec,
		"csc": defCsc,

		"abs":   defAbs,
		"ceil":  defCeil,
		"floor": defFloor,
		"round": defRound,
		"sqrt":  defSqrt,
		"cbrt":  defCbrt,

		"max": defMax,
		"min": defMin,
	}
}

// sin(pi/2) = 1
func defSin(expr []ExprAST) float64 {
	return math.Sin(expr2Radian(expr[0]))
}

// cos(0) = 1
func defCos(expr []ExprAST) float64 {
	return math.Cos(expr2Radian(expr[0]))
}

// tan(pi/4) = 1
func defTan(expr []ExprAST) float64 {
	return math.Tan(expr2Radian(expr[0]))
}

// cot(pi/4) = 1
func defCot(expr []ExprAST) float64 {
	return 1 / defTan(expr)
}

// sec(0) = 1
func defSec(expr []ExprAST) float64 {
	return 1 / defCos(expr)
}

// csc(pi/2) = 1
func defCsc(expr []ExprAST) float64 {
	return 1 / defSin(expr)
}

// abs(-2) = 2
func defAbs(expr []ExprAST) float64 {
	return math.Abs(ExprASTResult(expr[0]))
}

// ceil(4.2) = ceil(4.8) = 5
func defCeil(expr []ExprAST) float64 {
	return math.Ceil(ExprASTResult(expr[0]))
}

// floor(4.2) = floor(4.8) = 4
func defFloor(expr []ExprAST) float64 {
	return math.Floor(ExprASTResult(expr[0]))
}

// round(4.2) = 4
// round(4.6) = 5
func defRound(expr []ExprAST) float64 {
	return math.Round(ExprASTResult(expr[0]))
}

// sqrt(4) = 2
func defSqrt(expr []ExprAST) float64 {
	return math.Sqrt(ExprASTResult(expr[0]))
}

// cbrt(27) = 3
func defCbrt(expr []ExprAST) float64 {
	return math.Cbrt(ExprASTResult(expr[0]))
}

// max(2, 3) = 3
func defMax(expr []ExprAST) float64 {
	return math.Max(p2(expr))
}

// max(2, 3) = 2
func defMin(expr []ExprAST) float64 {
	return math.Min(p2(expr))
}

func p2(expr []ExprAST) (float64, float64) {
	return ExprASTResult(expr[0]), ExprASTResult(expr[1])
}

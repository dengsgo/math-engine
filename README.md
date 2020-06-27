## Math-Engine  

[![Build Status](https://travis-ci.org/dengsgo/math-engine.svg?branch=master)](https://travis-ci.org/dengsgo/math-engine)  [![Go Report Card](https://goreportcard.com/badge/github.com/dengsgo/math-engine)](https://goreportcard.com/report/github.com/dengsgo/math-engine)  [![godoc.org](https://godoc.org/github.com/dengsgo/math-engine/engine?status.svg)](https://godoc.org/github.com/dengsgo/math-engine/engine)  

使用 Go 实现的数学表达式解析计算引擎库，它小巧，无任何依赖，具有扩展性(比如可以注册自己的函数到引擎中)，比较完整的完成了数学表达式解析执行，包括词法分析、语法分析、构建AST、运行。  

`go get -u github.com/dengsgo/math-engine`  

能够处理的表达式样例：  
- `1+127-21+(3-4)*6/2.5`  
- `(88+(1+8)*6)/2+99`  
- `123_345_456 * 1.5 - 2 ^ 4`  
- `-4 * 6 + 2e2 - 1.6e-3`  
- `sin(pi/2)+cos(45-45*1)+tan(pi/4)`  
- `99+abs(-1)-ceil(88.8)+floor(88.8)`  
- `max(min(2^3, 3^2), 10*1.5-7)`  
- `double(6) + 3` , `double`是一个自定义的函数  

### Demo

[![asciicast](https://asciinema.org/a/276195.svg)](https://asciinema.org/a/276195)

## Method Support

| symbol      | explanation                  | e.g.                                  |
| ----------- | ---------------------------- | ------------------------------------- |
| `+`         | 加，plus                     | 1+2 = 3                               |
| `-`         | 减，sub                      | 8-3.5 = 4.5                           |
| `*`         | 乘，multiply                 | 2*3 = 6                               |
| `/`         | 除，division                 | 5/2 = 2.5                             |
| `%`         | 取余，remainder              | 5%2 = 1                               |
| `^`         | 整数次方，integer power      | 2^3 = 8, 3^2 = 9                      |
| `e`         | 科学计数法，E-notation       | 1.2e3 = 1.2e+3 = 1200，1.2e-2 = 0.012 |
| `()`        | 括号，brackets               | (2+3)*4 = 20                          |
| `_`         | 数字分隔符，number separator | 123_456_789 = 123456789               |
| `pi`        | π                            | pi = 3.141592653589793                |
| `sin(x)`    | 正弦函数，sine               | sin(pi/2) = 1                         |
| `cos(x)`    | 余弦函数，cosine             | cos(0) = 1                            |
| `tan(x)`    | 正切函数，tangent            | tan(pi/4) = 1                         |
| `cot(x)`    | 余切函数，cotangent          | cot(pi/4) = 1                         |
| `sec(x)`    | 正割函数，secant             | sec(0) = 1                            |
| `csc(x)`    | 余割函数，cosecant           | csc(pi/2) = 1                         |
| `abs(x)`    | 绝对值，absolute value       | abs(-6) = 6                           |
| `ceil(x)`   | 向上取整                     | ceil(4.2) = 5                         |
| `floor(x)`  | 向下取整                     | floor(4.8) = 4                        |
| `round(x)`  | 四舍五入取整                 | round(4.4) = 4, round(4.5) = 5        |
| `sqrt(x)`   | 平方根，square root          | sqrt(4) = 2                           |
| `cbrt(x)`   | 立方根，cube root            | cbrt(27) = 3                          |
| `max(x, y)` | x, y 中的较大值              | max(2, 3) = 3                         |
| `min(x, y)` | x, y 中的较小值              | min(2, 3) = 2                         |
| `noerr(x)`  | 计算 x 出错时返回 0          | noerr(1 / 1)  = 1, noerr( 1/ 0 ) = 0  |
| `double(x)`  | 返回 x 的双倍值，这是一个自定义的函数示例，你可以注册任意的自定义函数到引擎中  | double(6) = 12  |


## Usage  

你可以直接引用该库嵌入到自己的程序中：  
```bash
go get -u github.com/dengsgo/math-engine
```
在代码中引入：  
```go
import "github.com/dengsgo/math-engine/engine"
```
e.g. 1 常规用法： 直接调用解析执行函数 :

```go
import "github.com/dengsgo/math-engine/engine"

func main() {
  s := "1 + 2 * 6 / 4 + (456 - 8 * 9.2) - (2 + 4 ^ 5)"
  // call top level function
  r, err := engine.ParseAndExec(s)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("%s = %v", s, r)
}
```



e.g. 2 高级用法： 依次调用函数，手动执行 :  

```go
import "github.com/dengsgo/math-engine/engine"

func main() {
	s := "1 + 2 * 6 / 4 + (456 - 8 * 9.2) - (2 + 4 ^ 5)"
	exec(s)
}

// call engine
// one by one
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
	fmt.Println("progressing ...\t", r)
	fmt.Printf("%s = %v\n", exp, r)
}
```
编译运行，应该可以看到如下输出：  
```bash
ExprAST: {Op:- Lhs:{Op:+ Lhs:{Op:+ Lhs:{Val:1} Rhs:{Op:/ Lhs:{Op:* Lhs:{Val:2} Rhs:{Val:6}} Rhs:{Val:4}}} Rhs:{Op:- Lhs:{Val:456} Rhs:{Op:* Lhs:{Val:8} Rhs:{Val:9.2}}}} Rhs:{Op:+ Lhs:{Val:2} Rhs:{Op:^ Lhs:{Val:4} Rhs:{Val:5}}}}
progressing ...  -639.6
1+2*6/4+(456-8*9.2)-(2+4^5) = -639.6
```

## TrigonometricMode

三角函数的参数类型默认为弧度`RadianMode`，e.g. `sin(pi/2) = 1`.

你可以通过设置 `TrigonometricMode` 调整参数类型，可选 弧度`RadianMode`、角度`AngleMode`，e.g. :

```go
import "github.com/dengsgo/math-engine/engine"

func main() {
  s := "1 + sin(90)"
  engine.TrigonometricMode = engine.AngleMode
  engine.ParseAndExec(s) // will return 2, nil
  s = "1 + sin(pi/2)"
  engine.TrigonometricMode = engine.RadianMode
  engine.ParseAndExec(s) // will return 2, nil
}
```

## Register Function

`math-engine` 提供了自定义函数注册到引擎的能力。你可以把常用的函数注册到引擎中，然后就能像内置函数一样在输入的数学表达式中使用。

e.g 

```go
  // RegFunction is Top level function
  // the same function name only needs to be registered once.
  // double is register function name.
  // 1 is a number of parameter signatures.
  // func(expr ...engine.ExprAST) float64 is your function.
  engine.RegFunction("double", 1, func(expr ...engine.ExprAST) float64 {
    // you can use the index value directly according to the number of parameters
    // without worrying about crossing the boundary.
    // use ExprASTResult to get the result of the ExprAST structure.
    return engine.ExprASTResult(expr[0]) * 2
  })
```

然后你就可以在输入的表达式中使用这个函数 `double`:

```go 
//exp := "double(6) + 2"
r, err := engine.ParseAndExec("double(6) + 2")
if err != nil {
  panic(err)
}
fmt.Printf("double(6) + 2 = %f\n", r) // will print ： double(6) + 2 = 14.000000
```

注意事项：
- 注册的函数名只能是英文字母和数字，且必须英文字母开头（区分大小写）;
- 每一个函数名只能且只需注册一次；
- 注册的函数逻辑中如果有 panic ，需要程序自己捕获处理;

## Compile    

go version 1.12  
```go
# Compile Demo
go test
go build
./math-engine
```

也可以直接下载已编译好的二进制文件，直接运行：  

[Github Releases](https://github.com/dengsgo/math-engine/releases)  

## 实现细节    

请阅读我的博客文章：[用 Go 实现一个完整的数学表达式计算引擎](https://www.yoytang.com/math-expression-engine.html)  

## TODO  
### 已实现  

- [x] 加 `+`  
- [x] 减 `-`  
- [x] 乘 `*`  
- [x] 除 `/`  
- [x] 取余 `%`  
- [x] 整数次方 `^`  
- [x] 科学计数法 e.g. `1.2e7`、  `1.2e-7`
- [x] 括号 `()`  
- [x] 混合运算 e.g. `1+2*6/4+(456-8*9.2)-(2+4^5)*2e3+1.2e-2`  
- [x] 友好的长数字 e.g. `123_456_789`  
- [x] 三角函数 e.g. `sin, cos, tan, cot, sec, csc`
- [x] 常量 pi
- [x] 辅助函数 e.g. `abs, ceil, floor, sqrt, cbrt, max, min, noerr`
- [x] 提供自定义函数注册功能，注册后可以在表达式中使用
- [x] 精确的数据计算
- [x] 友好的错误消息 e.g.    
```bash
input /> 123+89-0.0.9
ERROR: strconv.ParseFloat: parsing "0.0.9": invalid syntax
want '(' or '0-9' but get '0.0.9'
------------
123+89-0.0.9
       ^
------------
```


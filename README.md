## Math-Engine  

使用 Go 实现的数学表达式微型计算引擎，无任何依赖，相对比较完整的完成了数学表达式解析执行，包括词法分析、语法分析、构建AST、运行。  

## TODO  
### 已实现  
  
- [x] 加 +  
- [x] 减 -  
- [x] 乘 *  
- [x] 除 /  
- [x] 取余 %  
- [x] 整数次方 ^  
- [x] 科学计数法 e.g. 1.2e7  
- [x] 括号 ()  
- [x] 混合运算 e.g. 1+2*6/4+(456-8*9.2)-(2+4^5)  
- [x] 友好的长数字 e.g. 123_456_789  
- [x] 友好的错误消息  

### 待实现  
  
- [ ] 科学计数法 e.g. 1.2-e7  
- [ ] 精确的浮点计算    

## Usage  

你可以直接引用该库嵌入到自己的程序中：  
```bash
go get -u github.com/dengsgo/math-engine
```  
在代码中引入：  
```go
import "github.com/dengsgo/math-engine/engine"
```  
e.g. :  
```go
import "github.com/dengsgo/math-engine/engine"

func main() {
	s := "1 + 2 * 6 / 4 + (456 - 8 * 9.2) - (2 + 4 ^ 5)"
	exec(s)
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


## Compile    

go version 1.12  
```go
go test
go build
./math-engine
```
  
也可以直接下载已编译好的二进制文件，直接运行：  

[Github Releases](https://github.com/dengsgo/math-engine/releases)  

## 实现原理  

TODO
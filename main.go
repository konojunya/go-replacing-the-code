package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	expr, err := parser.ParseExprFrom(fset, "input/in.go", `World`, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println("==== BEFORE ====")
	ast.Print(fset, expr)
	binaryExpr := expr.(*ast.Ident)
	binaryExpr.Name = "世界"

	fmt.Println("==== AFTER ====")
	ast.Print(fset, expr)
}

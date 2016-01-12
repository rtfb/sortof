package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type visitor struct {
	allCalls   []string
	allAssigns []string
	fileSet    *token.FileSet
}

func (v *visitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.CallExpr:
		call, _ := n.Fun.(*ast.Ident)
		if call != nil {
			fmt.Printf("call: %q\n", call.Name)
		}
	case *ast.AssignStmt:
		if len(n.Lhs) > 0 {
			if id, ok := n.Lhs[0].(*ast.Ident); ok {
				fmt.Printf("assign: %q\n", id.Name)
			}
		}
	}
	return v
}

func main() {
	files := []string{
		"parz.go",
	}
	v := &visitor{}
	for _, fileName := range files {
		v.fileSet = token.NewFileSet()
		f, err := parser.ParseFile(v.fileSet, fileName, nil, 0)
		if err != nil {
			panic(err) // XXX: better error handling
		}
		ast.Walk(v, f)
	}
}

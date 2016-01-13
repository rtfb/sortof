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
		var call string
		switch fun := n.Fun.(type) {
		case *ast.Ident:
			call = fmt.Sprintf("%s", n.Fun)
		case *ast.SelectorExpr:
			call = fmt.Sprintf("%s.%s", fun.X, fun.Sel)
		}
		v.allCalls = append(v.allCalls, call)
	case *ast.AssignStmt:
		if len(n.Lhs) > 0 {
			if id, ok := n.Lhs[0].(*ast.Ident); ok {
				v.allAssigns = append(v.allAssigns, id.Name)
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
	fmt.Println("Calls:")
	for _, c := range v.allCalls {
		fmt.Printf("\t%s\n", c)
	}
	fmt.Println("Assigns:")
	for _, a := range v.allAssigns {
		fmt.Printf("\t%s\n", a)
	}
}

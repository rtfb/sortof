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

type placeholderVisitor struct {
	meat    ast.Node
	fileSet *token.FileSet
}

func (v *placeholderVisitor) Visit(node ast.Node) ast.Visitor {
	switch n := node.(type) {
	case *ast.FuncDecl:
		fmt.Printf("FuncDecl: %q\n", n.Name)
		vv := &visitor{}
		vv.fileSet = token.NewFileSet()
		ast.Walk(vv, n)
		dump(vv)
		return nil
	}
	return v
}

func dump(v *visitor) {
	fmt.Println("Calls:")
	for _, c := range v.allCalls {
		fmt.Printf("\t%s\n", c)
	}
	fmt.Println("Assigns:")
	for _, a := range v.allAssigns {
		fmt.Printf("\t%s\n", a)
	}
}

func runWithBoilerplate() {
	fmt.Println("===== boilerplate ==========")
	src := `package placeholder
func noise() {
	a := 0
	b := noooize()
	c := moarnoize()
}
func placeholder() {
	signal := true
	return signalProcessor(temp)
}`
	v := &placeholderVisitor{}
	v.fileSet = token.NewFileSet()
	tree, err := parser.ParseFile(v.fileSet, "", src, 0)
	if err != nil {
		panic(err)
	}
	ast.Walk(v, tree)
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
	dump(v)
	fmt.Println("===============")
	src := `func foo() bool {
	temp := true
	return foo_r(temp)
}`
	vv := &visitor{}
	vv.fileSet = token.NewFileSet()
	prefix := "package placeholder\n"
	tree, err := parser.ParseFile(vv.fileSet, "", prefix+src, 0)
	if err != nil {
		panic(err)
	}
	ast.Walk(vv, tree)
	dump(vv)
	fmt.Println("===============")
	src = "blerk(\"param\")"
	vvv := &visitor{}
	vvv.fileSet = token.NewFileSet()
	expr, err := parser.ParseExpr(src)
	if err != nil {
		panic(err)
	}
	ast.Walk(vvv, expr)
	dump(vvv)
	runWithBoilerplate()
}

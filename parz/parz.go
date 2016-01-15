package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	//"reflect"
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

type channelPusher struct {
	fileSet *token.FileSet
	queue   chan *ast.Node
}

func (v *channelPusher) Visit(node ast.Node) ast.Visitor {
	v.queue <- &node
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

func equalNodes(n1, n2 *ast.Node) bool {
	//if reflect.TypeOf(n1).Name() != reflect.TypeOf(n2).Name() {
	//	return false
	//}
	if n1.Name != n2.Name {
		return false
	}
	// TODO: compare other criteria
	return true
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

func compareTwoTrees(src string) {
	v1 := &channelPusher{}
	v1.fileSet = token.NewFileSet()
	v1.queue = make(chan *ast.Node)

	v2 := &channelPusher{}
	v2.fileSet = token.NewFileSet()
	v2.queue = make(chan *ast.Node)

	tree1, err := parser.ParseExpr(src)
	if err != nil {
		panic(err)
	}

	src2 := "x + 2*y"
	tree2, err := parser.ParseExpr(src2)
	if err != nil {
		panic(err)
	}
	done := make(chan struct{})
	defer close(done)

	go func() {
		ast.Walk(v1, tree1)
		close(v1.queue)
		done <- struct{}{}
	}()
	go func() {
		ast.Walk(v2, tree2)
		close(v2.queue)
		done <- struct{}{}
	}()

	var n1, n2 *ast.Node
	quit := false
	for !quit {
		select {
		case n1 = <-v1.queue:
		case n2 = <-v2.queue:
		case <-done:
			quit = true
		}
		if n1 != nil && n2 != nil {
			if !equalNodes(n1, n2) {
				println("!equalNodes")
				break
			}
			println("equalNodes")
			n1 = nil
			n2 = nil
		}
	}
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
	compareTwoTrees(src)
}

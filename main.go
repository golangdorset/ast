package main

import (
	f "fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	node, err := parser.ParseFile(token.NewFileSet(), "main.go", nil, parser.ImportsOnly)
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	ast.Inspect(node, func(n ast.Node) bool {
		// Prints out import paths and grabs the alias if available.
		if imp, ok := n.(*ast.ImportSpec); ok {
			if imp.Name != nil {
				f.Println(imp.Path.Value, imp.Name.Name)
			} else {
				f.Println(imp.Path.Value)
			}
		}
		return true
	})
}

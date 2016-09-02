package main

import (
	//"bytes"
	"fmt"
	//"go/ast"
	//"go/format"
	//"go/parser"
	//"go/token"
	"reflect"
	"github.com/charleswisniewski/hecate/models"
)



func main() {
	//fset := token.NewFileSet() // positions are relative to fset

	// Parse the file containing this very example
	// but stop after processing the imports.
	//f, err := parser.ParseFile(fset, "/Users/bob/golang/src/github.com/charleswisniewski/hecate/models/deploy.go", nil, 0)//parser.ParseComments)
/*	if err != nil {
		fmt.Println(err)
		return
	}*/
	d := models.Deploy{}
//	fmt.Printf("%+v\n", d)
	dr := reflect.ValueOf(&d).Elem()
	typeOfT := dr.Type()
	for i := 0; i < dr.NumField(); i++ {
		f := dr.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(), f.Interface())
	}



	//cmap := ast.NewCommentMap(fset, f, f.Comments)
	//f.Decls = removeFirstVarDecl(f.Decls)


	// Use the comment map to filter comments that don't belong anymore
	// (the comments associated with the variable declaration), and create
	// the new comments list.
	//f.Comments = cmap.Filter(f).Comments()

	// Print the modified AST.
/*	var buf bytes.Buffer
	if err := format.Node(&buf, fset, f); err != nil {
		panic(err)
	}*/
	//fmt.Printf("%s", buf.Bytes())
	// Print the imports from the file's AST.
	//for _, s := range f.Nodes {
		//fmt.Println(s.)
	//	fmt.Println(FindTypes(s))
		 // Create an ast.CommentMap from the ast.File's comments.
    // This helps keeping the association between comments
    // and AST nodes.
		//cmap := ast.NewCommentMap(fset, f, f.Comments)
	//}
}
/*type VisitorFunc func(n ast.Node) ast.Visitor

func (f VisitorFunc) Visit(n ast.Node) ast.Visitor { return f(n) }

func FindTypes(n ast.Node) ast.Visitor {
    switch n := n.(type) {
    case *ast.Package:
        return VisitorFunc(FindTypes)
    case *ast.File:
        return VisitorFunc(FindTypes)
    case *ast.GenDecl:
        if n.Tok == token.TYPE {
            return VisitorFunc(FindTypes)
        }
    case *ast.TypeSpec:
        fmt.Println(n.Name.Name)
    }
    return nil
}
*/
/*func removeFirstVarDecl(list []ast.Decl) []ast.Decl {
	for i, decl := range list {
		//fmt.Printf()
		if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.TYPE {
			//fmt.Printf("%+v", gen)
			//fmt.Println(gen.Doc)
			//fmt.Println(gen.Spec)
			//fmt.Println(*ast.Spec)
			
			copy(list[i:], list[i+1:])
			return list[:len(list)-1]
		}
	}
	panic("variable declaration not found")
}*/

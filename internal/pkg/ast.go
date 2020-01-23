package pkg

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

func (c *Config) ParseDir() error {
	return afero.Walk(c.FS, c.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fset := token.NewFileSet()
			packages, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
			if err != nil {
				log.Fatalf("error parsing directory: %v", err)
			}

			for _, pkg := range packages {
				for p, f := range pkg.Files {
					if err := c.ProcessFile(fset, p, f); err != nil {
						log.Fatalf("error processing file: %v", err)
					}
				}
			}
		}
		return nil
	})
}

func (c *Config) ProcessFile(fset *token.FileSet, path string, node *ast.File) error {
	comments := []*ast.CommentGroup{}
	ast.Inspect(node, func(n ast.Node) bool {
		cg, ok := n.(*ast.CommentGroup)
		if ok {
			comments = append(comments, cg)
		}

		fn, ok := n.(*ast.FuncDecl)
		if ok && fn.Doc.Text() != "" {
			text, err := c.Justify(fn.Doc.Text())
			if err != nil {
				return false
			}
			comment := &ast.Comment{
				Text:  text,
				Slash: fn.Pos() - 1,
			}

			ncg := &ast.CommentGroup{
				List: []*ast.Comment{comment},
			}
			fn.Doc = ncg
		}

		return true
	})

	node.Comments = comments

	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "writing to file")
	}
	defer f.Close()

	if err := printer.Fprint(f, fset, node); err != nil {
		return errors.Wrapf(err, "writing updating code to %q", f.Name())
	}

	return nil
}

package pkg

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

func (c *Config) ParseRoot() error {
	return afero.Walk(c.FS, c.RootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fset := token.NewFileSet()
			packages, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
			if err != nil {
				return errors.Wrap(err, "parsing directory")
			}

			for _, pkg := range packages {
				for p, f := range pkg.Files {
					if err := c.ProcessFile(fset, p, f); err != nil {
						return errors.Wrap(err, "processing file")
					}
				}
			}
		}
		return nil
	})
}

func (c *Config) ProcessFile(fset *token.FileSet, path string, node *ast.File) error {
	cmap := ast.NewCommentMap(fset, node, node.Comments)

	for _, cg := range cmap.Comments() {
		pos := cg.List[len(cg.List)-1].Slash

		justified, err := c.Justify(cg.Text())
		if err != nil {
			return errors.Wrap(err, "justifying")
		}

		cg.List = []*ast.Comment{
			&ast.Comment{
				Text:  justified,
				Slash: pos,
			},
		}
	}

	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "creating file")
	}
	defer f.Close()

	if err := printer.Fprint(f, fset, node); err != nil {
		return errors.Wrapf(err, "writing updating code to %q", f.Name())
	}

	return nil
}

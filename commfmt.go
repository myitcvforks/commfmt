package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func main() {
	log.SetFlags(0)

	dir := flag.String("path", ".", "the relative/absolute path of the root directory.")
	prefix := flag.String("prefix", "// ", "the comment prefix.")
	width := flag.Int("width", 80, "the maximum width of comments.")
	flag.Parse()

	ctx := &context{
		rootDir: *dir,
		prefix:  *prefix,
		width:   *width,
	}

	if err := ctx.parseDir(); err != nil {
		log.Fatalf("error parsing directory: %v", err)
	}
}

type context struct {
	rootDir string
	prefix  string
	width   int
}

func (ctx *context) parseDir() error {
	return filepath.Walk(ctx.rootDir, func(path string, info os.FileInfo, err error) error {
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
					if err := ctx.processFile(fset, p, f); err != nil {
						log.Fatalf("error processing file: %v", err)
					}
				}
			}
		}
		return nil
	})
}

func (ctx *context) processFile(fset *token.FileSet, path string, node *ast.File) error {
	comments := []*ast.CommentGroup{}
	ast.Inspect(node, func(n ast.Node) bool {
		c, ok := n.(*ast.CommentGroup)
		if ok {
			comments = append(comments, c)
		}

		fn, ok := n.(*ast.FuncDecl)
		if ok {
			if fn.Doc.Text() != "" {
				comment := &ast.Comment{
					Text:  ctx.justify(fn.Doc.Text()),
					Slash: fn.Pos() - 1,
				}

				cg := &ast.CommentGroup{
					List: []*ast.Comment{comment},
				}
				fn.Doc = cg
			}
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

// justify  takes an input and a desired line length and returns a fully-justified version of
// the  text.  Taken wholesale from Malik Browne's great JavaScript walkthrough that does the
// same https://www.malikbrowne.com/blog/text-justification-coding-question
func (ctx *context) justify(input string) string {
	words := strings.Fields(input)

	lines := []string{}
	index := 0

	for index < len(words) {
		count := len(words[index])
		last := index + 1

		for last < len(words) {
			if len(words[last])+count+1 > ctx.width {
				break
			}
			count += len(words[last]) + 1
			last++
		}

		line := ctx.prefix
		difference := last - index - 1
		if last == len(words) || difference == 0 {
			for i := index; i < last; i++ {
				line += words[i] + " "
			}

			line = line[0 : len(line)-1]
			for i := len(line); i < ctx.width; i++ {
				line += " "
			}
		} else {
			spaces := (ctx.width - count) / difference
			remainder := (ctx.width - count) % difference

			for i := index; i < last; i++ {
				line += words[i]

				if i < last-1 {
					limit := spaces
					if (i - index) < remainder {
						limit++
					}

					for j := 0; j <= limit; j++ {
						line += " "
					}
				}
			}
		}

		lines = append(lines, line)
		index = last
	}

	return strings.Join(lines, "\n")
}

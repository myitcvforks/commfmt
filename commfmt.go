package main

import (
	"flag"
	"log"

	"github.com/codingconcepts/commfmt/internal/pkg"
	"github.com/spf13/afero"
)

func main() {
	log.SetFlags(0)

	dir := flag.String("path", ".", "the relative/absolute path of the root directory.")
	width := flag.Int("width", 80, "the maximum width of comments.")
	flag.Parse()

	c := &pkg.Config{
		FS:		afero.NewOsFs(),
		RootDir:	*dir,
		Width:		*width,
	}

	if err := c.ParseDir(); err != nil {
		log.Fatalf("error parsing directory: %v", err)
	}
}

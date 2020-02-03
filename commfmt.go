package main

import (
	"flag"
	"log"

	"github.com/codingconcepts/commfmt/internal/pkg"
	"github.com/spf13/afero"
)

func main() {
	log.SetFlags(0)

	path := flag.String("p", ".", "the relative/absolute path of the root directory.")
	width := flag.Int("w", 80, "the maximum width of comments.")
	flag.Parse()

	c := &pkg.Config{
		FS:       afero.NewOsFs(),
		RootPath: *path,
		Width:    *width,
	}

	if err := c.ParseRoot(); err != nil {
		log.Fatalf("error parsing directory: %v", err)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/codingconcepts/commfmt/internal/pkg"
	"github.com/spf13/afero"
)

func main() {
	os.Exit(main1())
}

func main1() int {
	switch err := mainerr(); err {
	case nil:
		return 0
	case flag.ErrHelp:
		return 2
	default:
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
}

func mainerr() error {
	log.SetFlags(0)
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	path := fs.String("p", ".", "the relative/absolute path of the root directory.")
	width := fs.Int("w", 80, "the maximum width of comments.")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return err
	}
	c := &pkg.Config{
		FS:       afero.NewOsFs(),
		RootPath: *path,
		Width:    *width,
	}
	if err := c.ParseRoot(); err != nil {
		return fmt.Errorf("error parsing directory: %v", err)
	}
	return nil
}

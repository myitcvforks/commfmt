package pkg

import "github.com/spf13/afero"

// Config holds the configuration required to justify text.
type Config struct {
	FS       afero.Fs
	RootPath string
	Width    int
}

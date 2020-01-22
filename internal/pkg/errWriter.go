package pkg

import "strings"

type errWriter struct {
	err	error
	builder	strings.Builder
}

func newErrWriter() *errWriter {
	return &errWriter{
		builder: strings.Builder{},
	}
}

func (ew *errWriter) write(s string) {
	if ew.err != nil {
		return
	}

	_, ew.err = ew.builder.WriteString(s)
}

func (ew *errWriter) string() string {
	return ew.builder.String()
}

func (ew *errWriter) error() error {
	return ew.err
}

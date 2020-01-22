package pkg

import (
	"testing"

	"github.com/codingconcepts/commfmt/internal/test"
)

func TestJustify(t *testing.T) {
	cases := []struct {
		name  string
		input string
		width int
		exp   string
	}{
		{
			name:  "single line less than width",
			input: "aa aa aa aa",
			width: 14,
			exp:   "// aa aa aa aa",
		},
		{
			name:  "single line more than width",
			input: "aa aa aa aa aa aa aa aa aa",
			width: 14,
			exp:   "// aa aa aa aa aa\n// aa aa aa aa",
		},
		{
			name:  "single line more than width end padding",
			input: "aa aa aa aa aa aa aa aa",
			width: 14,
			exp:   "// aa aa aa aa aa\n// aa aa aa",
		},
		{
			name:  "multi line more than width",
			input: "aa aa aa aa\naa aa aa aa aa",
			width: 14,
			exp:   "// aa aa aa aa aa\n// aa aa aa aa",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			cfg := &Config{
				Width: c.width,
			}

			act, err := cfg.Justify(c.input)
			test.Assert(t, err == nil, "unexpected error: %v", err)
			test.Equals(t, c.exp, act)
		})
	}
}

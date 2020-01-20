package main

import (
	"reflect"
	"testing"
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
			exp:   "// aa aa aa aa aa\n// aa aa aa   ",
		},
		{
			name:  "multi line more than width",
			input: "aa aa aa aa\naa aa aa aa aa",
			width: 14,
			exp:   "// aa aa aa aa aa\n// aa aa aa aa",
		},
		{
			name:  "single line more than width end padding",
			input: "aa aa aa aa aa aa aa aa",
			width: 14,
			exp:   "// aa aa aa aa aa\n// aa aa aa   ",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := &context{
				width:  c.width,
				prefix: "// ",
			}

			act := ctx.justify(c.input)
			equals(t, c.exp, act)
		})
	}
}

func equals(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("\nexp:\t'%[1]v' (%[1]T)\ngot:\t'%[2]v' (%[2]T)", exp, act)
	}
}

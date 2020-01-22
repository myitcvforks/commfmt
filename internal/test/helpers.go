package test

import (
	"reflect"
	"testing"
)

// Equals  performs  a deep equal against exp and act and fails the given
// test if they're found to be different.
func Equals(tb testing.TB, exp, act interface{}) {
	tb.Helper()
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("\nexp:\t'%[1]v' (%[1]T)\ngot:\t'%[2]v' (%[2]T)", exp, act)
	}
}

// Assert  fails  a  given  test  if  a  condition  is found to be false,
// optionally displaying a message to help diagnose the failure.
func Assert(tb testing.TB, cond bool, format string, args ...interface{}) {
	tb.Helper()
	if !cond {
		tb.Fatalf(format, args...)
	}
}

// ErrorNil fails a given test if an error is not nil and prints the error.
func ErrorNil(tb testing.TB, err error) {
	tb.Helper()
	if err != nil {
		tb.Fatal(err)
	}
}

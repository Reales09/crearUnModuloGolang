package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Juan"

	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Hello("Juan")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Juan) = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

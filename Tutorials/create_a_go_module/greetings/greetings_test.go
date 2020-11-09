package greetings

import (
	"testing"
	"regexp"
)

func TestHelloName (t *testing.T) {
	name := "Siddharth"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello ("Siddhart")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Siddharth") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
}

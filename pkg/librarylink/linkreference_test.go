package librarylink

import (
	"errors"
	"testing"
)

func TestLRValidate(t *testing.T) {
	var want error
	var got error
	l := LinkReference{}

	l.Name = "he!!0-Will"
	got = l.Validate()
	want = errors.New("LinkReference:name error 'Name 'he!!0-Will' does not conform " +
		"to the standard. Please check for extra whitespace and see " +
		"https://ethpm.github.io/ethpm-spec/glossary.html#term-identifier for the requirement.'")

	if got.Error() != want.Error() {
		t.Fatalf("Got '%v', expected '%v'", got, want)
	}

	l.Name = "hello-Me"
	got = l.Validate()
	if got != nil {
		t.Fatalf("Got '%v', expected <nil>", got)
	}
}

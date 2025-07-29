package di

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Oscar")

	got := buffer.String()
	want := "Hello, Oscar"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

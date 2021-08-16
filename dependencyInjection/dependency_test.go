package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Mbaku")

	got := buffer.String()
	want := "Hello, Mbaku"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

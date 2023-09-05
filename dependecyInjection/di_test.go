package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Angel") // Buffer implements the io.Writer interface

	got := buffer.String()
	want := "Hello, Angel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

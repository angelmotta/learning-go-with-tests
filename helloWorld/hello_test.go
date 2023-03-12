package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Angel")
	want := "Hello, Angel"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

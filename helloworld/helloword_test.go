package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Gladys")

	want := "Hello, Gladys"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

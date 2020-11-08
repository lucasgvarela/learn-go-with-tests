package main

import "testing"

func TestHelloYou(t *testing.T) {
	got := HelloYou("Lucas")
	want := "Hello, Lucas"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

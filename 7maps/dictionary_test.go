package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "testing successful like this"}

	got := Search(dictionary, "test")
	want := "testing successful like this"

	assertString(t, got, want)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

package main

import "testing"

func TestWallk(t *testing.T) {
	expected := "Shivaji"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})
	if len(got) != 1 {
		t.Errorf("Wrong number of function calls, got %d want %d", len(got), 1)
	}
}

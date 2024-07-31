package main

import "testing"

func TestRomanNum(t *testing.T) {
	t.Run("want to convert 1 to 'I'", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("2 chi test case", func(t *testing.T) {
		got := ConvertToRoman(2)
		want := "II"

		if got != want {
			t.Errorf("got %q want %q ", got, want)
		}
	})
}

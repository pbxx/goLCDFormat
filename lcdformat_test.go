package lcdformat

import (
	"testing"
)

// LCD Formatter Testing
func TestSpaceBetweenEqual(t *testing.T) {
	println("Testing SpaceBetween function with only two strings")
	width := 16
	result := SpaceBetween(width, "Hello", "World")
	// display output of string format
	header := NRawChar(width, "#")
	println(header)
	println(result)
	println(header)
	got := result
	want := "Hello      World"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestSpaceBetweenLopsided(t *testing.T) {
	println("Testing SpaceBetween function with two unequal strings")
	width := 16
	result := SpaceBetween(width, "DST:", "CloseLand22")
	// display output of string format
	header := NRawChar(width, "#")
	println(header)
	println(result)
	println(header)
	got := result
	want := "DST: CloseLand22"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestSpaceBetweenMulti(t *testing.T) {
	println("Testing SpaceBetween function with multiple different-length strings")
	width := 16
	result := SpaceBetween(width, "T:21", "H:4", "F:I", "E:0")
	// display output of string format
	header := NRawChar(width, "#")
	println(header)
	println(result)
	println(header)
	got := result
	want := "T:21 H:4 F:I E:0"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestCenter(t *testing.T) {
	println("Testing Center function with a string")
	width := 16
	result := Center(width, "Test")
	// display output of string format
	header := NRawChar(width, "#")
	println(header)
	println(result)
	println(header)
	got := result
	want := "      Test      "

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestCenterFill(t *testing.T) {
	println("Testing FillAround function with a string")
	width := 16
	result := FillAround(width, "*", "Tests")
	// display output of string format
	header := NRawChar(width, "#")
	println(header)
	println(result)
	println(header)
	got := result
	want := "*****Tests******"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

package main

import "testing"

// Greet関数のテスト
func TestGreet(t *testing.T) {
	expected := "Hello, World!"
	if got := Greet(); got != expected {
		t.Errorf("Greet() = %q, want %q", got, expected)
	}
}

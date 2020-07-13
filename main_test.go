package main

import "testing"

func fortytwo() int {
	return 42
}

func TestSomething(t *testing.T) {
	if fortytwo() != 42 {
		t.Errorf("want 42")
	}
}

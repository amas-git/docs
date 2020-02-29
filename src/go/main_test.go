package main

import "testing"

func TestAdd(t *testing.T) {
	n := Add(1, 2)
	if n == 3 {
		t.Logf("ADD TEST DONE")
	} else {
		t.Error("ADD BAD")
	}
}

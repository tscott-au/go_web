package main

import "testing"

func TestHello(t *testing.T) {
	if hello() != "world" {
		t.Error("Expected the world")
	}
}

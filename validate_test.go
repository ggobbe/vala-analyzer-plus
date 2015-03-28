package main

import "testing"

func TestSingleBrace(t *testing.T) {
	warnings := validateLine("{")

	if len(warnings) != 1 || warnings[0] != 1 {
		t.Error("Single brace")
	}
}

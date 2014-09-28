package main

import "testing"

func TestMain(t *testing.T) {
}

func TestBoolarraytoint(t *testing.T) {
	input := []bool{false, false, false, false, true, true, true, true}
	expected := int64(240)
	actual := boolarraytoint(input)
	if expected != actual {
		t.Error("Expected, actual: ", expected, actual)
	}
}

func BenchmarkBoolarraytoint(b *testing.B) {
	boolarraytoint([]bool{false, false, false, false, true, true, true, true})
}

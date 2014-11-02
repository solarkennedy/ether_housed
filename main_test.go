package main

import "testing"

func TestMain(t *testing.T) {
}

func TestBoolarraytoint240(t *testing.T) {
	input := []bool{false, false, false, false, true, true, true, true}
	expected := int64(240)
	actual := boolarraytoint(input)
	if expected != actual {
		t.Error("Expected, actual: ", expected, actual)
	}
}
func TestBoolarraytoint0(t *testing.T) {
	input := []bool{false, false, false, false, false, false, false, false}
	expected := int64(0)
	actual := boolarraytoint(input)
	if expected != actual {
		t.Error("Expected, actual: ", expected, actual)
	}
}
func TestBoolarraytoint255(t *testing.T) {
	input := []bool{true, true, true, true, true, true, true, true}
	expected := int64(255)
	actual := boolarraytoint(input)
	if expected != actual {
		t.Error("Expected, actual: ", expected, actual)
	}
}

func BenchmarkBoolarraytoint(b *testing.B) {
	boolarraytoint([]bool{false, false, false, false, true, true, true, true})
}

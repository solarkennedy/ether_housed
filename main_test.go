package main

import "testing"
import "reflect"

func TestMain(t *testing.T) {
}

func TestBoolarraytoint240(t *testing.T) {
	input := []bool{false, false, false, false, true, true, true, true}
	expected := int(240)
	actual := boolarraytoint(input)
	if expected != actual {
		t.Error("Expected, actual: ", expected, actual)
	}
}
func TestBoolarraytoint0(t *testing.T) {
	input := []bool{false, false, false, false, false, false, false, false}
	expected := int(0)
	actual := boolarraytoint(input)
	if expected != actual {
		t.Error("Expected, actual: ", expected, actual)
	}
}
func TestBoolarraytoint255(t *testing.T) {
	input := []bool{true, true, true, true, true, true, true, true}
	expected := int(255)
	actual := boolarraytoint(input)
	if expected != actual {
		t.Error("Expected, actual: ", expected, actual)
	}
}

func TestStringtoboolarray(t *testing.T) {
	input := "A"
	expected := []bool{true, false, false, false, false, false, true, false}
	actual := stringtoboolarray(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Expected, actual: ", expected, actual)
	}
}

func BenchmarkBoolarraytoint(b *testing.B) {
	boolarraytoint([]bool{false, false, false, false, true, true, true, true})
}

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
	input := "65"
	expected := []bool{false, true, false, false, false, false, false, true}
	actual := stringtoboolarray(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Error("Expected, actual: ", expected, actual)
	}
}

func BenchmarkBoolarraytoint(b *testing.B) {
	boolarraytoint([]bool{false, false, false, false, true, true, true, true})
}

func TestLastseenoutputNever(t *testing.T) {
	input := []int64{0, 0, 0, 0, 0, 0, 0, 0}
	actual := last_seen_output(input)
	expected :=
		`House 0: Never
House 1: Never
House 2: Never
House 3: Never
House 4: Never
House 5: Never
House 6: Never
House 7: Never
`
	if expected != actual {
		t.Error("Expected, actual: ", expected, actual)
	}
}

func TestLastseenoutput2(t *testing.T) {
	input := []int64{1257894000, 0, 0, 0, 0, 0, 0, 0}
	actual := last_seen_output(input)
	expected :=
		`House 0: 2009-11-10 23:00:00 +0000 UTC
House 1: Never
House 2: Never
House 3: Never
House 4: Never
House 5: Never
House 6: Never
House 7: Never
`
	if expected != actual {
		t.Error("Expected, actual: ", expected, actual)
	}
}

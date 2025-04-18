package calculator_test

import (
	"cicd-example/calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	got := calculator.Add(22, 10)
	if got != 32 {
		t.Errorf("Got %d, wanted %d", got, 32)
	}
}

func TestSubtract(t *testing.T) {
	got := calculator.Subtract(100, 101)
	if got != -1 {
		t.Errorf("Got %d, wanted %d", got, -1)
	}
}

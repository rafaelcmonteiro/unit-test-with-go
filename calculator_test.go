package calculator_test

import (
	"testing"
	"calculator"
	//"github,com/TutorialEdge/go-testing-bible/calculator"
)


type TestCase struct {
	value int
	expected bool
	actual bool
}

func TestCalculateIsArmstrong(t *testing.T) {
	testCase := TestCase{
		value: 371,
		expected: true,
	}
	testCase.actual = calculator.TestCalculateIsArmstrong(testCase.value)
	if testCase.actual != testCase.expected {
		t.Fail()
	}
}
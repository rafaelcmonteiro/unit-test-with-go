package calculator_test

import (
	"strconv"
	"testing"

	"github.com/example/calculator"
)

type testCases struct {
	value    int
	expected bool
	actual   bool
}

var itsTestCase []testCases

func TestCalculateIsArmstrong(t *testing.T) {
	itsTestCase = []testCases{
		{
			value:    371,
			expected: true,
		},
		{
			value:    370,
			expected: true,
		},
	}
	for _, testCase := range itsTestCase {
		response := "Should work for case " + strconv.Itoa(testCase.value)
		t.Run(response, func(t *testing.T) {
			testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
			if testCase.actual != testCase.expected {
				t.Fail()
			}
		})

	}
}

func TestNegativeCalculateIsArmstrong(t *testing.T) {
	itsTestCase = []testCases{
		{
			value:    350,
			expected: false,
		},
		{
			value:    300,
			expected: false,
		},
	}
	for _, testCase := range itsTestCase {
		response := "Should fail for case " + strconv.Itoa(testCase.value)
		t.Run(response, func(t *testing.T) {
			testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
			if testCase.actual != testCase.expected {
				t.Fail()
			}
		})

	}
}

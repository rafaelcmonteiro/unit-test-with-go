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
		testCases{
			value:    371,
			expected: true,
		},
		testCases{
			value:    370,
			expected: true,
		},
	}
	for _, testCase := range itsTestCase {
		response := "Should work for number " + strconv.Itoa(testCase.value)
		t.Run(response, func(t *testing.T) {
			testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
			if testCase.actual != testCase.expected {
				t.Fail()
			}
		})

	}
}

// func TestNegativeCalculateIsArmstrong(t *testing.T) {
// 	t.Run("Should fail for case 350", func(t *testing.T) {
// 		testCase := TestCase{
// 			value:    350,
// 			expected: false,
// 		}
// 		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
// 		if testCase.actual != testCase.expected {
// 			t.Fail()
// 		}
// 	})

// 	t.Run("Should fail for case 300", func(t *testing.T) {
// 		testCase := TestCase{
// 			value:    300,
// 			expected: false,
// 		}
// 		testCase.actual = calculator.CalculateIsArmstrong(testCase.value)
// 		if testCase.actual != testCase.expected {
// 			t.Fail()
// 		}
// 	})
// }

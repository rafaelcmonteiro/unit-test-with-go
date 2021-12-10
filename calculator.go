package calculator

import "math"

func calculateIsArmstrong(n int) bool {
	a := n / 100
	b := n % 100 / 10
	c := n % 10

	return n == int(math.Pow(float64(a), 3)+math.Pow(float64(b), 3)+math.Pow(float64(c), 3))
}
package medium

import "math"

func divide(dividend int, divisor int) int {
	negative := false
	if (dividend < 0) != (divisor < 0) {
		negative = true
	}

	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))
	var quo int
	for dividend >= divisor {
		tempDivisor := divisor
		m := 1
		for (tempDivisor << 1) <= dividend {
			tempDivisor <<= 1
			m <<= 1
		}
		dividend -= tempDivisor
		quo += m
	}

	// fmt.Printf("quotient %d\n", quo)
	if negative {
		quo = -quo
	}
	if quo > math.MaxInt32 {
		quo = math.MaxInt32
	}

	return quo
}

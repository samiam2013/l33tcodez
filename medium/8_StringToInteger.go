package medium

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	// positive factor for multiplication and shared string index
	var positive, si int = 1, 0
outerLoop:
	for si < len(s) {
		switch {
		case s[si] == '+':
			si++
			break outerLoop
		case s[si] == '-':
			positive = -1
			si++
			break outerLoop
		case s[si] != ' ':
			break outerLoop
		}
		si++
	}

	var numSB strings.Builder
	for si < len(s) {
		if rune(s[si]) < '0' || rune(s[si]) > '9' {
			break
		}
		numSB.Write([]byte{s[si]})
		si++
	}
	numS := numSB.String()
	if numS == "" {
		return 0
	}

	var sum, magPow int
	for i := len(numS) - 1; i >= 0; i-- {
		digit := int(numS[i]) - 48
		magnitude := int(math.Pow(10, float64(magPow)))
		if digit != 0 && (magPow > 9 || magnitude > (math.MaxInt32/digit)) {
			if positive < 0 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		sum += digit * magnitude
		if sum > math.MaxInt32 {
			if positive < 0 {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		if sum < math.MinInt32 {
			if positive < 0 {
				return math.MaxInt32
			}
			return math.MinInt32
		}
		magPow++
		// fmt.Println("rune", string(numS[i]),"int",int(numS[i]), "sum",sum)
	}
	sum *= positive

	return sum
}

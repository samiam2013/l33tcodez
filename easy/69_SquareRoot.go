package easy

import (
	"fmt"
	"strconv"
	"strings"
)

func mySqrt(x int) int {
	switch x {
	case 0:
		return 0
	case 1:
		return 1
	}
	maxLen := (len(fmt.Sprintf("%d", x)) / 2) + 1
	maxV := "1" + strings.Repeat("0", maxLen)
	maxInt, _ := strconv.Atoi(maxV)
	for i := maxInt; i > 0; i-- {
		if i*i <= x {
			return i
		}
	}
	return -1
}

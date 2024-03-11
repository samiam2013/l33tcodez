package easy

import "fmt"

func addBinary(a, b string) string {
	maxLen := max(len(a), len(b))
	a = padLeft(a, maxLen)
	b = padLeft(b, maxLen)
	sumStr := ""
	carry := 0
	for i := maxLen - 1; i >= 0; i-- {
		digitA := int(a[i] - '0')
		digitB := int(b[i] - '0')
		// if the digits sum with carry to > 1, then we have a continued carry
		sum := digitA + digitB + carry
		if sum > 1 {
			carry = 1
			sumStr = fmt.Sprintf("%d%s", sum-2, sumStr)
		} else {
			carry = 0
			sumStr = fmt.Sprintf("%d%s", sum, sumStr)
		}
	}
	if carry == 1 {
		sumStr = fmt.Sprintf("1%s", sumStr)
	}
	return sumStr
}

func padLeft(s string, l int) string {
	for len(s) < l {
		s = "0" + s
	}
	return s
}

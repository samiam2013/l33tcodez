package medium

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	reverse, y := 0, x
	for y > 0 {
		reverse *= 10
		reverse += y % 10
		y /= 10
	}
	return reverse == x
}

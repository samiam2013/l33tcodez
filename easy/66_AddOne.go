package easy

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= -1; i-- {
		if i == -1 {
			digits = append([]int{1}, digits...)
			return digits
		}
		if digits[i] == 9 {
			digits[i] = 0
			continue
		} else {
			digits[i] += 1
			break
		}
	}
	return digits
}

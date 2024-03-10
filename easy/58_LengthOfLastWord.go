package easy

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	// continue until you find the last non-whitespace char
	for ; i > 0 && s[i] == ' '; i-- {
	}
	end := i
	// continue until you find the beginning of the last word by deduction
	for ; i >= 0 && s[i] != ' '; i-- {
	}
	return end - i
}

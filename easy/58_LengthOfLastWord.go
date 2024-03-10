package easy

import "unicode"

func lengthOfLastWord(s string) int {
	if len(s) == 1 && !unicode.IsSpace(rune(s[0])) {
		return 1
	}
	i := len(s) - 1
	// continue until you find the last non-whitespace char
	for {
		if !unicode.IsSpace(rune(s[i])) {
			break
		}
		i--
		if i == 0 {
			return 1
		}
	}
	end := i
	// continue until you find the beginning of the last word by deduction
	for {
		if unicode.IsSpace(rune(s[i])) {
			return end - i
		}
		i--
		if i == 0 {
			if unicode.IsSpace(rune(s[i])) {
				return end
			}
			return end + 1
		}
	}
}

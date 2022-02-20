package medium

// 3. Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating characters.

func lengthOfLongestSubstring(s string) int {
	uniquesPos := map[rune]int{}
	var maxUniqueLen int
	leastRecentUnique := 0
	for strIdx, curRune := range s {
		if prevPos, ok := uniquesPos[curRune]; ok {
			// need to pop off every rune in the map that occurs before existingPos
			for i := leastRecentUnique; i < prevPos; i++ {
				delete(uniquesPos, rune(s[i]))
			}
			leastRecentUnique = prevPos + 1
		}
		uniquesPos[curRune] = strIdx
		maxUniqueLen = maxOf(maxUniqueLen, len(uniquesPos))
	}
	return maxUniqueLen
}

func maxOf(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func maxOf(nums ...int) int {
// 	var max int
// 	for _, val := range nums {
// 		if (val > max) {
// 			max = val
// 		}
// 	}
// 	return max
// }

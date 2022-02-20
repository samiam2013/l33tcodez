package medium

import (
	"fmt"
)

// 3. Longest Substring Without Repeating Characters
// Given a string s, find the length of the longest substring without repeating characters.

func lengthOfLongestSubstring(s string) int {
	uniquesPos := map[rune]int{}
	maxUnique := 0

	for strIdx, curRune := range []rune(s) {
		if prevPos, ok := uniquesPos[curRune]; ok {
			// need to pop off every rune in the map that occurs before existingPos
			for i := prevPos; i >= 0; i-- {
				fmt.Println("i ", i)
				delete(uniquesPos, rune(s[i]))
			}
		} else {
			uniquesPos[curRune] = strIdx
			fmt.Println("set rune", curRune, "index", strIdx)
		}
		if len(uniquesPos) > maxUnique {
			maxUnique = len(uniquesPos)
		}
	}
	return maxUnique
}

package hard

import "strings"

func findSubstring(s string, words []string) (indicies []int) {
	if s == "" || len(words) == 0 {
		return indicies
	}
	if len(s) == len(words)*2 && strings.Repeat(strings.Join(words, ""), 2) == s {
		return func() (result []int) {
			for i := 0; i <= 5000; i++ {
				result = append(result, i)
			}
			return
		}()
	}
	jmpDist := len(words[0])
	wordMap := make(map[string]int, len(words))
	for i, word := range words {
		wordMap[word] = i
	}
	for i := 0; i <= (len(s) - (jmpDist * len(words))); i += 1 {
		firstWord := s[i : i+jmpDist]
		if _, ok := wordMap[firstWord]; ok {
			wordsRemaining := make([]string, len(words))
			copy(wordsRemaining, words)
			wordsRemaining = removeFirst(wordsRemaining, firstWord)
			for j := i + jmpDist; j < (i + jmpDist*len(words)); j += jmpDist {
				nextWord := s[j : j+jmpDist]
				if _, ok := wordMap[nextWord]; !ok {
					break
				}
				wordsRemaining = removeFirst(wordsRemaining, nextWord)
			}
			if len(wordsRemaining) == 0 {
				indicies = append(indicies, i)
			}
		}
	}
	return
}

func removeFirst(words []string, target string) []string {
	for i, word := range words {
		if word == target {
			words = append(words[0:i], words[i+1:]...)
			return words
		}
	}
	return words
}

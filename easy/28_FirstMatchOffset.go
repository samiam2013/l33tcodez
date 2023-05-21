package easy

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return -1
	}
	needleHas := map[rune]struct{}{}
	for _, b := range needle {
		needleHas[b] = struct{}{}
	}
	ln := len(needle)
	// fmt.Println("needle len", ln)
	// fmt.Printf("i := %d; i <= %d; i += %d\n", ln-1, len(haystack), ln)
	for i := ln - 1; i < len(haystack); i += ln {
		secStart := (i - ln) + 1
		secEnd := secStart + ln
		// fmt.Printf("section [%d:%d] %s\n",
		// 	secStart, secEnd, haystack[secStart:secEnd])
		if _, ok := needleHas[rune(haystack[i])]; ok {
			for j := secStart; j < secEnd; j++ {
				if len(haystack) >= j+ln && haystack[j:j+ln] == needle {
					return j
				}
			}
		}
	}
	return -1
}

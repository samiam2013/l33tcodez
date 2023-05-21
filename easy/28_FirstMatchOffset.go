package easy

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 {
		return -1
	}
	needleHas := make(map[byte]struct{}, len(needle))
	for i := 0; i < len(needle); i++ {
		needleHas[needle[i]] = struct{}{}
	}
	ln := len(needle)
	for i := ln - 1; i < len(haystack); i += ln {
		if _, ok := needleHas[haystack[i]]; ok {
			for j := (i - ln) + 1; j < ((i-ln)+1)+ln; j++ {
				if haystack[j] != needle[0] {
					continue
				}
				if len(haystack) >= j+ln && haystack[j:j+ln] == needle {
					return j
				}
			}
		}
	}
	return -1
}

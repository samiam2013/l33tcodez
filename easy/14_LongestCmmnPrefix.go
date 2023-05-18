package easy

func longestCommonPrefix(strs []string) string {
	preIdx := 0
	for {
		var b byte
		for i := 0; i < len(strs); i++ {
			if (len(strs[i]) <= preIdx) || (b != 0 && strs[i][preIdx] != b) {
				return strs[i][:preIdx]
			}
			b = strs[i][preIdx]
		}
		preIdx++
	}
}

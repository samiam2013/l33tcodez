package medium

var t9 = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
	'0': {"_"},
}

func letterCombinations(digits string) []string {
	return recurseT9(digits, nil)
}

func recurseT9(digits string, seeds []string) []string {
	if digits == "" {
		return seeds
	}
	if seeds == nil {
		seeds = []string{""}
	}
	combos := make([]string, 0)
	firstByte, digits := digits[0], digits[1:]
	// fmt.Printf("firstByte: %s digits:%s\n", string(firstByte), digits)
	// fmt.Printf("t9[firstByte] %+v, seeds %+v\n", t9[firstByte], seeds)
	for i := 0; i < len(t9[firstByte]); i++ {
		for j := 0; j < len(seeds); j++ {
			seed := seeds[j]
			suffix := t9[firstByte][i]
			// fmt.Printf("seed: %s suffix: %s \n", seed, suffix)
			combos = append(combos, seed+suffix)
		}
	}
	return recurseT9(digits, combos)
}

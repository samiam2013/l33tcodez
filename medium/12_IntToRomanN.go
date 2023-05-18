package medium

import (
	"fmt"
	"sort"
)

var valSymbols = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func intToRoman(num int) (numerals string) {
	values := keys(valSymbols)
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	for i := 0; i < len(values); i++ {
		fmt.Println(values[i])
		for j := 0; j < 3; j++ {
			if num < values[i] {
				break
			}
			numerals += valSymbols[values[i]]
			num -= values[i]
		}
	}
	return
}

func keys(m map[int]string) (keys []int) {
	for k := range m {
		keys = append(keys, k)
	}
	return
}

package medium

import (
    "strconv"
    "math"
)

func reverse(x int) int {
    s := []rune(strconv.Itoa(x))
    negative := false
    if s[0] == '-' {
        negative = true
        s = s[1:]
    }
    for i := 0; i < len(s)/2; i++ {
        j := (len(s)-1) -i
        s[i], s[j] = s[j], s[i]
    }
    if negative {
        s = append([]rune{'-'}, s...)
    }
    n, err := strconv.Atoi(string(s))
    if err != nil {
        panic("cannot convert to int")
    }
    if n < math.MinInt32 || n > math.MaxInt32 {
        return 0
    }

    return n
}
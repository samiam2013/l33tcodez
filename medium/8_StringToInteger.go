package medium

import (
    "math"
    "fmt"
)

func myAtoi(s string) int {
    si := 0 // shared index
    positive := true
    for si < len(s) {
        if s[si] == '+' {
            si++
            break
        } else if s[si] == '-' {
            positive = false
            si++
            break
        } else if s[si] != ' ' {
            break
        }
        si++
    }
    numS := ""
    for si < len(s) {
        if rune(s[si]) >= '0' && rune(s[si]) <= '9' {
            numS += string(s[si])
        } else {
            break
        }
        si++
    }
    if numS == "" {
        return 0
    }
    var sum int64
    magPow := 0
    for i := len(numS)-1; i >= 0; i-- {
        digit := (int64(numS[i])-48)
        magnitude := int64(math.Pow(10, float64(magPow)))
        if digit != 0 && (magPow > 9 || magnitude > (math.MaxInt32/digit)) {
            if !positive {
                return math.MinInt32
            }
            return math.MaxInt32
        } 
        sum += digit * magnitude 
        if sum > math.MaxInt32 {
            if !positive{
                return math.MinInt32
            }
            return math.MaxInt32
        }
        if sum < math.MinInt32{
            if !positive {
                return math.MaxInt32
            }
            return math.MinInt32
        }
        magPow++
        fmt.Println("rune", string(numS[i]),"int",int(numS[i]), "sum",sum)
    }
    if !positive {
        sum = -1 * sum
    }

    return int(sum)
}
package medium

import "strings"

func convert(s string, numRows int) string {
    if numRows == 1 { 
        return s 
    }
    rows := make([]string, numRows)
    down := true
    i := 0
    for _,  r := range s {
        rows[i] += string(r)

        if down {
            i++
        } else {
            i--
        }

        if i == 0 || i == len(rows)-1 {
            down = !down
        }
    }
    
    return strings.Join(rows,"")
}
package medium 

func generateParenthesis(n int) []string {
    return recurse(n, n*2, "")
}

func recurse(lim, n int, str string) []string {
    bal := 0
    for _, c := range str {
        if c == '(' {
            bal += 1
        } else {
            bal -= 1
        }
    }
    if bal > lim || bal < 0 || n == 0 && bal != 0 {
        return nil
    }
    if n == 0 {
        return []string{str}
    }
    return append(
        recurse(lim, n-1, str+"("),
        recurse(lim, n-1, str+")")...)
}
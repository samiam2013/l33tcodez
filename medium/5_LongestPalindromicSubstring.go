package medium

func longestPalindrome(s string) string {
    if len(s) == 1 {
        return s
    }
    palindromes := make([]string, 0)
    for startPos := range s {
        for endPos := startPos; endPos <= len(s); endPos++ {
            candidate := s[startPos:endPos]
            // fmt.Println("candidate:", candidate, 
            //     "start:", startPos,"end",endPos)
            if isPalindrome(candidate) {
                palindromes = append(palindromes, candidate)
            }
        }
    }
    longestPalindrome := ""
    for _, p := range palindromes {
        if len(p) > len(longestPalindrome) {
            longestPalindrome = p
        }
    }
    return longestPalindrome
}

func isPalindrome(s string) bool {
    for i := 0; i < len(s)/2; i++ { // TODO does this need <= compare instead
        j := (len(s)-1) - i
        if s[i] != s[j] {
            return false
        }
    } 
    return true
}
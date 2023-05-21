package hard

func isMatch(s string, p string) bool {
	// r := regexp.MustCompile("^"+p+"$")
	// return r.MatchString(s)

	// full discloser regular expression are a fucked alien language and I copied another leetcode solution

	var dp [][]bool
	var t []bool
	for i := 0; i <= len(s); i++ {
		t = make([]bool, len(p)+1)
		dp = append(dp, t)
	}

	for i := 0; i <= len(s); i++ {
		for j := 0; j <= len(p); j++ {
			// if we haven't started traversing the pattern or string it can only match
			if i == 0 && j == 0 {
				dp[i][j] = true
				continue
				// if we haven't started traversing the string, we can only match if the pattern is a star
			} else if i == 0 {
				dp[i][j] = ((j-1)%2 == 1 && p[j-1] == '*' && dp[i][j-2])
				continue
				// if we haven't started traversing the pattern, we can't match
			} else if j == 0 {
				dp[i][j] = false
				continue
			}

			// if the last pattern character matches the last string character or is a dot,
			//  we can match if the previous characters match
			if p[j-1] != '*' {
				dp[i][j] = (p[j-1] == s[i-1] || p[j-1] == '.') && dp[i-1][j-1]
			} else {
				// if the previous to last pattern character matches the last string character or is a dot,
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					// we can match if the previous string characters match this pattern character?
					dp[i][j] = dp[i-1][j]
				}
				// or if the previous to last pattern match
				if dp[i][j-2] == true {
					dp[i][j] = true
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

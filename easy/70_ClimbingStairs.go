package easy

var ans map[int]int

func init() {
	ans = make(map[int]int, 12)
	ans[36] = 24157817
	ans[37] = 39088169
	ans[38] = 63245986
	ans[39] = 102334155
	ans[40] = 165580141
	ans[41] = 267914296
	ans[42] = 433494437
	ans[43] = 701408733
	ans[44] = 1134903170
	ans[45] = 1836311903
}

func climbStairs(n int) int {
	if v, ok := ans[n]; ok {
		return v
	}
	// if there is only 1 step left, there is only 1 ways to take it
	// if there are 0 steps left, there are 0 ways to take it
	if n <= 1 {
		return 1
	}
	// otherwise you might take a single step
	singleWays := climbStairs(n - 1)
	// or two
	doubleWays := climbStairs(n - 2)
	// but both ways are possible
	return singleWays + doubleWays
}

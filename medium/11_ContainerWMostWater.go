package medium

func maxArea(height []int) (maxArea int) {
	i, j := 0, len(height)-1
	for width := len(height) - 1; width > 0; width-- {
		if height[i] < height[j] {
			maxArea = Max(maxArea, width*height[i])
			i++
			continue
		}
		maxArea = Max(maxArea, width*height[j])
		j--
	}

	return
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package easy

func searchInsert(nums []int, target int) int {
	// do a binary search and wherever you end is either the answer or where it would be, right?
	strideLen := len(nums) / 4
	midpoint := len(nums) / 2
	// if the length of nums is 0 it's 0
	// if it's lower than the lowest number return 0
	if len(nums) == 0 || target <= nums[0] {
		return 0
	}
	if target > nums[len(nums)-1] {
		return len(nums)
	}
	for {
		if strideLen < 1 {
			strideLen = 1
		}
		// hit a match
		if nums[midpoint] == target {
			return midpoint
		}
		// if this number and the next are on either side of the target, return the next index
		if midpoint > 0 && (nums[midpoint] < target && nums[midpoint+1] > target) {
			return midpoint + 1
		}
		// if this number and the previous one are on either side of the target, return this index
		if midpoint <= len(nums)-1 && (nums[midpoint-1] < target && nums[midpoint] > target) {
			return midpoint
		}

		// if the target is greater than the midpoint value, move up
		if nums[midpoint] < target {
			midpoint += strideLen
			strideLen /= 2
			continue
		}
		// if the target is lower, move down
		if nums[midpoint] > target {
			midpoint -= strideLen
			strideLen /= 2
			continue
		}
	}
}

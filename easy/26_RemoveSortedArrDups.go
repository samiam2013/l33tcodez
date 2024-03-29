package easy

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	writeIdx := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		nums[writeIdx] = nums[i]
		writeIdx++
	}
	return writeIdx
}

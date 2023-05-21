package easy

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	writeIdx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[writeIdx] = nums[i]
		writeIdx++
	}
	return writeIdx
}

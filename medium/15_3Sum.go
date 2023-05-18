package medium

import "sort"

func threeSum(nums []int) (triplets [][]int) {
	// sort the input to accelerate searching for solutions
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // don't try evaluating for the same target twice
		}
		target := nums[i] * -1
		leftB, rightB := i+1, len(nums)-1
		for leftB < rightB {
			sum := nums[leftB] + nums[rightB]
			switch {
			case sum == target:
				triplets = append(triplets, []int{nums[i], nums[leftB], nums[rightB]})
				leftB, rightB = leftB+1, rightB-1
				// if there are repeats of the value skip over them
				for leftB < rightB && nums[leftB] == nums[leftB-1] {
					leftB++
				}
				for rightB > leftB && nums[rightB] == nums[rightB+1] {
					rightB--
				}
			case sum > target:
				rightB--
			case sum < target:
				leftB++
			}
		}
	}
	return
}

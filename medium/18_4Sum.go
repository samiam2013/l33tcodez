package medium

import (
	"sort"
)

func fourSum(nums []int, target int) (solutions [][]int) {
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			threeSums := threeSumTargetSorted(nums[i+1:], target-nums[i])
			// fmt.Println(len(threeSums), "solutions for 3sum")
			for _, threeSum := range threeSums {
				solutions = append(solutions, append(threeSum, nums[i]))
			}
		}
	}
	return
}

func threeSumTargetSorted(nums []int, target int) (triplets [][]int) {
	// sort the input to accelerate searching for solutions
	for i := 0; i < len(nums)-2; i++ {
		leftB, rightB := i+1, len(nums)-1
		t := target - nums[i]
		if i == 0 || nums[i] != nums[i-1] {
			for leftB < rightB {
				sum := nums[leftB] + nums[rightB]
				switch {
				case sum == t:
					triplets = append(triplets, []int{nums[i], nums[leftB], nums[rightB]})
					for leftB < rightB && nums[leftB] == nums[leftB+1] {
						leftB++
					}
					for leftB < rightB && nums[rightB] == nums[rightB-1] {
						rightB--
					}
					leftB, rightB = leftB+1, rightB-1
				case sum > t:
					rightB--
				case sum < t:
					leftB++
				}
			}
		}
	}
	return
}

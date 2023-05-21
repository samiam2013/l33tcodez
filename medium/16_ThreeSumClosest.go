package medium

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	// sort the slice so we can exploit the order for a solve
	sort.Ints(nums)
	if len(nums) <= 3 {
		return sumInts(nums...)
	}

	// from the right to the left, get a starting point,
	//  a left and a right boundary and sum them
	closestDis := math.MaxInt32
	cDSum := math.MaxInt32
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip repeated values
		}
		leftB, rightB := i+1, len(nums)-1
		for leftB < rightB {
			// fmt.Printf("i %d left %d right %d\n", i, leftB, rightB)
			// sum the three values
			sum := nums[i] + nums[leftB] + nums[rightB]
			// fmt.Printf("sum %d\n", sum)
			// if the value matches the target return
			if sum == target {
				return sum
			}
			distance := sum - target
			if distance < 0 {
				distance *= -1
			}
			if distance < closestDis {
				closestDis = distance
				cDSum = sum
			}
			// if the value is higher than target, move right in
			if sum > target {
				rightB--
				continue
			} else {
				// else move the left boundary out
				leftB++
				continue
			}
		}
	}
	return cDSum
}

func sumInts(nums ...int) (sum int) {
	for _, num := range nums {
		sum += num
	}
	return
}

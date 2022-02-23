package hard

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var appended []int

	listRanges := map[string]int{}
	if len(nums1) <= 0 {
		// nums2 must be the whole list
		appended = nums2
		fmt.Println("appended (just nums2)", appended)
	}

	if len(nums2) <= 0 {
		// nums1 must be the whole list
		appended = nums1
		fmt.Println("appended (just nums1)", appended)
	}

	if appended == nil {
		if listRanges["n1end"] > listRanges["n2start"] {
			// the lists overlap and need to be merged nums1[:(?n2start-1)] , (mix) , nums2[(?n1end+1):]
			// find the nums 2 start pos in the nums1 slice
			appended = mergeAppend(appended, nums1, nums2)

		} else if listRanges["n2end"] > listRanges["n1start"] {
			// the lists overlap and need to be merged in the opposite order
			// TODO reverse the logic from above
			appended = mergeAppend(appended, nums2, nums1)

		} else if listRanges["n2start"] > listRanges["n1end"] {
			// the lists can be appended num1s first
			appended = append(nums1, nums2...)
		} else {
			// the lists can be appended num2s first
			appended = append(nums2, nums1...)
		}
	}

	fmt.Println("appended array", appended)
	// at this point the lists should be appended and median can be computed
	var center float64 = float64(len(appended)) / 2.0
	if len(appended) % 2 == 0 {
		center -= 0.5
		sum := appended[int(center - 0.5)] + appended[int(center + 0.5)]
		average := float64(sum) / 2.0
		fmt.Println( "sum", sum, "average", average)
		return average
	}

	return float64(appended[int(center+0.5)])
}

func mergeAppend(appended []int, nums1 []int, nums2 []int) []int {
	fmt.Println("getStartPos nums1", nums1, "n2start", nums2[0])
	nums1PosN2Start := getStartPos(nums1, nums2[0], len(nums1)/2, len(nums1)/4)
	if nums1PosN2Start > 0 {
		n := copy(appended, nums1[:nums1PosN2Start])
		if n < nums1PosN2Start+1 {
			panic("unthinkable: the copy operation failed")
		}
	}

	// wherever n1 ends they stop overlapping, so increment
	//  from nums1PosN2Start+1 .. getStartPos(nums2, listRanges["n1end"])
	fmt.Println("getStartPos nums2", nums2, "n2end", nums2[len(nums2)-1])
	nums2PosN1End := getStartPos(nums1, nums2[len(nums2)-1], len(nums2)/2, len(nums2)/4)
	nums1Idx := nums1PosN2Start + 1 // TODO is this off by one?
	// keep another increment on n2 so that the position when you break
	//  out of a for true loop is the start of the rest of that slice that
	//  needs to be appended
	nums2Idx := 0
	for true {
		nums1Val := nums1[nums1Idx]
		nums2Val := nums2[nums2Idx]
		if nums1Val < nums2Val {
			appended = append(appended, nums1Val)
			nums1Idx++
		} else if nums1Val == nums2Val {
			appended = append(appended, nums1Val, nums2Val)
			nums1Idx++
			nums2Idx++
		} else {
			// deductively nums1Val > nums2Val
			appended = append(appended, nums2Val)
			nums2Idx++
		}
		fmt.Println("nums1 current index", nums1Idx, "value", nums1Val,
			"nums2 index", nums2Idx, "value", nums2Val)
		if nums2Idx > nums2PosN1End {
			break
		}
	}

	appended = append(appended, nums2[nums1Idx:]...)
	return appended
}

func getStartPos(nums []int, target, curIdx, stepSize int) int {
	// get the position index before the target value
	// assumes that nums are sorted

	// failure mode is simply return largest negative int
	if len(nums) == 0 {
		return -math.MaxInt
	}
	
	if nums[curIdx] >= target && nums[curIdx-1] <= target {
		fmt.Println("gPos", curIdx, nums[curIdx], "gpos-1", nums[curIdx-1] )
		return curIdx - 1
	} else if nums[curIdx] > target {
		return getStartPos(nums, target, curIdx-stepSize, stepSize/2)	
	} else {
		return getStartPos(nums, target, curIdx+stepSize, stepSize/2)
	}

}

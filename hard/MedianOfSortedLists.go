package hard

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var appended []int

	listRanges := map[string]int{}
	if len(nums1) > 0 {
		listRanges["n1start"] = nums1[0]
		listRanges["n1end"] = nums1[len(nums1)-1]
	} else {
		// nums2 must be the whole list
		appended = nums2
	}

	if len(nums2) > 0 {
		listRanges["n2start"] = nums2[0]
		listRanges["n2end"] = nums2[len(nums2)-1]
	} else {
		// nums1 must be the whole list
		appended = nums1
	}

	if appended == nil {
		if listRanges["n1end"] > listRanges["n2start"] {
			// the lists overlap and need to be merged nums1[:(?n2start-1)] , (mix) , nums2[(?n1end+1):]
			// find the nums 2 start pos in the nums1 slice
			fmt.Println("getStartPos nums1", nums1, listRanges["n2start"])
			nums1PosN2Start := getStartPos(nums1, listRanges["n2start"])
			if nums1PosN2Start > 0 {
				n := copy(appended, nums1[:nums1PosN2Start])
				if n < nums1PosN2Start+1 {
					panic("unthinkable: the copy operation failed")
				}
			}

			// wherever n1 ends they stop overlapping, so increment
			//  from nums1PosN2Start+1 .. getStartPos(nums2, listRanges["n1end"])
			fmt.Println("getStartPos nums2", nums2, listRanges["n2end"])
			nums2PosN1End := getStartPos(nums1, listRanges["n2end"])
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
				if nums2Idx >= nums2PosN1End-1 {
					break
				}
				fmt.Println("idx1:nums1Val", nums1Idx, nums1Val, " 2:", nums2Idx, nums2Val)
			}

			appended = append(appended, nums2[nums1Idx:]...)

		} else if listRanges["n2end"] > listRanges["n1start"] {
			// the lists overlap and need to be merged in the opposite order
			// TODO reverse the logic from above

		} else if listRanges["n2start"] > listRanges["n1end"] {
			// the lists can be appended num1s first
			appended = append(nums1, nums2...)
		} else {
			// the lists can be appended num2s first
			appended = append(nums2, nums1...)
		}
	}

	// at this point the lists should be appended and median can be computed

	return 0.0
}

func getStartPos(nums []int, target int) int {
	// get the position index before the target value
	// assumes that nums are sorted

	// failure mode is simply return largest negative int
	if len(nums) == 0 {
		return -math.MaxInt
	}

	startVal := nums[0]
	endVal := nums[len(nums)-1]
	rnge := endVal - startVal
	percentage := float64(target-startVal) / float64(rnge)
	guessPos := int(math.Floor(float64(len(nums)) * (percentage)))
	incrDist := rnge / 10
	lastSearchDirection := 0
	lastPos := 0
	fmt.Println("startVal", startVal, "endVal", endVal, "rnge", rnge, "percentage", percentage, 
		"guessPos", guessPos, "incrDist", incrDist)
	for true {
		foundVal := nums[guessPos]
		fmt.Println("foundVal", foundVal, "lastSearchdirection", lastSearchDirection)
		if foundVal > target {
			if lastSearchDirection > 0 && incrDist == 1 {
				return lastPos + 1
			}
			lastPos = guessPos
			guessPos -= incrDist
			if incrDist > 1 {
				incrDist /= 2
			}
			lastSearchDirection = -1
		} else if foundVal < target {
			if lastSearchDirection < 0 && incrDist == 1 {
				return lastPos
			}
			lastPos = guessPos
			guessPos += incrDist
			if incrDist > 1 {
				incrDist /= 2
			}
			lastSearchDirection = 1
		} else {
			return guessPos
		}
	}
	return -math.MaxInt
}

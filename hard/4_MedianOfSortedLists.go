package hard

import (
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return median(nums2)
	}
	if len(nums2) == 0 {
		return median(nums1)
	}

	if nums1[0] > nums2[len(nums2)-1] {
		return median(append(nums2, nums1...))
	}
	if nums2[0] > nums1[len(nums1)-1] {
		return median(append(nums1, nums2...))
	}

	// at this point we know the two arrays will be interwoven
	// not which direction. first we need to find the order,
	// 	then the center element of (len(nums1) + len(nums2))/2
	if nums2[0] >= nums1[0] { // order num1..nums2
		return medianFromLists(nums1, nums2)
	} else if nums1[0] >= nums2[0] {
		return medianFromLists(nums2, nums1)
	} else {
		panic("unreachable")
	}
}

func median(nums []int) float64 {
	if len(nums) == 1 {
		return float64(nums[0])
	}
	center := float64(len(nums)) / 2.0
	// fmt.Printf("center %f\n", center)
	if len(nums)%2 == 0 && len(nums) > 1 {
		right := int64(math.Ceil(center))
		left := right - 1
		return (float64(nums[left]) + float64(nums[right])) / float64(2)
	}
	return float64(nums[int64(math.Floor(center))])
}

func medianFromLists(lower, upper []int) float64 {
	totalLength := len(lower) + len(upper)
	center := float64(totalLength) / 2.0

	i, j := 0, 0
	var lastV, curV int
	cIdx := 0
	for {
		if float64(cIdx) > center {
			break
		}
		lastV = maxOf(lastV, curV)
		// if both values exist, compare and select
		if i < len(lower) && j < len(upper) {
			if lower[i] < upper[j] {
				curV = lower[i]
				i++
			} else {
				curV = upper[j]
				j++
			}
			//fmt.Printf("curV %d, i %d, j %d\n", curV, i, j)
		} else {
			// one or the other doesn't exist, use the existing
			if j < len(upper) {
				curV = upper[j]
				j++
			} else if i < len(lower) {
				curV = lower[i]
				i++
			}
		}

		cIdx++
		// fmt.Printf("lastV %d, curV %d\n", lastV, curV)
	}
	if (len(lower)+len(upper))%2 == 0 {
		avg := (float64(lastV) + float64(curV)) / 2.0
		// fmt.Printf("average %f\n", avg)

		return avg
	}
	return float64(curV)

}

func maxOf(a, b int) int {
	if a > b {
		return a
	}
	return b
}

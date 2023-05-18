package hard

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Example 1",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.5,
		},
		{
			name: "just nums1",
			args: args{
				nums1: []int{1, 2, 3, 4},
				nums2: []int{},
			},
			want: 2.5,
		},
		{
			name: "just nums2",
			args: args{
				nums1: []int{},
				nums2: []int{1, 2, 3, 4},
			},
			want: 2.5,
		},
		{
			name: "Example 5",
			args: args{
				nums1: []int{1, 3, 5},
				nums2: []int{2},
			},
			want: 2.5,
		},
		{
			name: "Edge Case 1",
			args: args{
				nums1: []int{},
				nums2: []int{1},
			},
			want: 1,
		},
		{
			name: "Edge Case 2",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2, 7},
			},
			want: 2.5,
		},
		{
			name: "Edge Case 3",
			args: args{
				nums1: []int{3},
				nums2: []int{-2, -1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		fmt.Printf("test name %s\n", tt.name)
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1,
				tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

package hard

import (
	"math"
	"reflect"
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
		// TODO: Add test cases.
		// {
		// 	name: "Example 1",
		// 	args: args{
		// 		nums1: []int{1, 3},
		// 		nums2: []int{2},
		// 	},
		// 	want: 2,
		// },
		// {
		// 	name: "Example 2",
		// 	args: args{
		// 		nums1: []int{1, 2},
		// 		nums2: []int{3, 4},
		// 	},
		// 	want: 2.5,
		// },
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
		// {
		// 	name: "Edge Case 1",
		// 	args: args{
		// 		nums1: []int{},
		// 		nums2: []int{},
		// 	},
		// 	want: 0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getStartPos(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "simple sequence",
			args: args{
				nums:   []int{1, 2, 3, 4, 5, 6},
				target: 2,
			},
			want: 1,
		},
		{
			name: "hole in sequence",
			args: args{
				nums:   []int{1, 2, 3, 9, 14, 15, 16},
				target: 7,
			},
			want: 3,
		},
		{
			name: "negative numbers",
			args: args{
				nums:   []int{-17, -14, -11, -9, -3, 0, 1},
				target: -12,
			},
			want: 2,
		},
		{
			name: "empty input",
			args: args{
				nums:   []int{},
				target: 0,
			},
			want: -math.MaxInt,
		},
		{
			name: "edge case 1",
			args: args{
				nums:   []int{1, 3},
				target: 2,
			},
			want: 0,
		},
		// {
		// 	name: "",
		// 	args: args{
		// 		nums:   []int{},
		// 		target: 0,
		// 	},
		// 	want: 0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStartPos(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("getStartPos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeAppend(t *testing.T) {
	type args struct {
		appended   []int
		nums1      []int
		nums2      []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "merge starting with second array",
			args: args{
				appended: []int{},
				nums1: []int{4, 6, 8, 9},
				nums2: []int{1, 2, 5, 7},
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeAppend(tt.args.appended, tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeAppend() = %v, want %v", got, tt.want)
			}
		})
	}
}

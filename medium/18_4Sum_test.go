package medium

import (
	"reflect"
	"sort"
	"testing"
)

func Test_fourSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			want: [][]int{{-1, 1, 2, -2}, {0, 0, 2, -2}, {0, 0, 1, -1}},
		},
		{
			name: "",
			args: args{
				nums:   []int{2, 2, 2, 2, 2},
				target: 8,
			},
			want: [][]int{{2, 2, 2, 2}},
		},
		{
			name: "zeroes",
			args: args{
				nums:   []int{0, 0, 0, 0},
				target: 0,
			},
			want: [][]int{{0, 0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fourSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_threeSumTargetSorted(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name         string
		args         args
		wantTriplets [][]int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			wantTriplets: [][]int{{-2, 0, 2}, {-1, 0, 1}},
		},
		{
			name: "",
			args: args{
				nums:   []int{2, 2, 2, 2},
				target: 6,
			},
			wantTriplets: [][]int{{2, 2, 2}},
		},
		{
			name: "zeroes",
			args: args{
				nums:   []int{0, 0, 0},
				target: 0,
			},
			wantTriplets: [][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Ints(tt.args.nums) // the 4Sum function sorts the input
			if gotTriplets := threeSumTargetSorted(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotTriplets, tt.wantTriplets) {
				t.Errorf("threeSumTargetSorted() = %v, want %v", gotTriplets, tt.wantTriplets)
			}
		})
	}
}

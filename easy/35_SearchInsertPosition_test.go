package easy

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{[]int{1, 3, 5, 6}, 5}, 2},
		{"Case 2", args{[]int{1, 3, 5, 6}, 2}, 1},
		{"Case 3", args{[]int{1, 3, 5, 6}, 7}, 4},
		{"Case 3", args{[]int{1, 3}, 2}, 1},
		{"Case 3", args{[]int{1, 3}, 1}, 0},
		{"Case 3", args{[]int{1, 3}, 3}, 1},
		{"Case3SearchInsert", args{[]int{1, 2, 3, 4, 5, 10}, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

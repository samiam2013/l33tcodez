package easy

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{2}, 2},
		{"Case 2", args{3}, 3},
		{"Case 3", args{4}, 5},
		{"Case 4", args{5}, 8},
		{"Case 5", args{6}, 13},
		{"Case 6", args{7}, 21},
		{"Case 7", args{8}, 34},
		{"Case 8", args{9}, 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

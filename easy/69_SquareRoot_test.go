package easy

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{4}, 2},
		{"Case 2", args{8}, 2},
		{"Case 3", args{9}, 3},
		{"Case 4", args{10}, 3},
		{"Case 5", args{16}, 4},
		{"Case 6", args{25}, 5},
		{"Case 7", args{36}, 6},
		{"Case 8", args{123456789 * 123456788}, 123456788},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}

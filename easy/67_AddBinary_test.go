package easy

import "testing"

func Test_addOne(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Case 1", args{"11", "1"}, "100"},
		{"Case 2", args{"1010", "1011"}, "10101"},
		{"Case 3", args{"1", "111"}, "1000"},
		{"Case 4", args{"0", "0"}, "0"},
		{"Case 5", args{"0", "1"}, "1"},
		{"Case 6", args{"1", "0"}, "1"},
		{"Case 7", args{"1", "1"}, "10"},
		{"Case 8", args{"111", "111"}, "1110"},
		{"Case 9", args{"111", "1111"}, "10110"},
		{"Case 10", args{"1111", "1111"}, "11110"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("addOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

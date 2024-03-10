package easy

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{"III"}, 3},
		{"Case 2", args{"IV"}, 4},
		{"Case 3", args{"IX"}, 9},
		{"Case 4", args{"LVIII"}, 58},
		{"Case 5", args{"MCMXCIV"}, 1994},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

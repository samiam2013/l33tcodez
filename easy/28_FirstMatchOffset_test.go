package easy

import "testing"

func Test_strStr(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"hello", "ll"}, 2},
		{"", args{"sadbutsad", "sad"}, 0},
		{"", args{"leetcode", "leeto"}, -1},
		{"", args{"aaa", "aaa"}, 0},
		{"", args{"aaaaa", "bba"}, -1},
		{"", args{"mississippi", "issi"}, 1},
		{"", args{"mississippi", "sippj"}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

package medium

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "Example 3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "Found Edge Case 1",
			args: args{
				s: "aab",
			},
			want: 2,
		},
		{
			name: "Found Edge Case 2",
			args: args{
				s: "aabaab!bb",
			},
			want: 3,
		},
		{
			name: "Found Edge Case 3",
			args: args{
				s: "dvdf",
			},
			want: 3,
		},
		{
			name: "Found Edge Case 5",
			args: args{
				s: "bpfbhmipx",
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

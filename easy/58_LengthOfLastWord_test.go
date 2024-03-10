package easy

import "testing"

func Test_lengthOfLastWord(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Case 1", args{"Hello World"}, 5},
		{"Case 3", args{"a "}, 1},
		{"Case 4", args{"a"}, 1},
		{"Case 5", args{"a b"}, 1},
		{"Case 6", args{"a b c"}, 1},
		{"Case 7", args{"a b c "}, 1},
		{"Case 8", args{" a"}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLastWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

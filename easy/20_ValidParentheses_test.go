package easy

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Case 1", args{"()"}, true},
		{"Case 2", args{"()[]{}"}, true},
		{"Case 3", args{"(]"}, false},
		{"Case 4", args{"([)]"}, false},
		{"Case 5", args{"{[]}"}, true},
		{"Case 6", args{"["}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

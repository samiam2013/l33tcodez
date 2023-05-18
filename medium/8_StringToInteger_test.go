package medium

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "basic usage",
			args: args{
				s: "123",
			},
			want: 123,
		},
		{
			name: "negative",
			args: args{
				s: "  -123",
			},
			want: -123,
		},
		{
			name: "positive with words",
			args: args{
				s: "  +123words",
			},
			want: 123,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}

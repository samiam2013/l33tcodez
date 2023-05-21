package hard

import (
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
				2,
			},
			want: []int{2, 1, 4, 3, 5},
		},
		{
			name: "Example 2",
			args: args{
				&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}},
				3,
			},
			want: []int{3, 2, 1, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseKGroup(tt.args.head, tt.args.k)
			gotSlice := listToSlice(got)
			if !slicesMatch(gotSlice, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}

func listToSlice(head *ListNode) []int {
	var slice []int
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	return slice
}

func slicesMatch(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

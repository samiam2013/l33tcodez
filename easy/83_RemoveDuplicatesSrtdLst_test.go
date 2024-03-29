package easy

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"Case 1", args{&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}}, &ListNode{Val: 1, Next: &ListNode{Val: 2}}},
		{"Case 2", args{&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 3}}}}}}, &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}},
		{"Case 3", args{&ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1}}}}, &ListNode{Val: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

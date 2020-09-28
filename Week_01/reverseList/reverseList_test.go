package reverseList

import "testing"


func TestReverseList(t *testing.T) {
	arr := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					4,
					&ListNode{},
				},
			},
		},
	}
	info := reverseList(arr)
	t.Log(*info)
}

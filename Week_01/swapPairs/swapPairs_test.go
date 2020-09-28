package swapPairs

import "testing"

func TestSwapPairs(t *testing.T)  {
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
	info := swapPairs(arr)
	t.Log(*info)
}
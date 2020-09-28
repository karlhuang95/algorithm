package hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return false
	}
	slow := head
	runner := head
	for runner.Next != nil && runner.Next.Next != nil {
		slow = slow.Next
		runner = runner.Next.Next
		if slow == runner {
			return true
		}
	}
	return false
}

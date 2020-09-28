package reverseList

type ListNode struct {
	Val  int
	Next *ListNode
}
// 方法一： 迭代法
func reverseList(head *ListNode) *ListNode {
	var newHead,next *ListNode
	for head != nil {
		next = head.Next
		head.Next = newHead
		newHead = head
		head = next
	}
	return newHead
}

// 方法二：递归
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
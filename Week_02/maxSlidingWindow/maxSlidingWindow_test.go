package maxSlidingWindow

import "testing"

func TestMaxSlidingWindow(t *testing.T)  {
	arr := []int{1,3,-1,-3,5,3,6,7}
	info := maxSlidingWindow1(arr,3)
	t.Log(info)
}

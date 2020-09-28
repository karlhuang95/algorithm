package maxArea

import "testing"

func TestMaxArea(t *testing.T) {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}

	info := maxArea(arr)
	t.Log(info)
}

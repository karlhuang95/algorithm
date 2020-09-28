package threeSum

import "testing"

func TestThreeSum(t *testing.T) {
	arr := []int{
		-1, 0, 1, 2, -1, -4,
	}
	info := threeSum(arr)
	t.Log(info)
}


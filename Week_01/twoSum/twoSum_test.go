package twoSum

import "testing"

func TestTowSum(t *testing.T) {
	arr := []int{2, 7, 11, 15}
	info := twoSum(arr, 9)
	t.Log(info)
}

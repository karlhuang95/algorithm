package maxSlidingWindow

//
func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil {
		return nil
	}
	var q = make([]int, 0, len(nums))
	res := make([]int, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		for len(q) != 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		if q[0] == i-k {
			q = q[1:]
		}
		if i+1-k >= 0 {
			res[i+1-k] = nums[q[0]]
		}
	}
	return res
}





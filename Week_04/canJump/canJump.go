package canJump

func canJump(nums []int) bool {
	l := len(nums) - 1
	for i := l - 1; i >= 0; i-- {
		if nums[i]+i >= l {
			l = i
		}
	}
	return l <= 0

}

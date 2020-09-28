package jump

func max(x, y int) int {
	if x < y {
		return y
	} else {
		return x
	}
}
func jump(nums []int) int {
	var steps = 0
	var end = 0
	var maxPos = 0
	for i := 0; i < len(nums)-1; i++ {
		maxPos = max(maxPos, nums[i]+i)
		if i == end {
			end = maxPos
			steps++
		}
	}
	return steps
}

package maxArea

func maxArea(height []int) int {
	i := 0
	max := 0
	j := len(height) - 1
	for ; i < j; {
		long := j - i
		high := height[i]
		if height[j] < height[i] {
			high = height[j]
			j--
		} else {
			high = height[i]
			i++
		}
		if long*high > max {
			max = long * high
		}
	}
	return max
}

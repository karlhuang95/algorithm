package maxCoins

func maxCoins(nums []int) int {
	// Store the input's length.
	n := len(nums)
	// Padding 1s to head and tail of nums.
	nums = append(nums, 1)
	nums = append([]int{1}, nums...)

	// Create 2D-DP with n+2 width and height.
	// c[i][j] represents the max coins from i to j.
	c := make([][]int, n+2)
	for i := range c {
		c[i] = make([]int, n+2)
	}

	// l is the length of subarray. We start with l= 1, end with l = n.
	for l := 1; l <= n; l++ {
		// i is the start point in this subarray.
		for i := 1; i <= n-l+1; i++ {
			// j is the subarray's end.
			j := i + l - 1
			// k is the break point to separate.
			for k := i; k <= j; k++ {
				c[i][j] = max(c[i][j], c[i][k-1]+nums[i-1]*nums[k]*nums[j+1]+c[k+1][j])
			}
		}
	}

	return c[1][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


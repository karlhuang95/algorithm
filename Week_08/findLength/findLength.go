package findLength

// 滑动窗口
func findLength(A []int, B []int) int {
	var ans int

	n, m := len(A), len(B)
	for i := 0; i < n; i++ {
		len := min(m, n-i)
		maxLenth := maxLen(A, B, i, 0, len)
		ans = max(ans, maxLenth)
	}

	for i := 0; i < m; i++ {
		len := min(n, m-i)
		maxLenth := maxLen(A, B, 0, i, len)
		ans = max(ans, maxLenth)
	}

	return ans


}

func maxLen(A []int, B []int, a int, b int, len int) int {
	ret, k := 0, 0
	for i := 0; i < len; i++ {
		if A[a+i] == B[b+i] {
			k++
		} else {
			k = 0
		}
		ret = max(ret, k)
	}
	return ret
}

func max(a int, b int) int {
	switch {
	case a > b:
		return a
	}
	return b
}

func min(a int, b int) int {
	switch {
	case a < b:
		return a
	}
	return b
}

// 动态规划
/**
func findLength(A []int, B []int) int {
    var ans int
    n, m := len(A), len(B)
    dp := make([][]int, n+1)
    for i := 0; i < len(dp); i++ { //初始化二维数组
        dp[i] = make([]int, m+1)
    }
    for i := n-1; i >= 0; i-- {
        for j := m-1; j >= 0; j-- {
            if A[i] == B[j] {
                dp[i][j] = dp[i+1][j+1] + 1  //从最后一个开始比较
            } else {
                dp[i][j] = 0
            }
            ans = max(ans, dp[i][j])  //取最大值
        }
    }
    return ans
}

func max(a int, b int) int {
    switch {
        case a > b:
            return a
    }
    return b
}


 */
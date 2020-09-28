package countSubstrings

func countSubstrings(s string) int {
	ret := 0
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for gap := 0; gap < len(s); gap++ {
		for i := 0; i+gap < len(s); i++ {
			j := i + gap  // dp[i][j] 间隔gap 即以子序列长度为标准遍历
			if gap == 0 { // 初始化base
				dp[i][j] = true
			} else if gap == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			if dp[i][j] {
				ret++
			}
		}
	}
	return ret
}

package coinChange
//动态规划（不压缩空间）
func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	availableCoins := make([]int, 0)
	for _, coin := range coins {
		if coin > amount {
			continue
		}
		availableCoins = append(availableCoins, coin)
		dp[coin] = 1
	}
	for i := 1; i < len(dp); i++ {
		if dp[i] != 0 {
			continue
		}
		dp[i] = -1
		for _, coin := range coins {
			lastIdx := i-coin
			if lastIdx <= 0 || dp[lastIdx] == -1{
				continue
			}
			if dp[i] == -1 || (dp[i] != -1 && dp[lastIdx]+1 < dp[i]) {
				dp[i] = dp[lastIdx]+1
			}
		}
	}
	return dp[amount]

}


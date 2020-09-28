package canCross

func canCross(stones []int) bool {
	// dp[stones[i]] 用个map保存第个位置i所有能跳的k值
	dp := make(map[int]map[int]int, len(stones))
	for _, i := range stones {
		dp[i] = make(map[int]int)
	}
	dp[0][0] = 0
	for i := 0; i < len(stones); i++ {
		for _, k := range dp[stones[i]] {
			for step := k - 1; step <= k + 1; step++ {
				//  跳的距离大于0 且 跳到数组中存在的位置
				//  如[0,1,3,5,6,8,12,17] 1 点能跳 0，1， 2 共3种距离 只有2跳到了 3里面
				if step > 0 {
					if _, ok := dp[stones[i] + step]; ok {
						dp[stones[i] + step][i] = step
					}
				}
			}
		}
	}

	return  len(dp[stones[len(stones)-1]]) != 0
}


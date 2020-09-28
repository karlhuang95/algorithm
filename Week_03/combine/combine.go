package combine

func combine(n int, k int) [][]int {

	ans := make([][]int, 0)
	option := make([]int, k)
	var dfs func(start int, idx int)
	dfs = func(start int, idx int) {
		for i := start; i <= n-(k-1-idx); i++ {
			option[idx] = i
			if idx == k-1 {
				ans = append(ans, append([]int{}, option...))
			} else {
				dfs(i+1, idx+1)
			}
		}
	}
	dfs(1, 0)
	return ans

}

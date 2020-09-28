package findCircleNum


func findCircleNum(M [][]int) int {
	n := len(M)
	p := make([]*[]int, n)
	for i := 0; i < n; i++ {
		p[i] = &[]int{i}
	}
	ans := n
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			i, j := i, j
			if M[i][j] == 1 && p[i] != p[j] {
				if len(*p[i]) < len(*p[j]) {
					i, j = j, i
				}
				*p[i] = append(*p[i], *p[j]...)
				for _, k := range *p[j] {
					p[k] = p[i]
				}
				ans--
			}
		}
	}
	return ans
}

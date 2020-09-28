package minMutation
// 深度优先
var mutationMap = map[uint8][3]string{
	'A': [...]string{"T", "G", "C"},
	'C': [...]string{"T", "G", "A"},
	'T': [...]string{"A", "G", "C"},
	'G': [...]string{"T", "A", "C"},
}

func idxOf(str string, bank []string) int {
	for i, s := range bank {
		if s == str {
			return i
		}
	}
	return -1
}

func minMutation(start string, end string, bank []string) int {
	if idxOf(end, bank) == -1 {
		return -1
	}
	minCount := len(bank) + 1
	isUsed := make([]bool, len(bank))

	var dfs func(curr string, count int)
	dfs = func(curr string, count int) {
		if count >= minCount {
			return
		}
		if curr == end {
			if count < minCount {
				minCount = count
			}
			return
		}

		for i := 0; i < len(curr); i++ {
			for _, c := range mutationMap[curr[i]] {
				gene := curr[:i] + c + curr[i+1:]
				if idx := idxOf(gene, bank); idx != -1 && !isUsed[idx] {
					isUsed[idx] = true
					dfs(gene, count+1)
					isUsed[idx] = false
				}
			}
		}
	}
	dfs(start, 0)
	if minCount <= len(bank) {
		return minCount
	}
	return -1
}


// 广度优先
func minMutation1(start string, end string, bank []string) int {
	if idxOf(end, bank) == -1 {
		return -1
	}
	isUsed := make([]bool, len(bank))

	queue := []string{start}
	count := 0
	for len(queue) > 0 {
		l := len(queue)
		for i:= 0; i < l; i++ {
			curr := queue[i]
			if curr == end {
				return count
			}
			for j:= 0;j< len(curr);j++ {
				for _, s := range mutationMap[curr[j]] {
					if idx := idxOf(curr[:j] + s + curr[j+1:], bank); idx != -1 && !isUsed[idx] {
						queue = append(queue, bank[idx])
						isUsed[idx] = true
					}
				}

			}

		}
		count++
		queue = queue[l:]
	}
	return -1
}

// 双向广度优先
func minMutation2(start string, end string, bank []string) int {
	if idxOf(end, bank) == -1 {
		return -1
	}
	if start == end {
		return 0
	}
	count := 0
	isUsed := make([]bool, len(bank))

	startQueue := []string{start}
	endQueue := []string{end}
	for len(startQueue) > 0 {
		count++
		l := len(startQueue)
		for _, curr := range startQueue {
			for i := 0; i < len(curr); i++ {
				for _, c := range mutationMap[curr[i]] {
					gene := curr[:i] + c + curr[i+1:]
					if idx := idxOf(gene, endQueue); idx != -1 {
						return count
					}
					if idx := idxOf(gene, bank); idx != -1 && !isUsed[idx] {
						isUsed[idx] = true
						startQueue = append(startQueue, gene)
					}
				}
			}
		}
		startQueue = startQueue[l:]
		if len(startQueue) > len(endQueue) {
			startQueue, endQueue = endQueue, startQueue
		}
	}
	return -1
}



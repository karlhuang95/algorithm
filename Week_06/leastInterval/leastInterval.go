package leastInterval

func Count(b []byte) (bool, int, []byte) {
	if len(b) == 0 {
		return false, 0, nil
	}
	/*定义一个字典统计每个字符的出现次数*/
	m := make(map[byte]int)
	for _, v := range b {
		/*若字典中没有v这个键，则m[v]的值默认为0*/
		m[v]++
	}
	/*把所有出现次数存储到一个切片内*/
	s := make([]int, 0)
	for _, v := range m {
		s = append(s, v)
	}
	/*找到所有出现次数中最大的那一个*/
	maxAccount := s[0]
	for _, v := range s[1:] {
		if v > maxAccount {
			maxAccount = v
		}
	}
	/*遍历字典中的值，如果和maxAccount相等，就把它存进maxElements*/
	maxElements := make([]byte, 0)
	for k, v := range m {
		if v == maxAccount {
			maxElements = append(maxElements, k)
		}
	}
	return true, maxAccount, maxElements
}

func leastInterval(tasks []byte, n int) int {
	flag, max, ele := Count(tasks)
	var point int
	if flag {
		point = (n+1)*max - n + (len(ele) - 1)
		if len(tasks) < point {
			return point
		} else {
			return len(tasks)
		}
	} else {
		panic("字符切片为空，请检查测试用例")
	}
}


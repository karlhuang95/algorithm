package findAnagrams
func findAnagrams(s string, p string) []int {

	need, window := map[byte]int{}, map[byte]int{}
	for i := range p {
		need[p[i]]++
	}
	left, right := 0, 0
	valid := 0
	indexs := []int{}
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left >= len(p) {
			if valid == len(need) {
				indexs = append(indexs, left)
			}
			d := s[left]
			left++
			if _, ok := window[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return indexs
}


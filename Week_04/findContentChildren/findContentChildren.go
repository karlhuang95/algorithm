package findContentChildren

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i, j := 0, 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i

}
func findContentChildren1(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var count int

	i, j := 0, 0

	for {
		if i >= len(g) {
			break
		}
		if j >= len(s) {
			break
		}

		if s[j] >= g[i] {
			count++
			j++
			i++
		} else {
			j++
		}
	}

	return count

}
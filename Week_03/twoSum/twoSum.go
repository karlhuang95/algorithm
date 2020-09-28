package twoSum

func twoSum(nums []int, target int) []int {
	result := []int{}
	m := make(map[int]int)
	for i, k := range nums {
		if v, n := m[target-k]; n {
			result = append(result, v)
			result = append(result, i)
		}
		m[k] = i
	}
	return result
}

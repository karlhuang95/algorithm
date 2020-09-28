package relativeSortArray

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	hash := make(map[int]int, 0)
	res := make([]int, 0)

	if len(arr2) == 0 {
		sort.Ints(arr1)
		return arr1
	}

	for _, v := range arr1 {
		hash[v]++
	}
	for i := 0; i < len(arr2); i++ {
		for j := 0; j < hash[arr2[i]]; j++ {
			res = append(res, arr2[i])
		}
		hash[arr2[i]] = 0
	}

	temp := make([]int, 0)
	for k, v := range hash {
		for v > 0 {
			temp = append(temp, k)
			v--
		}
	}
	sort.Ints(temp)

	res = append(res, temp...)
	return res

}
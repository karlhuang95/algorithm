package intersect

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	result := make(map[int]int)
	res := make([]int, 0)
	for _,v := range nums1 {
		fmt.Println(result)
		result[v]++
	}
	for _,v := range nums2 {
		if result[v]> 0 {
			res = append(res,v)
			result[v]--
		}
	}
	return res
}

package reversePairs

import "sort"
// 注意: 这不是最好的版本
func reversePairs(nums []int) int {
	// 树状数组：BIT
	cache := make(map[int]struct{}) // 1
	for _, v := range nums {        // 1.1
		cache[v] = struct{}{}
		cache[v<<1] = struct{}{}
	}
	n := len(cache)
	a, c, count, i := make([]int, n), make([]int, n+1), 0, 0 // 2.1
	for k, _ := range cache {                                // 1.2
		a[i], i = k, i+1
	}
	sort.Ints(a)
	for i := len(nums) - 1; i >= 0; i-- { // 3
		idx1, idx2 := binarySearch(a, 0, n, nums[i]), binarySearch(a, 0, n, nums[i]<<1) // 3.1
		count += query(c, idx1)                                                         // 3.2
		update(c, idx2+1)                                                               // 2 & 3.3
	}
	return count
}
func binarySearch(arr []int, l, r int, val int) int {
	for l < r {
		mid := (l + r) >> 1
		if arr[mid] < val {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
func update(bit []int, i int) { // 2.3
	for i < len(bit) {
		bit[i] ++
		i += i & -i
	}
}
func query(bit []int, i int) int { // 2.2
	count := 0
	for i > 0 {
		count += bit[i]
		i -= i & -i
	}
	return count
}

package permuteUnique

func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	var recursivePermute func(int)
	recursivePermute = func(pos int) {
		if pos == len(nums)-1 {
			tmp := make([]int, len(nums))
			copy(tmp, nums)
			result = append(result, tmp)
			return
		}

		seen := map[int]bool{} // 记录出现过的数字，出现过就不做交换了
		for i := pos; i < len(nums); i++ {
			if !seen[nums[i]] {
				nums[i], nums[pos] = nums[pos], nums[i]
				recursivePermute(pos + 1)
				nums[i], nums[pos] = nums[pos], nums[i]
				seen[nums[i]] = true
			}
		}
	}

	recursivePermute(0)
	return result

}

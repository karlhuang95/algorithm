package moveZeroes

func moveZeroes(nums []int) []int {
	l := len(nums)
	i := 0
	for {
		if i >= l {
			break
		}
		if nums[i] == 0 {
			l -= 1
			nums = append(nums[0:i], nums[i+1:]...)
			nums = append(nums, 0)
		} else {
			i += 1
		}
	}
	return nums

}

// 方法二
func moveZeroes1(nums []int) []int {
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if n != i {
				nums[n], nums[i] = nums[i], 0
			}
			n++
		}
	}
	return nums
}

func moveZeroes2(nums []int) {
	sonwBallSize := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sonwBallSize++
		} else if sonwBallSize > 0 {
			t := nums[i]
			nums[i] = 0
			nums[i-sonwBallSize] = t
		}
	}
}

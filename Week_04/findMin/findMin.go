package findMin

func findMin(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = right - 1
		}
	}
	return nums[left]

}

func findMin1(nums []int) int {
	var min = nums[0]
	for _, i := range nums {
		if i < min {
			return i
		}
	}
	return min

}

package splitArray

func splitArray(nums []int, m int) int {
	left, right, length := 0, 0, len(nums)
	for i := 0; i < length; i++ {
		right += nums[i]
		if nums[i] > left {
			left = nums[i]
		}
	}
	ans := right
	for left <= right {
		mid := left + (right-left)/2
		sum, cnt := 0, 1
		for i := 0; i < length; i++ {
			if sum+nums[i] > mid {
				cnt++
				sum = nums[i]
			} else {
				sum += nums[i]
			}
		}
		if cnt <= m {
			ans = min(mid, ans)
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
func min(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

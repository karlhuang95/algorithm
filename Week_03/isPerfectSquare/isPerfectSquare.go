package isPerfectSquare


func isPerfectSquare(num int) bool {
	if num < 2{
		return true
	}
	min , max := 0 , num/2 + 1
	for max - min >= 0{
		mid := (min + max) / 2
		if num > mid * mid{
			min = mid + 1
		}else if num < mid * mid{
			max = mid - 1
		}else if num == mid * mid{
			return true
		}
	}
	return false
}


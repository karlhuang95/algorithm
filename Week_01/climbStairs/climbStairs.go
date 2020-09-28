package climbStairs

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	one_setps_before := 2
	two_setps_before := 1
	all_ways := 0
	for i := 2; n > i; i++ {
		all_ways = one_setps_before + two_setps_before
		two_setps_before = one_setps_before
		one_setps_before = all_ways
	}
	return all_ways
}

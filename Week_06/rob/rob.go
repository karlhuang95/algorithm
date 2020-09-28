package rob
func Max(a int,b int)int{
	if a>b{
		return a
	}else {
		return b
	}
}

func rob(nums []int) int {
	n:=len(nums)
	if n==0{
		return 0
	}else if n==1{
		return nums[0]
	}
	dp1:=make([]int,n)
	dp2:=make([]int,n)
	dp1[0]=nums[0]
	dp2[0]=0
	dp1[1]=Max(nums[0],nums[1])
	dp2[1]=nums[1]
	for i:=2;i<n;i++{
		dp1[i]=Max(dp1[i-1],dp1[i-2]+nums[i])
		dp2[i]=Max(dp2[i-1],dp2[i-2]+nums[i])
	}
	return Max(dp1[n-2],dp2[n-1])
}


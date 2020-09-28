package minPathSum

func minPathSum(grid [][]int) int {

	m:=len(grid)
	n:=len(grid[0])
	var dp=make([][]int,m)
	//dp[0][0]=grid[0][0]
	for i:=0;i<m;i++{
		dp[i]=make([]int,n)
		for j:=0;j<n;j++{

			if i==0  && j>0{//当行为0时，列不等于0
				dp[i][j]=dp[i][j-1]+grid[i][j]
			}else if j==0 && i>0{//当列等于0，行不等于0
				dp[i][j]=dp[i-1][j]+grid[i][j]
			}else if i==0 && j==0{//当列等于0，行等于0
				dp[i][j]=grid[i][j]
			}else{//求出i，j位置中的最短距离
				dp[i][j]=grid[i][j]+mix(dp[i-1][j],dp[i][j-1])
			}

		}

	}
	return dp[m-1][n-1]
}

func mix(a,b int) int{//求出最小值
	if a>b{
		return b
	}

	return a
}


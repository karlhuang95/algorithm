package solveNQueens
func solveNQueens(n int) [][]string {
	res := make([][][]int,0)
	tmpRes := make([][]int,0)
	for i:=0;i<n;i++{
		tmp := make([]int,n)
		tmpRes = append(tmpRes,tmp)
	}
	dfsQueens(n,tmpRes,&res,0,0)
	finalRes := [][]string{}
	for i:=0;i<len(res);i++{
		tmp2 := []string{}
		for j:=0;j<len(res[i]);j++{
			tmp := ""
			for k:=0;k<len(res[i][j]);k++{
				if res[i][j][k] == 0{
					tmp += "."
				}else if res[i][j][k] == 1{
					tmp += "Q"
				}
			}
			tmp2 = append(tmp2,tmp)
		}
		finalRes = append(finalRes,tmp2)
	}
	return finalRes
}

func dfsQueens(n int,tmpRes [][]int,res *[][][]int,now int,col int){
	if now == n{
		rr := make([][]int,0,0)
		for i:=0;i<n;i++{
			r := make([]int,n,n)
			copy(r,tmpRes[i])
			rr = append(rr,r)
		}
		*res = append(*res,rr)
		return
	}
	for i:=0;i<n;i++{
		//剪掉直线，相邻对角线
		if now != 0 && (col == i || col-1 == i || col+1 == i){
			continue
		}
		//剪掉直线
		flag := true
		for j:=0;j<now;j++{
			if tmpRes[j][i] == 1{
				flag = false
			}
		}
		if flag == false{
			continue
		}
		//剪对角线1
		flag = true
		tmp1 := now-1
		tmp2 := i-1
		for tmp1 >= 0 && tmp2 >= 0{
			if tmpRes[tmp1][tmp2] == 1{
				flag = false
				break
			}
			tmp1--
			tmp2--
		}
		if flag == false{
			continue
		}
		//剪对角线2
		flag = true
		tmp1 = now-1
		tmp2 = i+1
		for tmp1 >= 0 && tmp2 <= n-1{
			if tmpRes[tmp1][tmp2] == 1{
				flag = false
				break
			}
			tmp1--
			tmp2++
		}
		if flag == false{
			continue
		}
		tmpRes[now][i] = 1
		dfsQueens(n,tmpRes,res,now+1,i)
		tmpRes[now][i] = 0
	}
}


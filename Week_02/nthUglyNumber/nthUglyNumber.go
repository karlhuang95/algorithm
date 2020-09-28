package nthUglyNumber

func nthUglyNumber(n int) int {

	i,j,k,res:=0,0,0,make([]int,n)
	min:=func(x,y int)int{
		if x<y{
			return x
		}
		return y
	}

	res[0]=1
	for count:=1;count<n;count++{
		tmp := min(2*res[i],min(3*res[j],5*res[k]))
		res[count] = tmp
		if tmp==res[i]*2{
			i++
		}
		if tmp==res[j]*3{
			j++
		}
		if tmp==res[k]*5{
			k++
		}
	}
	return  res[n-1]
}

package robotSim

func robotSim(commands []int, obstacles [][]int) int {
	type point struct{
		x int
		y int
	}
	obsMap:=make(map[point]bool)
	for i:=0;i< len(obstacles);i++{
		tmp:=point{obstacles[i][0],obstacles[i][1]}
		obsMap[tmp]=true
	}
	dir:=0
	x:=0
	y:=0
	res:=0
	points:=[][]int{{0,1},{-1,0},{0,-1},{1,0}}
	for i:=0;i< len(commands);i++{
		if commands[i]==-1{
			dir=(dir+4-1)%4
		}else if commands[i]==-2{
			dir=(dir+1)%4
		}else{
			tmpX:=x
			tmpY:=y
			for j:=1;j<=commands[i];j++{
				if _,ok:=obsMap[point{tmpX+points[dir][0],tmpY+points[dir][1]}];!ok{
					tmpX,tmpY=tmpX+points[dir][0],tmpY+points[dir][1]
				}else{
					x,y=tmpX,tmpY
					break
				}
			}
			x,y=tmpX,tmpY
			if res<x*x+y*y{
				res=x*x+y*y
			}
		}
	}
	return res

}
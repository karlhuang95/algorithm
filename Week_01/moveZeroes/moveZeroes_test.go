package moveZeroes

import "testing"

func TestMoveXeroes(t *testing.T)  {
	arr := []int{0,1,0,3,12}
	info := moveZeroes(arr)
	t.Log(info)

}
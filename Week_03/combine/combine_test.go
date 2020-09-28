package combine

import "testing"

func TestCombine(t *testing.T)  {
	n := 4
	k := 2
	info := combine(n,k)
	t.Log(info)
}

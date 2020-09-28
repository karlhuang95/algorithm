package intersect

import "testing"

func TestIntersect(t *testing.T)  {
	nums1:= []int{1,2,2,1}
	nums2:= []int{2,2}
	info := intersect(nums1,nums2)
	t.Log(info)
}
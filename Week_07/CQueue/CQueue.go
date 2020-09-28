package CQueue

type CQueue struct {
	p []int
	q []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.p = append(this.p, value)
}
func popFrom(q *[]int) (v int) {
	v, *q = (*q)[len(*q)-1], (*q)[:len(*q)-1]
	return
}

func (this *CQueue) DeleteHead() int {
	if len(this.q) > 0 {
		return popFrom(&this.q)
	}
	for i := len(this.p) - 1; i >= 0; i-- {
		this.q = append(this.q, this.p[i])
	}
	this.p = []int{}
	if len(this.q) > 0 {
		return popFrom(&this.q)
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

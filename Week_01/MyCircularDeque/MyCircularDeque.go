package MyCircularDeque


type MyCircularDeque struct {
	data []int
	maxLen int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		maxLen: k,
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if len(this.data) == this.maxLen {
		return false
	}
	this.data = append(this.data, value)
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if len(this.data) == this.maxLen {
		return false
	}
	this.data = append([]int{value}, this.data...)
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[0:len(this.data)-1]
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[1:]
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[0]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.data) == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	return len(this.data) == this.maxLen
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */

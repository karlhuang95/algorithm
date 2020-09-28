# 学习笔记

# 用add first或add last这套新的API改写Deque的代码

```
Dequedeque = new LinkedList();
deque.addLast("a");
deque.addLast("b");
deque.addLast("c");
System.out.println(deque);

String str = deque.peekFirst();
System.out.println(str);
System.out.println(deque);

while(deque.size()>0){
    System.out.println(deque.removeFirst())
}
System.out.println(deque);
```
# 分析Queue和Priority Queue的源码
```go
// 链接地址 github.com/golang-collections/go-datastructures/queue

//  PriorityQueue
func NewPriorityQueue(hint int) *PriorityQueue
func (pq *PriorityQueue) Dispose()
func (pq *PriorityQueue) Disposed() bool
func (pq *PriorityQueue) Empty() bool
func (pq *PriorityQueue) Get(number int) ([]Item, error)
func (pq *PriorityQueue) Len() int
func (pq *PriorityQueue) Peek() Item
func (pq *PriorityQueue) Put(items ...Item) error

// Queue
func New(hint int64) *Queue
func (q *Queue) Dispose()
func (q *Queue) Disposed() bool
func (q *Queue) Empty() bool
func (q *Queue) Get(number int64) ([]interface{}, error)
func (q *Queue) Len() int64
func (q *Queue) Put(items ...interface{}) error
func (q *Queue) TakeUntil(checker func(item interface{}) bool) ([]interface{}, error)
```

# PriorityQueue
github地址：https://github.com/golang-collections/go-datastructures/blob/master/queue/priority_queue.go

```go
type PriorityQueue struct {}
func (pq *PriorityQueue) Get(number int) ([]Item, error) {}
func (pq *PriorityQueue) Put(items ...Item) error {}
func (pq *PriorityQueue) Peek() Item {}
func (pq *PriorityQueue) Empty() bool {}
func (pq *PriorityQueue) Len() int {}
func (pq *PriorityQueue) Disposed() bool {}
func (pq *PriorityQueue) Disposed() bool {}
```
我们主要看一下`Get`和`Put`这两个方法，这两个方法还是比较有意思的。
```go
func (pq *PriorityQueue) Put(items ...Item) error {
	if len(items) == 0 {
		return nil
	}

	pq.lock.Lock()
	if pq.disposed {
		pq.lock.Unlock()
		return disposedError
	}
    //  注意1
	for _, item := range items {
		pq.items.insert(item)
	}

	for {
		sema := pq.waiters.get()
		if sema == nil {
			break
		}

		sema.response.Add(1)
		sema.wg.Done()
		sema.response.Wait()
		if len(pq.items) == 0 {
			break
		}
	}

	pq.lock.Unlock()
	return nil
}
```
其他的地方就是通过协程去操作，比较有意思的是这里`注意1`,通过循环items进行插入。我们可以看一下插入的具体操作
```go
func (items *priorityItems) insert(item Item) {
	if len(*items) == 0 {
		*items = append(*items, item)
		return
	}

	equalFound := false
   //  搜索元素 i 是否在已经升序排好的切片 *items 中
	i := sort.Search(len(*items), func(i int) bool {
		result := (*items)[i].Compare(item)
		if result == 0 {
			equalFound = true
		}
		return result >= 0
	})

	if equalFound {
		return
	}

	if i == len(*items) {
		*items = append(*items, item)
		return
	}

	*items = append(*items, nil)
    // 这里是通过切片去处理进行copy
	copy((*items)[i+1:], (*items)[i:])
	(*items)[i] = item
}
```
GET方法
```go

func (pq *PriorityQueue) Get(number int) ([]Item, error) {
	if number < 1 {
		return nil, nil
	}

	pq.lock.Lock()

	if pq.disposed {
		pq.lock.Unlock()
		return nil, disposedError
	}

	var items []Item

	if len(pq.items) == 0 {
		sema := newSema()
		pq.waiters.put(sema)
		sema.wg.Add(1)
		pq.lock.Unlock()

		sema.wg.Wait()
		pq.disposeLock.Lock()
		if pq.disposed {
			pq.disposeLock.Unlock()
			return nil, disposedError
		}
		pq.disposeLock.Unlock()

		items = pq.items.get(number)
		sema.response.Done()
		return items, nil
	}
    // 注意2
	items = pq.items.get(number)
	pq.lock.Unlock()
	return items, nil
}
```
Get从队列中检索项目。如果队列为空，此调用将阻塞，直到将下一个项目添加到队列中为止。这个将尝试检索项目数。

我们直接看他的内部get方法
```go 
func (items *priorityItems) get(number int) []Item {
	returnItems := make([]Item, 0, number)
	index := 0
	for i := 0; i < number; i++ {
		if i >= len(*items) {
			break
		}

		returnItems = append(returnItems, (*items)[i])
		(*items)[i] = nil
		index++
	}

	*items = (*items)[index:]
	return returnItems
}
```
这里我们可以看到，他初始化一个长度为0，容量为number的数组，如果循环的number次大于了items的长度就会退出。继续，将items中的值添加到returnItems中。
将相应的位置变为nil.

# Queue
对于queue还是主要看put和get
```go
type Queue struct {
	waiters  waiters
	items    items
	lock     sync.Mutex
	disposed bool
}
```
```go
func (q *Queue) Put(items ...interface{}) error {
	if len(items) == 0 {
		return nil
	}

	q.lock.Lock()

	if q.disposed {
		q.lock.Unlock()
		return disposedError
	}

	q.items = append(q.items, items...)
	for {
		sema := q.waiters.get()
		if sema == nil {
			break
		}
		sema.response.Add(1)
		sema.wg.Done()
		sema.response.Wait()
		if len(q.items) == 0 {
			break
		}
	}

	q.lock.Unlock()
	return nil
}
```
这里的put方法相对与Priority Queue就比较简单多了，它就直接添加到Queue结构体中的items中。

get方法
```go 
func (q *Queue) Get(number int64) ([]interface{}, error) {
	if number < 1 {
		// thanks again go
		return []interface{}{}, nil
	}

	q.lock.Lock()

	if q.disposed {
		q.lock.Unlock()
		return nil, disposedError
	}

	var items []interface{}

	if len(q.items) == 0 {
		sema := newSema()
		q.waiters.put(sema)
		sema.wg.Add(1)
		q.lock.Unlock()

		sema.wg.Wait()
		// we are now inside the put's lock
		if q.disposed {
			return nil, disposedError
		}
		items = q.items.get(number)
		sema.response.Done()
		return items, nil
	}

	items = q.items.get(number)
	q.lock.Unlock()
	return items, nil
}
```
通上，我们还是来看他的get方法
```go 
func (items *items) get(number int64) []interface{} {
	returnItems := make([]interface{}, 0, number)
	index := int64(0)
	for i := int64(0); i < number; i++ {
		if i >= int64(len(*items)) {
			break
		}

		returnItems = append(returnItems, (*items)[i])
		(*items)[i] = nil
		index++
	}

	*items = (*items)[index:]
	return returnItems
}
```
与Priority Queue类似。

# 对列
是一种特殊的线性表，特殊之处在于它只允许在表的前端（front）进行删除操作，而在表的后端（rear）进行插入操作，和栈一样，队列是一种操作受限制的线性表。进行插入操作的端称为队尾，进行删除操作的端称为队头。
队列的数据元素又称为队列元素。在队列中插入一个队列元素称为入队，从队列中删除一个队列元素称为出队。因为队列只允许在一端插入，在另一端删除，所以只有最早进入队列的元素才能最先从队列中删除，故队列又称为先进先出（FIFO—first in first out）
# 优先队列
在优先队列中，元素被赋予优先级，当访问元素时，具有最高级优先级的元素先被访问。即优先队列具有最高级先出的行为特征。它可以说是队列和排序的完美结合体，不仅可以存储数据，还可以将这些数据按照我们设定的规则进行排序。


# 个人总结
开始在数组还好一点，但是到了链表代码上就力不从心了，没有办法只能死记硬背，用这样的方式去记住内容。
总体来说还是感觉现在的代码水平有明显提高，开始注意自己的代码复杂度问题。这一点还是比较感谢老师的。

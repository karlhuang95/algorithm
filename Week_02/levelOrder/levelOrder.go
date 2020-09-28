package levelOrder

type Node struct {
	Val      int
	Children []*Node
}

var res [][]int

func levelOrder(root *Node) [][]int {
	res = [][]int{}
	if root == nil {
		return res
	}

	queue := []*Node{root}
	var level int
	for 0 < len(queue) {
		counter := len(queue)
		res = append(res, []int{})
		for i := 0; i < counter; i++ {
			if queue[i] != nil {
				res[level] = append(res[level], queue[i].Val)
				for _, n := range queue[i].Children {
					queue = append(queue, n)
				}
			}
		}
		queue = queue[counter:]
		level++
	}
	return res
}

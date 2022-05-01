package trees

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int32
}

type dummyStack struct {
	values []*TreeNode
}

func (d *dummyStack) size() int {
	return len(d.values)
}

func (d *dummyStack) Push(value *TreeNode) {
	d.values = append(d.values, value)
}

func (d *dummyStack) Pop() *TreeNode {
	if d.size() == 0 {
		return nil
	}
	value := d.values[d.size()-1]
	d.values = d.values[0 : d.size()-1]
	return value
}

func (t *TreeNode) InOrder() {
	stack := dummyStack{[]*TreeNode{}}
	fmt.Println(stack.values)
	node := t
	for node != nil || stack.size() > 0 {
		for node != nil {
			stack.Push(node)
			node = node.Left
		}
		node = stack.Pop()
		fmt.Printf("%d ", node.Value)
		node = node.Right
	}
}

func (t *TreeNode) BFS() {
	queue := []*TreeNode{t}
	for len(queue) > 0 {
		tmpQueue := make([]*TreeNode, 0, len(queue)*2)
		for _, node := range queue {
			fmt.Printf("%d ", node.Value)
			if node.Left != nil {
				tmpQueue = append(tmpQueue, node.Left)
			}
			if node.Right != nil {
				tmpQueue = append(tmpQueue, node.Right)
			}
		}
		fmt.Print("\n")
		queue = tmpQueue
	}
}

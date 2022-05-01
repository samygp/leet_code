package trees

import "fmt"

func CreateMinBST(sortedArray []int32) *TreeNode {
	fmt.Println(sortedArray)
	t := createMinBST(sortedArray, 0, len(sortedArray))
	t.InOrder()
	fmt.Print("\n")
	t.BFS()
	return t
}

func createMinBST(array []int32, from, to int) *TreeNode {
	if from >= to {
		return nil
	}
	middle := (from + to) / 2
	node := &TreeNode{Value: array[middle]}
	node.Left = createMinBST(array, from, middle)
	node.Right = createMinBST(array, middle+1, to)
	return node
}

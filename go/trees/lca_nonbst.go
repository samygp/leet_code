package trees

import "fmt"

var testTree *TreeNode = &TreeNode{
	Value: 4,
	Left: &TreeNode{
		Value: 2,
		Left:  &TreeNode{Value: 1},
		Right: &TreeNode{Value: 3},
	},
	Right: &TreeNode{
		Value: 6,
		Left:  &TreeNode{Value: 5},
		Right: &TreeNode{
			Value: 7,
			Right: &TreeNode{Value: 8},
		},
	},
}

type path struct {
	vals []int32
}

func LCANonBST() int32 {
	patha, pathb := &path{}, &path{}
	path_to(testTree, int32(5), patha)
	path_to(testTree, int32(8), pathb)
	fmt.Println(patha.vals)
	fmt.Println(pathb.vals)
	if len(patha.vals) == 0 || len(pathb.vals) == 0 {
		return -1
	}
	i, j := len(patha.vals)-1, len(pathb.vals)-1
	fmt.Printf("\n%d - %d", i, j)
	lca := patha.vals[i]
	for i >= 0 && j >= 0 {
		a, b := patha.vals[i], pathb.vals[j]
		fmt.Printf("\n%d %d", a, b)
		if a != b {
			fmt.Printf("\nLCA: %d", lca)
			return lca
		}
		if i > 0 {
			lca = a
			i--
		}
		if j > 0 {
			lca = b
			j--
		}
	}
	return -1
}

func path_to(t *TreeNode, val int32, result *path) {
	if len(result.vals) > 0 {
		return
	}
	if t.Value == val {
		result.vals = append(result.vals, t.Value)
		return
	}

	if t.Left != nil {
		path_to(t.Left, val, result)
		if len(result.vals) > 0 {
			result.vals = append(result.vals, t.Value)
			return
		}
	}
	if t.Right != nil {
		path_to(t.Right, val, result)
		if len(result.vals) > 0 {
			result.vals = append(result.vals, t.Value)
			return
		}
	}

}

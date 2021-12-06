package M_102_Binary_Tree_Level_Order_Traversal

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var values []int
	res := [][]int{{root.Val}}
	stack := []*TreeNode{root, nil}
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]

		if node == nil {
			if len(values) > 0 {
				res = append(res, values)
			}
			values = []int{}

			if len(stack) > 0 {
				stack = append(stack, nil)
			}
		} else {
			if node.Left != nil {
				values = append(values, node.Left.Val)
				stack = append(stack, node.Left)
			}

			if node.Right != nil {
				values = append(values, node.Right.Val)
				stack = append(stack, node.Right)
			}
		}
	}
	return res
}

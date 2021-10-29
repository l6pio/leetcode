package utils

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTree(arr []int) *TreeNode {
	var root *TreeNode
	queue := make([]*TreeNode, 0)
	for i := 0; i < len(arr); {
		if len(queue) == 0 {
			root = &TreeNode{
				Val:   arr[i],
				Left:  nil,
				Right: nil,
			}
			queue = append(queue, root)
			i += 1
		} else {
			leftTree := &TreeNode{
				Val:   arr[i],
				Left:  nil,
				Right: nil,
			}

			rightTree := &TreeNode{
				Val:   arr[i+1],
				Left:  nil,
				Right: nil,
			}

			head := queue[0]
			head.Left = leftTree
			head.Right = rightTree

			queue = queue[1:]
			queue = append(queue, leftTree, rightTree)
			i += 2
		}
	}
	return root
}

func PrintTree(t *TreeNode) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, t)
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		if head.Left != nil {
			queue = append(queue, head.Left)
		}
		if head.Right != nil {
			queue = append(queue, head.Right)
		}
		fmt.Printf("%v ", head.Val)
	}
}

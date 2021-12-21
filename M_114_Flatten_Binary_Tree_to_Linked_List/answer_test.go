package M_114_Flatten_Binary_Tree_to_Linked_List

import (
	"fmt"
	"testing"
)

func Test_flatten(t *testing.T) {
	tree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	flatten(tree)
	PrintTree(tree)
}

func PrintTree(t *TreeNode) {
	queue := make([]*TreeNode, 0)
	queue = append(queue, t, nil)
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		if head == nil {
			println("")
			if len(queue) > 0 {
				queue = append(queue, nil)
			}
		} else {
			tag := ""
			if head.Left != nil {
				queue = append(queue, head.Left)
				tag += "<"
			}
			if head.Right != nil {
				queue = append(queue, head.Right)
				tag += ">"
			}
			fmt.Printf("%v%s   ", head.Val, tag)
		}
	}
}

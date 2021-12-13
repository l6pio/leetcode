package M_105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

import (
	"fmt"
	"testing"
)

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

func Test_buildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	tree := buildTree(preorder, inorder)
	PrintTree(tree)
}

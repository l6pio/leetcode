package M_114_Flatten_Binary_Tree_to_Linked_List

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Flatten(root *TreeNode, arr *[]*TreeNode) {
	if root == nil {
		return
	}

	*arr = append(*arr, root)
	Flatten(root.Left, arr)
	Flatten(root.Right, arr)
}

func flatten(root *TreeNode) {
	arr := make([]*TreeNode, 0)
	Flatten(root, &arr)

	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		arr[i].Left = nil
		arr[i].Right = arr[i+1]
	}
}

package M_105_Construct_Binary_Tree_from_Preorder_and_Inorder_Traversal

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func BuildTree(preorder []int, inorder []int, root *TreeNode) {
	root.Val = preorder[0]
	inorderLen := len(inorder)
	for i := 0; i < inorderLen; i++ {
		if inorder[i] == root.Val {
			leftInorder := inorder[:i]
			leftInorderLen := len(leftInorder)
			if leftInorderLen > 0 {
				leftTree := &TreeNode{}
				root.Left = leftTree
				BuildTree(preorder[1:leftInorderLen+1], leftInorder, leftTree)
			}

			rightInorder := inorder[i+1:]
			rightInorderLen := len(rightInorder)
			if rightInorderLen > 0 {
				rightTree := &TreeNode{}
				root.Right = rightTree
				BuildTree(preorder[leftInorderLen+1:], rightInorder, rightTree)
			}
		}
	}
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	root := &TreeNode{}
	BuildTree(preorder, inorder, root)
	return root
}

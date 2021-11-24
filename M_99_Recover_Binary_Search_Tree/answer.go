package M_99_Recover_Binary_Search_Tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrderTraversal(root *TreeNode, arr *[]*TreeNode) {
	if root.Left != nil {
		inOrderTraversal(root.Left, arr)
	}

	*arr = append(*arr, root)

	if root.Right != nil {
		inOrderTraversal(root.Right, arr)
	}
}

func recoverTree(root *TreeNode) {
	arr := make([]*TreeNode, 0)
	inOrderTraversal(root, &arr)

	x, y := -1, -1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i].Val > arr[i+1].Val {
			if x == -1 {
				x = i
			} else {
				y = i
			}
		}
	}

	if x > -1 && y > -1 {
		tmp := arr[x].Val
		arr[x].Val = arr[y+1].Val
		arr[y+1].Val = tmp
	} else if x > -1 {
		tmp := arr[x].Val
		arr[x].Val = arr[x+1].Val
		arr[x+1].Val = tmp
	}
}

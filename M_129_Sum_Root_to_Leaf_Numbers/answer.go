package M_129_Sum_Root_to_Leaf_Numbers

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	oldStack := []*TreeNode{root}
	oldNums := []int{root.Val}
	var newStack []*TreeNode
	var newNums []int
	loop := true
	for loop {
		loop = false
		for i := 0; i < len(oldStack); i++ {
			if oldStack[i].Left != nil {
				newStack = append(newStack, oldStack[i].Left)
				newNums = append(newNums, oldNums[i]*10+oldStack[i].Left.Val)
				loop = true
			}
			if oldStack[i].Right != nil {
				newStack = append(newStack, oldStack[i].Right)
				newNums = append(newNums, oldNums[i]*10+oldStack[i].Right.Val)
				loop = true
			}
			if oldStack[i].Left == nil && oldStack[i].Right == nil {
				newStack = append(newStack, oldStack[i])
				newNums = append(newNums, oldNums[i])
			}
		}
		oldStack = newStack
		oldNums = newNums
		newStack = []*TreeNode{}
		newNums = []int{}
	}
	sum := 0
	for i := 0; i < len(oldNums); i++ {
		sum += oldNums[i]
	}
	return sum
}

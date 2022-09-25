package leetcode

import "fmt" 


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func StartBinaryTreeProblem() {
	fourth := &TreeNode{4, nil, nil}
	third := &TreeNode{3, nil, fourth}
	second := &TreeNode{1, nil, nil}
	first := &TreeNode{2, second, third}
	fmt.Println(maxDepth(first))
}
 
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	} 
	return right + 1
}
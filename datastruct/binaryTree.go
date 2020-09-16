package datastruct

import "math"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func InitTree(source []int, index int) *TreeNode {
	length := len(source)
	if length <= index || source[index] == math.MaxInt64 {
		return nil
	}

	root := &TreeNode{
		Val:source[index],
		Left: nil,
		Right: nil,
	}

	left := 2 * index + 1
	if left  < length {
		root.Left = InitTree(source, left)
	}

	right := 2 * index + 2
	if right < length {
		root.Right = InitTree(source, right)
	}
	return root
}

func PreOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := make([]int, 0)
	ret = append(ret, root.Val)
	if root.Left != nil {
		ret = append(ret, PreOrderTraversal(root.Left)...)
	}

	if root.Right != nil {
		ret = append(ret, PreOrderTraversal(root.Right)...)
	}
	return ret
}

func InOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := make([]int, 0)

	if root.Left != nil {
		ret = append(ret, PreOrderTraversal(root.Left)...)
	}

	ret = append(ret, root.Val)

	if root.Right != nil {
		ret = append(ret, PreOrderTraversal(root.Right)...)
	}
	return ret
}

func PostOrderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	ret := make([]int, 0)

	if root.Left != nil {
		ret = append(ret, PreOrderTraversal(root.Left)...)
	}

	if root.Right != nil {
		ret = append(ret, PreOrderTraversal(root.Right)...)
	}

	ret = append(ret, root.Val)
	return ret
}


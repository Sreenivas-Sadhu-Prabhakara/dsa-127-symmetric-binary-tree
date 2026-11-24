package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// IsSymmetric reports whether the tree is a mirror of itself around its
// center: the left subtree must mirror the right.
func IsSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return mirror(root.Left, root.Right)
}

// mirror reports whether a and b are mirror images: equal roots, a's left
// mirroring b's right and vice versa.
func mirror(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && mirror(a.Left, b.Right) && mirror(a.Right, b.Left)
}

func tn(v int, l, r *TreeNode) *TreeNode { return &TreeNode{Val: v, Left: l, Right: r} }

func main() {
	sym := tn(1, tn(2, tn(3, nil, nil), tn(4, nil, nil)), tn(2, tn(4, nil, nil), tn(3, nil, nil)))
	fmt.Println(IsSymmetric(sym))
}

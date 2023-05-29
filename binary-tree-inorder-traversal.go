package algorithm

import "fmt"

// 地址：https://leetcode.cn/problems/binary-tree-inorder-traversal/
// 94. 二叉树的中序遍历
// 给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
// 示例 1：
// 输入：root = [1,null,2,3]
// 输出：[1,3,2]
// 示例 2：
// 输入：root = []
// 输出：[]
// 示例 3：
// 输入：root = [1]
// 输出：[1]
// 提示：
// 树中节点数目在范围 [0, 100] 内
//-100 <= Node.val <= 100

func main() {
	tree := TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	fmt.Println(inorderTraversal(&tree))
}

// inorderTraversal 采用迭代+栈
func inorderTraversal(root *TreeNode) []int {
	re := make([]int, 0)
	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		re = append(re, root.Val)
		stack = stack[:len(stack)-1]
		root = root.Right
	}
	return re
}

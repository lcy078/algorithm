package algorithm

import "fmt"

// 地址：https://leetcode.cn/problems/balanced-binary-tree/
// 110. 平衡二叉树
// 给定一个二叉树，判断它是否是高度平衡的二叉树。
// 本题中，一棵高度平衡二叉树定义为：
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。
// 示例 1：
// 输入：root = [3,9,20,null,null,15,7]
// 输出：true

func main() {
	tree := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   9,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val:   15,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   7,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(isBalanced(&tree))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 获取左子树的最大深度和右子树的最大深度之差
	result := getDeep(root.Left) - getDeep(root.Right)
	if result < 0 {
		result = -result
	}
	if result > 1 {
		// 深度之差大于1，就是不平衡的，直接返回
		return false
	} else {
		// 递归比较左子树的左子树，右子树的右子树，都为true则是平衡的
		return isBalanced(root.Left) && isBalanced(root.Right)
	}
}

// 获取树的深度
func getDeep(node *TreeNode) int {
	if node == nil {
		return 0
	}
	// 递归获取深度
	leftDeep := getDeep(node.Left)
	rightDeep := getDeep(node.Right)
	// 比较左树、右树深度，获取最大深度+1
	if leftDeep > rightDeep {
		return leftDeep + 1
	}
	return rightDeep + 1
}

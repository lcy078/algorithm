package algorithm

import "fmt"

//剑指 Offer 54. 二叉搜索树的第k大节点
//给定一棵二叉搜索树，请找出其中第 k 大的节点的值。
//示例 1:
//输入: root = [3,1,4,null,2], k = 1
//   3
//  / \
// 1   4
//  \
//   2
// 输出: 4
// 示例 2:
// 输入: root = [5,3,6,2,4,null,null,1], k = 3
//     5
//    / \
//   3   6
//  / \
// 2   4
// /
// 1
// 输出: 4
// 限制：
// 1 ≤ k ≤ 二叉搜索树元素个数

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  1,
			Left: nil,
			Right: &TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(kthLargest(root, 1))
}

func kthLargest(root *TreeNode, k int) int {
	s := make([]int, 0, 0)
	// 进行中序遍历
	var recur func(*TreeNode)
	recur = func(node *TreeNode) {
		if node == nil {
			return
		}
		recur(node.Left)
		s = append(s, node.Val)
		recur(node.Right)
	}
	recur(root)
	// 取出相应的值
	return s[len(s)-k]
}

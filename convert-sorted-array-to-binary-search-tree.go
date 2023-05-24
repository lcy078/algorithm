package algorithm

import "fmt"

// 地址：https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/
// 108. 将有序数组转换为二叉搜索树
// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
// 高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。
// 示例 1：
// 输入：nums = [-10,-3,0,5,9]
// 输出：[0,-3,9,-10,null,5]
// 解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：
// 示例 2：
// 输入：nums = [1,3]
// 输出：[3,1]
// 解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。
// 提示：
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums 按 严格递增 顺序排列

func main() {
	nums := []int{1, 3, 4, 6, 8, 9, 11, 20, 44}
	fmt.Println(sortedArrayToBST(nums))
}

func sortedArrayToBST(nums []int) *TreeNode {
	// 已知nums严格升序，可以根据中位数递归赋值
	var recur func(left, right int) *TreeNode
	recur = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		node := &TreeNode{Val: nums[mid]}
		node.Left = recur(left, mid-1)
		node.Right = recur(mid+1, right)
		return node
	}
	return recur(0, len(nums)-1)
}

package algorithm

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 地址：https://leetcode.cn/problems/linked-list-cycle/
// 141. 环形链表
// 给你一个链表的头节点 head ，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
// 如果链表中存在环 ，则返回 true 。 否则，返回 false 。
// 示例 1：
// 输入：head = [3,2,0,-4], pos = 1
// 输出：true
// 解释：链表中有一个环，其尾部连接到第二个节点。
// 示例 2：
// 输入：head = [1,2], pos = 0
// 输出：true
// 解释：链表中有一个环，其尾部连接到第一个节点。
// 提示：
// 链表中节点的数目范围是 [0, 104]
// -105 <= Node.val <= 105
// pos 为 -1 或者链表中的一个 有效索引 。

func main() {
	nums := []int{3, 2, 0, -4}
	head := createCircularLinkedList(nums)
	fmt.Println(hasCycle(head))
}

// createCircularLinkedList 创建环形链表
func createCircularLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	// 创建第一个节点
	head := &ListNode{
		Val: nums[0],
	}

	// 保存头节点的指针
	current := head

	// 创建剩余节点，并将下一个节点指向头节点
	for i := 1; i < len(nums); i++ {
		node := &ListNode{
			Val: nums[i],
		}
		current.Next = node
		current = node
	}

	// 将最后一个节点的下一个节点指向头节点，形成环形结构
	current.Next = head

	return head
}

// hasCycle 判断是否是环形链表算法
func hasCycle(head *ListNode) bool {
	// 快慢指针
	if head == nil {
		return false
	}
	one := head
	two := head.Next
	if two == nil {
		return false
	}
	for {
		// 没有不能比较的值，可以比较大小
		if one == two {
			return true
		}
		if one.Next != nil && two.Next != nil && two.Next.Next != nil {
			one = one.Next
			two = two.Next.Next
		} else {
			return false
		}
	}
}

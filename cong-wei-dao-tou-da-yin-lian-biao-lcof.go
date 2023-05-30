package algorithm

import "fmt"

// 地址：https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
// 剑指 Offer 06. 从尾到头打印链表
// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。
// 示例 1：
// 输入：head = [1,3,2]
// 输出：[2,3,1]
// 限制：
// 0 <= 链表长度 <= 10000

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val:  2,
				Next: nil,
			},
		},
	}
	fmt.Println(reversePrint(head))
}

// reversePrint // 先遍历链表拿值，再进行原地交换
func reversePrint(head *ListNode) []int {
	re := make([]int, 0)
	for head != nil {
		re = append(re, head.Val)
		head = head.Next
	}
	l := len(re)
	for i := 0; i < l/2; i++ {
		re[i], re[l-i-1] = re[l-i-1], re[i]
	}
	return re
}

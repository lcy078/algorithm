package algorithm

import "fmt"

// 地址：https://leetcode.cn/problems/search-insert-position/
// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。
// 示例 1:
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2
// 输入: nums = [1,3,5,6], target = 2
// 输出: 1
// 示例 3:
// 输入: nums = [1,3,5,6], target = 7
// 输出: 4
// 提示:
// 1 <= nums.length <= 104
// -104 <= nums[i] <= 104
// nums 为无重复元素的升序排列数组
// -104 <= target <= 104

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))
}

// searchInsert 采用二分搜索
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	position := 0
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return target
		} else if nums[mid] > target {
			right = mid - 1
			position = mid
		} else {
			left = mid + 1
		}
	}
	return position
}

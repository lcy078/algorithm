package algorithm

import "fmt"

// 地址：https://leetcode.cn/problems/pascals-triangle/
// 118. 杨辉三角
// 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
// 示例 1:
// 输入: numRows = 5
// 输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
// 示例 2:
// 输入: numRows = 1
// 输出: [[1]]
// 提示:
// 1 <= numRows <= 30

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	s := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		s[i] = make([]int, i+1)
		s[i][0] = 1
		s[i][i] = 1
		for j := 1; j < i; j++ {
			s[i][j] = s[i-1][j] + s[i-1][j-1]
		}
	}
	return s
}

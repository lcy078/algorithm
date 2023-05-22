package algorithm

import "fmt"

// 地址：https://leetcode.cn/problems/longest-palindromic-substring/
// 5. 最长回文子串
// 给你一个字符串 s，找到 s 中最长的回文子串。
// 如果字符串的反序与原始字符串相同，则该字符串称为回文字符串。
// 示例 1：
// 输入：s = "babad"
// 输出："bab"
// 解释："aba" 同样是符合题意的答案。
// 示例 2：
// 输入：s = "cbbd"
// 输出："bb"
// 提示：
// 1 <= s.length <= 1000
// s 仅由数字和英文字母组成

func main() {
	s := "babaddcvbvcdd"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) string {
	l := len(s)
	if l <= 1 {
		return s
	}
	// 以第一个字符开始
	maxStr := s[0:1]
	for i := 0; i < l-1; i++ {
		// aba这种情况
		one := getStr(s, i, i)
		// abba这种情况
		double := getStr(s, i, i+1)
		maxStr = getMax(maxStr, one, double)
	}
	return maxStr
}

func getStr(s string, begin, end int) string {
	for {
		// 边界判断，同时如果是回文数的话，begin和end对应的值必须是相等的
		if begin <= end && begin >= 0 && end <= len(s)-1 && s[begin] == s[end] {
			begin--
			end++
		} else {
			break
		}
	}
	return s[begin+1 : end]
}

func getMax(str1, str2, str3 string) string {
	str1Len := len(str1)
	str2Len := len(str2)
	str3Len := len(str3)
	if str1Len > str2Len {
		if str1Len >= str3Len {
			return str1
		} else {
			return str3
		}
	} else {
		if str2Len >= str3Len {
			return str2
		} else {
			return str3
		}
	}
}

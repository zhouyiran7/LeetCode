/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */
package leetcode

import (
	"strings"
)

// @lc code=start
func IsPalindrome2(s string) bool {
	if len(s) == 0 {
		return true
	}
	// 过滤非数字和字母的字符(ASCII码)
	newS := strings.ToLower(s)
	filterStr := saveNumberAndChar(newS)

	// 折半首尾对比字符,全部相同则为回文串
	return halfHeadAndEnd(filterStr)

}

// 比较: 正则匹配效率低
func saveNumberAndChar(arr string) string {
	out := make([]byte, 0, len(arr))
	for i := range arr {
		if (arr[i] >= 97 && arr[i] <= 122) || (arr[i] >= 48 && arr[i] <= 57) {

			out = append(out, arr[i])
		}
	}
	return string(out)
}

func halfHeadAndEnd(s string) bool {
	flag := true
	chars := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		if chars[i] != chars[len(chars)-i-1] {
			flag = false
		}
	}
	return flag
}

// @lc code=end

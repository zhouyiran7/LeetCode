/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 * 自顶而上调用
 */
package leetcode

import (
	"regexp"
	"strings"
)

// @lc code=start
func IsPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	// 过滤非数字和字母的字符
	filterStr := filterNonumberAndChars(s)
	// 反转验证字符有效性
	reservedStr := reserveString(filterStr)
	return strings.ToLower(filterStr) == strings.ToLower(reservedStr)
}

func filterNonumberAndChars(s string) string {
	reg, err := regexp.Compile("[^0-9A-Za-z_]")
	if err != nil {
		panic(err)
	}

	return reg.ReplaceAllString(s, "")
}

func reserveString(s string) string {
	chars := []rune(s)
	// 反转: 首位字符对调
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func reserverString2(s string) string {
	chars := []rune(s)
	for i := 0; i < len(chars)/2; i++ {
		temp := chars[i]
		chars[i] = chars[len(chars)-i-1]
		chars[len(chars)-i-1] = temp
	}
	return string(chars)
}

// @lc code=end

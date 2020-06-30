package main

/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashMap := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
		hashMap[t[i]]--
	}

	for _, countVal := range hashMap {
		if countVal != 0 {
			return false
		}
	}

	return true
}

// @lc code=end

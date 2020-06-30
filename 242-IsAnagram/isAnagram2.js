/*
 * @lc app=leetcode.cn id=242 lang=javascript
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function (s, t) {
  if (s.length != t.length) {
    return false
  }

  let hashMap = {}

  for (let i = 0; i < s.length; i++) {
    hashMap[s.charAt(i)] ? hashMap[s.charAt(i)]++ : hashMap[s.charAt(i)] = 1
    hashMap[t.charAt(i)] ? hashMap[t.charAt(i)]-- : hashMap[t.charAt(i)] = -1
  }

  for (const key in hashMap) {
    if (hashMap[key] != 0) {
      return false
    }
  }

  return true
};
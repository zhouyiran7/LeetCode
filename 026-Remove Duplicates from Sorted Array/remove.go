/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 * 前提条件: 已排序数组
 * 要求: 原地删除
 * 官方题解: 双指针,通过一维数组索引特性, 假设两个指针游标, 块指针游标向后依次偏移,和固定
 * 在左侧的慢指针游标比对, 进行慢索引位置的元素替换
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 慢指针
	i := 0

	// 快指针右移动比较是否和慢指针位置元素相同
	for j := 1; j < len(nums); j++ {
		// 如果存在和慢指针当前位置不重复的元素,则慢指针后移置换元素为当前快指针元素
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

// @lc code=end

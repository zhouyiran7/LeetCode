package main

import (
	"fmt"
)

// 哈希表索引
// 通过哈希表充当数据临时"集装箱",成对数据进行匹配,时间复杂度O(n)
func twoSum(nums []int, target int) []int {
	// 哈希表
	sumHash := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sumHash[nums[i]] = i
	}

	// 再次遍历
	for j := 0; j < len(nums); j++ {
		maybeV := nums[j]
		// 剔除自身,进行哈希表中数据相匹配
		if index, ok := sumHash[target-maybeV]; ok && index != j {
			return []int{j, index}
		}
	}
	return nil
}

// 哈希表
// 时间复杂度O(n)
func twoSum2(nums []int, target int) []int {

	sumHash := make(map[int]int, len(nums))

	// 遍历时做假设性, 不成立时填充数据项
	for i, num := range nums {
		maybeV := target - num
		if j, ok := sumHash[maybeV]; ok && j != i {
			return []int{j, i}
		}
		sumHash[num] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 2, 7, 1}, 4))
	fmt.Println(twoSum2([]int{2, 2, 7, 1}, 4))
}

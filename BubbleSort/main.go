package main

import (
	"fmt"
)

// 冒泡排序
func bubbleSort(nums []int) []int {
	if len(nums) < 1 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	return nums
}

func main() {
	nums := []int{8, 4, 6, 5, 3, 7, 2, 1}
	nums = bubbleSort(nums)
	fmt.Println(nums)
}

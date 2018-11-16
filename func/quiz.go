package main

import (
	"fmt"
)

func practice() {
	resultAvg := avgInSlice([]float64{1.12, 32.2, 232.22})

	fmt.Println(resultAvg)

	resultList := sortNums([]int{7, 1, 20, 9, 31, 0, 3})

	fmt.Println(resultList)
}

/*
 * Q1: 求一个 [float64]slice 的平均值
 */
func avgInSlice(nums []float64) (result float64) {
	sum := 0.0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sum / (float64)(len(nums))
}

/*
* Q2: 编写函数返回两个参数的自然数顺序
*	f(7,2) -> f(2,7)
*	f(2,7) -> f(2,7)
*  扩展下 冒泡排序
 */
func sortNums(nums []int) (result []int) {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	return nums
}

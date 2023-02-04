package tasks

import "fmt"

func BubbleSort(nums []int) {
	var count int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
			count++
		}
	}
	fmt.Println(nums)
}

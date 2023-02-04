package main

import (
	"fmt"
	"go-uacademy/tasks"
)

func main() {
	// fmt.Println("calc increase from another package", homeworks.Increase(6))

	arr := []int{5, 2, 3, 1, 87}
	tasks.BubbleSort(arr)
	fmt.Println(arr)
}

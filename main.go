package main

import (
	"fmt"
	"go-uacademy/homeworks"
	// "go-uacademy/tasks"
)

func main() {


	fmt.Println(homeworks.GetOddNums([]int{2,3,4,5,6,7,8,53}))
	fmt.Println(homeworks.SumOfTheArray([]int{2,3,4,5,6,7,8,53}))
	fmt.Println(homeworks.GetSecondOddNum([]int{2,3,4,5,6,7,8,53}))
	fmt.Println(homeworks.ReverseArr([]int{2,3,4,5,6,7,8,53}))
	fmt.Println(homeworks.GetOddIndexNums([]int{2,3,4,5,6,7,8,53}))
	fmt.Println(homeworks.CompareSumOfOddAndEven([]int{2,3,4,5,6,7,8,53}))


	}

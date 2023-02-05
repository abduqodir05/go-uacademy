package main

import (
	"fmt"
	"go-uacademy/homeworks"
)

func main() {
	// home 1

	var key string = "the quick brown fox jumps over the lazy dog"
	var message string = "vkbs bs t suepuv"

	fmt.Println(key)
	fmt.Println(homeworks.DecodeMessage(key, message))

	// home 2

	fmt.Println(homeworks.SumOfUnique([]int{8, 1, 2, 2, 3}))

}

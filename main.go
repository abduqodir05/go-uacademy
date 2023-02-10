package main

import (
	"fmt"
	"go-uacademy/homeworks"
)

func main() {

	// home 1
	fmt.Println(homeworks.SumOfUniqueNums([]int{1,3,1,2}))

	// home 2
	fmt.Println(homeworks.SmallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))

	// home 3
	var key string = "the quick brown fox jumps over the lazy dog"
	var message string = "vkbs bs t suepuv"
	fmt.Println(key)
	fmt.Println(homeworks.DecodeMessage(key, message))

	// home 4
	fmt.Println(homeworks.RepeatedCharacter("abccbaacz"))

	//home 5
	fmt.Println(homeworks.UniqueOccurrences([]int{1,2,3,2}))

	
}

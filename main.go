package main

import (
	"fmt"
	"go-uacademy/homeworks"
)

func main() {

	// home 1
	day, month := 12, 2
	fmt.Println(homeworks.GetNextDayDate(day, month))

	// home-2
	fmt.Println(homeworks.CountMinNums())

	// home3
	fmt.Println(homeworks.PlusDigitToOneNum(9643))

	// home 4
	fmt.Println(homeworks.PrintDegreeOfTwo(1035))

	// home 5
	fmt.Println(homeworks.FindPerfectNum(1000))

	// home 6
	NewSet := homeworks.MyHashSet{
		Obj: []int{5, 8, 9, 1, 2, 3, 4, 5, 6},
	}
	fmt.Println(NewSet.Add(7))
	fmt.Println(NewSet.Contains(5))
	fmt.Println(NewSet.Remove(2))
}

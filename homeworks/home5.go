package homeworks

import (
	"math/rand"
	"time"
)

func GetOddIndexNums(array []int) []int {

	(rand.Seed(time.Now().UnixNano()))
	
	newArr := []int{}
	for i, v := range array {
		if i%2 != 0 {
			newArr = append(newArr, v)
		}
	}
	return newArr
}


// exercise 5
// Faqat toq indexda turgan elementlarni ekrangi chiqaruvchi function tuzing.

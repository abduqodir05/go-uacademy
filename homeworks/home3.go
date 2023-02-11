package homeworks

import (
	"math/rand"
	"time"
)

func GetSecondOddNum(arr []int) int {

	(rand.Seed(time.Now().UnixNano()))

	newSlice := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			newSlice = append(newSlice, arr[i])
		}
	}
	second := newSlice[1]
	return second
}

// exercise 3
// Shu sonlar ichidan 2-toq sonni qaytaruvchi funksiya tuzing.

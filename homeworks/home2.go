package homeworks

import (
	"math/rand"
	"time"
)

func SumOfTheArray(array []int) int {
	
	(rand.Seed(time.Now().UnixNano()))
	sum := 0

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

// exercise 2
// Shu sonlar yigindisini qaytaruvchi funksiya tuzing.
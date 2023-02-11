package homeworks

import (
	"math/rand"
	"time"
)

func ReverseArr(arr []int) []int {

	(rand.Seed(time.Now().UnixNano()))
	
	for i,j := 0, len(arr)-1; i<j; i,j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// exercise 4
// Shu sonlarni teskari tartibda chiqaring.
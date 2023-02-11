package homeworks

import (
	"math/rand"
	"time"
	"fmt"
)


func GetOddNums(arr []int) int {

	(rand.Seed(time.Now().UnixNano()))

	var flag int = 0

	fmt.Printf("Prime Numbers: \n")
	for i := 0; i <= 5; i++ {
		flag = 0
		for j := 2; j < arr[i]/2; j++ {
			if arr[i]%j == 0 {
				flag = 1
				break
			}
		}
		if flag == 0 && arr[i]>1 {
			fmt.Printf("%d ", arr[i])
		}
	}
	return flag
}

// execsise 1
// Shu sonlar ichidan faqat tublarini chiqaruvchi funksiya tuzing.
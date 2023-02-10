package homeworks

func GetSecondOddNum(arr []int) int {
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
// Shu sonlar ichidan 2-tub sonni qaytaruvchi funksiya tuzing.

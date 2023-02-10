package homeworks

func GetOddNums(arr []int)[]int  {
	newArr := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 != 0 {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// execsise 1
// Shu sonlar ichidan faqat tublarini chiqaruvchi funksiya tuzing.
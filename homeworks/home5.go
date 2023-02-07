package homeworks

func GetOddIndexNums(array []int) []int {
	newArr := []int{}
	for i, v := range array {
		if i%2 != 0 {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

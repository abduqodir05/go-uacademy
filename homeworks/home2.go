package homeworks

func CountMinNums() int {
	Myslice := []int{5, 3, 2, 1, 4, 8, 9, 1, 1, 2, 3, 4, 1, 1, 1, 1}
	myMap := make(map[int]int)

	min := Myslice[0]
	for _, v := range Myslice {
		myMap[v]++
	}
	for _, v := range myMap {
		if min > v {
			min = v
		}
	}

	return myMap[min]
}

package tasks

func GetMinCount() int {
	var nums = []int{7, 8, 1, 4, 1, 1, 3, 1}
	var getMinNums []int

	min := nums[0]
	for _, v := range nums {
		if min > v {
			min = v
		}
	}
	for _, v := range nums {
		if min == v {
			getMinNums = append(getMinNums, min)
		}
	}

	return len(getMinNums)
}

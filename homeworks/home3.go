package homeworks

func SumOfUniqueNums(nums []int) int {

	var newMap = make(map[int]int)

	for _, v := range nums {
		newMap[v]++
	}
	sum := 0
	for i, v := range newMap {	
		if v==1 {
			sum +=i
		}
	}
	return sum
}

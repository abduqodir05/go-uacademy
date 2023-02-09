package homeworks

func UniqueOccurrences(nums []int) bool {

	var newMap = make(map[int]int)

	for _, v := range nums {
		newMap[v]++
	}
	for _,v := range newMap {	
		if v>1 {
			return true
		}
	}
	return false	
}
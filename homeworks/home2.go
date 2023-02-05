package homeworks

func SumOfUnique(nums []int) int {
	uniq := []int{}
	keys := map[int]bool{}
	sum := 0

	for _, v := range nums {
		if value := keys[v]; !value {
			keys[v] = true
			uniq = append(uniq, v)
		}
	}

	for _, v := range uniq {
		v += sum
	}
	return sum
}

package homeworks

func SumOfTheArray(array []int) int {
	sum := 0

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

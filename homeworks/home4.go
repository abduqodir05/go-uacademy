package homeworks

import "fmt"

func PrintDegreeOfTwo(n int) int {

	result := 1

	for result*2 < n {
		result *= 2
		fmt.Println(result)
	}

	return result
}

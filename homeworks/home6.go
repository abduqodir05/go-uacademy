package homeworks

import "fmt"

func CompareSumOfOddAndEven(myArray []int) int {
	evenArray := []int{}
	oddArray := []int{}
	sumOdd := 0
	sumEven := 0

	for i, v := range myArray {
		if i%2 != 0 {
			evenArray = append(evenArray, v)
		}
	}

	for i := 0; i < len(evenArray); i++ {
		sumEven += evenArray[i]
	}

	for i, v := range myArray {
		if i%2 == 0 {
			oddArray = append(oddArray, v)
		}
	}

	for j := 0; j < len(oddArray); j++ {
		sumOdd += oddArray[j]
	}

	if sumEven > sumOdd {
		fmt.Println("even older")
		return sumEven
	}

	fmt.Println("odd older")
	return sumOdd
}

// 6 exercise
// Shu slicening toq indexsidagi sonlarning yig'indisi bilan juft indexsidagi
// sonlar yig'indisini taqqoslab kattasini qaytaruvchi function tuzing.

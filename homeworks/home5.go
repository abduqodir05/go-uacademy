package homeworks

func findDivisors(n int) []int {
	myArray := []int{}

	for i := 1; i < n; i++ {
		if n%i == 0 {
			myArray = append(myArray, i)
		}
	}
	return myArray
}

func sumOffArrElements(n []int) int {
	sum := 0
	for i := 0; i < len(n); i++ {
		sum += n[i]
	}
	return sum
}

func FindPerfectNum(n int) []int {
	perfectNumbers := []int{}

	for i := 1; i < n; i++ {
		// check i if it is perfect number
		divisors := findDivisors(i)
		sum := sumOffArrElements(divisors)
		if sum == i {
			perfectNumbers = append(perfectNumbers, i)
		}
	}
	return perfectNumbers
}

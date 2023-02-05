package homeworks

func PlusDigitToOneNum(n int) int {

	sum := 0

	for n > 0 || sum > 9 {
		if n == 0 {
			n = sum
			sum = 0
		}
		sum += n % 10
		n /= 10
	}
	return sum
}

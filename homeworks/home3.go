package homeworks


func GetSecondOddNum(arr []int) int {
	num1 := arr[0]
	for i := 0; i < len(arr); i++ {
		if num1 % 2 == 0{
		num1 = arr[i]
	}
}
return num1
}
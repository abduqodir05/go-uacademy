package tasks

import "fmt"

func Add(a, b, res *int) {
	res1 := *a + *b
	fmt.Println(&res1)
}
func Sub(a, b, res *int) {
	res2 := *a - *b
	fmt.Println(&res2)
}

func Div(a, b,res *int) {
	res3:= *a / *b
	fmt.Println(&res3)
}
func Mult(a, b,res *int) {
	res4:= *a * *b
	fmt.Println(&res4)
}

	package homeworks

// import "fmt"

// func main() {
// 	err, result, message := calc("-", 10, 4)

// 	if err == nil {
// 		fmt.Println(result)
// 	} else {
// 		fmt.Println(message) //6
// 	}

// }

// func calc(symbol string, a, b int) (error, int, string) {
// 	switch symbol {
// 	case "-":
// 		return nil, a - b, "successfully calculated"
// 	case "+":
// 		return nil, a + b, "added"
// 	}
// 	return fmt.Errorf("not defined symbol"), 0, "false"
// }
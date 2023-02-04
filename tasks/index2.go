package tasks

import "fmt"

func main() {
	slice1 := []int{91, 21, 21, 21, 91, 31, 91, 21}
	map1 :=make( map[int]int)

	for _, v := range slice1 {
		map1[v]++
	}
	max := slice1[0]
	for _, v := range map1 {
		if max < v {
			max = v
		} 
	}
	fmt.Println(max)
}

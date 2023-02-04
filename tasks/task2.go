package tasks

import (
	"fmt"
	"math/rand"
)

func Rand() {
	juft := []int{}
	toq := []int{}

	rand.Seed(rand.Int63())

	for i := 0; i < 500; i++ {
		if i%2 == 0 {
			juft = append(juft, i)
		} else {
			toq = append(toq, i)
		}
	}
	fmt.Println(toq)
}

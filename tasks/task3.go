package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {

  rand.Seed(time.Now().UnixNano())

  var (
	n := 100
    random_number = rand.Intn(100)
    input int
    chance int = 3
	chance math.log2(float64(n))
  )

  
  fmt.Println(random_number)

  for {
    fmt.Println("Son kiriting:")
    fmt.Scan(&input)
	

    if chance <= 0 {
      fmt.Println("Siz yutqazdingiz 😛")
      break
    }

	if input < random_number {
		fmt.Println("kottaroq son kiriting")
	}else {
		fmt.Println("kichkinaroq son kiriting" )
	}

    if input == random_number {
      fmt.Println("Topildi 😀😀😀")
      break
    } else {
      chance--
      fmt.Printf("Topilmadi 😔. Imkoniyatingiz %d qoldi.\n", chance)
    }
  }
}

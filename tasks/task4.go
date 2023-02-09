package main

import (
	"os"
	"fmt"
	"time"
	"os/exec"
)

func main()  {
	
	password := 1234
	
	var input int
	var chance int = 3
	
	for  {
		fmt.Println("enter password")
		fmt.Scan(&input)

		if input != password {
			chance--
			fmt.Printf("left %d chances\n",chance)
		}else {
			fmt.Println("Topildi 😀😀😀")
			break
		}
		
		if chance == 0 {
			for i := 0; i < 30; i++ {

				cmd := exec.Command("clear")
				cmd.Stdout = os.Stdout
				cmd.Run()
				fmt.Println("Time:", i)
				time.Sleep(time.Second * 1)
				continue
			}
		}
	}

}
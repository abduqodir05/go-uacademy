package tasks

// import "fmt"

// func main() {
	
// }

func numJewelsInStones(jewels string, stones string) int {
    
    newMap := map[byte]int{}
    count := 0
    for i :=range  stones {
        newMap[stones[i]]++
    }
    for i:= range jewels {
        if value,ok:=newMap[jewels[i]];ok{
            count+=value
        }
       
    }
    return count
}
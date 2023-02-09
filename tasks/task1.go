package main

import "fmt"

type Component struct {
	weight, sugar float64
	count int
}
type Fruit struct {
	Component
	name string
}
type Cake struct {
	Component
	name string
}

func (c Component) GetSugar() string {
return fmt.Sprintf("Sugar %fgr",c.sugar)
}
func (d Component) Total() int{
	
	return d.count
}

func main() {

	var fruits = Fruit{
		Component: Component{100.0,54.0, 12},
		name: "apple",
	}
	var cakes = Cake{
		Component: Component{12.2,53.0,12},
		name: "medove",
	}
	
	fmt.Println(fruits.GetSugar())
	fmt.Println(cakes.Total())
}

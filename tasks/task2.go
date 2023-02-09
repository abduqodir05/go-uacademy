package main

import "fmt"

type Engine struct {
	weight float64
	timeSpeed int
}
type Car struct {
	Engine
}
type Plane struct {
	Engine
}

// func (e Engine)GetSpeed()int {
// 	return e.timeSpeed
// }
func (c Car)GetSpeed()int {
	return c.Engine.timeSpeed * c.timeSpeed
}
func (p Plane)GetSpeed()int {
	return p.Engine.timeSpeed * p.timeSpeed
}
type EngineInterface interface {
	GetSpeed() int
}
func Run(Engine EngineInterface)  {
	fmt.Println(Engine.GetSpeed())
}

func main() {
	cobalt := Car{Engine{weight: 23.4,timeSpeed:900}}
	plane := Plane{Engine{weight: 63.6,timeSpeed:2100}}
	Run(cobalt)
	Run(plane)
}

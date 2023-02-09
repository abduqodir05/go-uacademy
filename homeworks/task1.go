package main

import "fmt"

type Car struct {
	engine int
	info   string
	speed  int
}
type Merc struct {
	Car
}
type BMW struct {
	Car
}

func (c Car) GetSpeed() int {
	return c.speed
}
func (m Merc) GetSpeed() int {
	return m.Car.speed * m.speed
}
func (b BMW) GetSpeed() int {
	return b.Car.speed * b.speed
}
func (c Car) GetEngine() int {
	return c.engine
}
func (m Merc) GetEngine() int {
	return m.Car.GetEngine()
}
func (b BMW) GetEngine() int {
	return b.Car.GetEngine()
}
func (m Merc) GetInfo() string {
	return fmt.Sprintln(m.Car.info)
}
func (b BMW) GetInfo() string {
	return fmt.Sprintln(b.Car.info)
}

type CarInterface interface {
	GetSpeed() int
	GetEngine() int
	GetInfo() string
}

func aboutCar(car CarInterface) {
	fmt.Printf("info " + car.GetInfo())
	fmt.Printf("Engine %d\n",car.GetEngine())
	fmt.Printf("Speed %d\n",car.GetSpeed())
}

func main() {
	bmw := (BMW{Car{info: "BMW", engine: 4, speed: 320}})
	merc := (Merc{Car{info: "Merc", engine: 3, speed: 300}})
	aboutCar(bmw)
	aboutCar(merc)
}
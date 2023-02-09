package main

import "fmt"

type Os struct {
	name string
}
type MacOs struct {
	Os
}
type Linux struct {
	Os
}
type Windows struct {
	Os
}

func (o Os) GetName()string  {
	return o.name
}

func (m MacOs) GetName() string {
	return m.Os.GetName() + " M2"
}

func (w Windows) GetName()string {
	return w.Os.GetName() + "Max"
}

type OsInterface interface {
	GetName() string
}

func Run(os OsInterface)  {
	fmt.Println("Loading " + os.GetName() + "...")
}

func main()  {
	Run(MacOs{Os{name: "Macintosh"}})
	Run(Linux{Os{name: "Ubuntu"}})
	Run(Windows{Os{name: "windows"}})
}
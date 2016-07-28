package main

import (
	"fmt"
)

type Pizza struct {
	Color string
}
type Cooker interface {
	Prepare(*Pizza)
	Cook(*Pizza)
}
type DefaultCooker struct {
	Name string
}

func (c *DefaultCooker) Prepare(p *Pizza) {
	fmt.Println("Default cooker prepare :", p.Color)
}
func (c *DefaultCooker) Cook(p *Pizza) {
	fmt.Println("Default cooker cook :", p.Color)
}

type MyDefaultCooker struct {
	DefaultCooker
}

func (c *MyDefaultCooker) Prepare(p *Pizza) {
	fmt.Println("My cooker prepare :", p.Color)
}
func (c *MyDefaultCooker) Cook(p *Pizza) {
	fmt.Println("My cooker cook :", p.Color)
}
func main() {
	pizza := &Pizza{Color: "yellow"}

	var c Cooker

	c1 := &DefaultCooker{}
	c = c1
	c.Prepare(pizza)
	c.Cook(pizza)

	c2 := &MyDefaultCooker{}
	c = c2
	c.Prepare(pizza)
	c.Cook(pizza)

}

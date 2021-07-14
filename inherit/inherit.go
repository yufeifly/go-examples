package main

import "fmt"

type Parent struct{}

func (p *Parent) Test() {
	p.Inner()
}

func (p *Parent) Inner() {
	fmt.Println("base inner")
}

type Child struct {
	Parent
}

func (c *Child) Inner() {
	fmt.Println("child inner")
}

func main() {
	p := &Parent{}
	c := &Child{}

	p.Inner()
	p.Test()
	c.Inner()
	c.Test()
}

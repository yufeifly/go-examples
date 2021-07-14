package main

import (
	"fmt"
	"github.com/yufeifly/go-examples/xiface/animal"
)

func main() {
	c := &animal.Cat{}
	d := &animal.Dog{}
	h := &animal.Human{}
	animals := []animal.Animal{c, d, h}
	for _, animal := range animals {
		speak(animal)
	}
}

func speak(s animal.Animal) {
	fmt.Println(s.Speak())
}

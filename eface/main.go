package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Test struct {
		name string
		val  int
	}
	v := Test{name: "yufei", val: 100}
	Print(v)

}

func Print(v interface{}) {
	print(v)
	fmt.Printf("%v\n", reflect.ValueOf(v))
	fmt.Printf("%v\n", reflect.TypeOf(v))
}

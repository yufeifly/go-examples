/*
package main
https://mp.weixin.qq.com/s/2eED64tnb3bDong_3_HJoA
*/
package main

import (
	"fmt"
)

type worker interface {
	work()
}

type person struct {
	name string
	worker
}

func (p *person) work() {
	fmt.Printf("%s is working\n", p.name)
}

func main() {
	p := &person{} // its members will be init by default
	p.worker = &person{
		name:   "yufei",
		worker: nil,
	}
	p.worker.work()
}

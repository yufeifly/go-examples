package main

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

var lpool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 1024)
	},
}

func copyBuffer(dst, src io.ReadWriter) error {
	buf := lpool.Get().([]byte)
	defer lpool.Put(buf)

	_, err := io.CopyBuffer(dst, src, buf)
	return err
}

func transport(rw1, rw2 io.ReadWriter) error {
	errc := make(chan error, 1)

	go func() {
		errc <- copyBuffer(rw1, rw2)
		defer func() {
			fmt.Println("fuck rw2 => rw1")
		}()
	}()

	go func() {
		errc <- copyBuffer(rw2, rw1)
		defer func() {
			fmt.Println("fuck rw1 => rw2")
		}()
	}()

	err := <-errc

	if err != nil && err == io.EOF {
		err = nil
	}
	fmt.Println("rw1: ", rw1)
	fmt.Println("rw2: ", rw2)
	return err
}

func main() {
	lpool.Put([]byte("fuck you, alan"))
	lpool.Put([]byte("fuck you, bob"))
	lpool.Put([]byte("fuck you, cindy"))
	rw1 := bytes.Buffer{}
	rw1.Write([]byte("hello"))
	rw2 := bytes.Buffer{}
	rw2.Write([]byte("world"))
	err := transport(&rw1, &rw2)
	if err != nil {
		fmt.Println("fuck, ", err)
	}
	rw3 := bytes.Buffer{}
	rw3.Write([]byte("hi"))
	rw4 := bytes.Buffer{}
	rw4.Write([]byte("yufei"))
	b, ok := lpool.Get().([]byte)
	if ok == false {
		return
	}
	_, errC := io.CopyBuffer(&rw3, &rw4, b)
	if errC != nil {
		fmt.Println("errC: ", errC)
	}
	fmt.Println("rw3: ", rw3.String())
	fmt.Println("rw4: ", rw4.String())
}

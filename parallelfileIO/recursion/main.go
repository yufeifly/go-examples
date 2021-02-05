package main

import (
	"fmt"
	"os"
)

func lsFiles(dir string) {
	file, err := os.Open(dir)
	if err != nil {
		fmt.Println("error opening directory")
	}
	defer file.Close()
	files, err := file.Readdir(-1)
	if err != nil {
		fmt.Println("error reading directory")
	}
	for _, f := range files {
		if f.IsDir() {
			lsFiles(dir + "/" + f.Name())
		}
		fmt.Println(dir + "/" + f.Name())
	}
}

func main() {
	//err := filepath.Walk(os.Args[1], func(path string, fi os.FileInfo, err error) error {
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Println(path)
	//	return nil
	// })
	//if err != nil {
	//	log.Fatal(err)
	//}
	lsFiles(os.Args[1])
}

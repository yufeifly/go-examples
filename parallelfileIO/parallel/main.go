package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

func main() {
	//
	logName := "log"
	logFile, err := os.Create(logName)
	if err != nil {
		logrus.Fatalf("os create failed, err: %v", err)
	}
	defer logFile.Close()
	logrus.SetOutput(logFile)
	//debugLog = log.New(logFile,"[Debug]",log.LstdFlags)

	dir := os.Args[1]
	fmt.Println(dir)
	wg.Add(1)
	lsFiles(dir)
	wg.Wait()
	fmt.Printf("total files: %v\n", fileCount)
}

const INT_MAX = int(^uint(0) >> 1)
const INT_MIN = ^INT_MAX

var fileCount int
var lock sync.Mutex

var wg sync.WaitGroup

var (
//concurrent    = runtime.GOMAXPROCS(runtime.NumCPU())
//semaphoreChan = make(chan struct{}, concurrent)
)

func lsFiles(dir string) {
	// block while full
	//semaphoreChan <- struct{}{}
	go func() {
		defer func() {
			// read to release a slot
			//<-semaphoreChan
			wg.Done()
		}()

		file, err := os.Open(dir)
		if err != nil {
			fmt.Println("error opening directory", err.Error())
			logrus.Printf("error opening directory %v", err.Error())
			lock.Lock()
			fileCount = INT_MIN
			lock.Unlock()
		}
		defer file.Close()
		files, err := file.Readdir(-1) // Loads all children files into memory. A more efficient way?
		if err != nil {
			fmt.Println("error reading directory", err.Error())
			logrus.Printf("error reading directory %v", err.Error())
			lock.Lock()
			fileCount = INT_MIN
			lock.Unlock()
		}
		for _, f := range files {
			lock.Lock()
			fileCount++
			lock.Unlock()
			fmt.Println(dir + "/" + f.Name())
			if f.IsDir() {
				wg.Add(1)
				go lsFiles(dir + "/" + f.Name())
			}
		}
	}()
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func adjustName(Names string) string {
	names := strings.Split(Names, ",")
	var newNames []string
	for i, p := range names {
		names[i] = strings.TrimSpace(p)
		pieces := strings.Split(names[i], " ")
		s := 0
		t := len(pieces) - 1
		for s < t {
			pieces[s], pieces[t] = pieces[t], pieces[s]
			s++
			t--
		}
		for ind := range pieces[:len(pieces)-1] {
			pieces[ind] += "."
		}
		var newName string
		for _, piece := range pieces[:len(pieces)-1] {
			newName += piece + " "
		}
		newName += pieces[len(pieces)-1]
		newNames = append(newNames, newName)
	}
	var outputNames string
	for _, newN := range newNames[:len(newNames)-1] {
		outputNames += newN + ", "
	}
	if newNames != nil {
		outputNames += newNames[len(newNames)-1]
	}
	return outputNames
}

func main() {
	fi, err := os.Open("names.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(adjustName(string(a)))
	}
}

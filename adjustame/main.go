package main

import (
	"fmt"
	"strings"
)

func adjustName(Names []string) {
	for _, rawName := range Names {
		names := strings.Split(rawName, ",")
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
			for ind := range pieces {
				pieces[ind] += "."
			}
			var newName string
			for _, piece := range pieces {
				newName += piece
			}
			fmt.Println(newName)
		}

	}
}

func main() {
	names := []string{
		"Puliafito C, Vallati C, Mingozzi E",
	}
	adjustName(names)
}

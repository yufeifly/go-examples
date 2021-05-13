package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Age > people[j].Age }) // 按年龄降序排序
	fmt.Println("Sort by age:", people)
	fmt.Println("Sorted:", sort.SliceIsSorted(people, func(i, j int) bool { return people[i].Age > people[j].Age }))
}

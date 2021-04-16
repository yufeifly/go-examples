package main

import "fmt"

// time : O(n*logn)
func HeapSortN(nums []int) []int {
	n := len(nums)
	buildHeapN(nums)

	for i := n - 1; i >= 0; i-- {
		swapN(nums, 0, i)
		adjustHeapN(nums, 0, i)
	}
	//fmt.Println(nums)
	return nums
}

func buildHeapN(a []int) {
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		adjustHeapN(a, i, n)
	}
}

func adjustHeapN(nums []int, i, n int) {
	for {
		j := i*2 + 1
		if j >= n {
			break
		}
		if j+1 < n && nums[j+1] > nums[j] {
			j = j + 1
		}
		if nums[i] < nums[j] {
			swapN(nums, i, j)
		} else {
			break
		}
		i = j
	}
}

func swapN(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	vec := []int{4, 3, 2, 100, 6, 7, 9}
	vec = HeapSortN(vec)
	fmt.Println(vec)
	//vec = append(vec, -1)
	//fmt.Println(vec)
	//vec = up(vec, len(vec)-1)
	//fmt.Println(vec)
}

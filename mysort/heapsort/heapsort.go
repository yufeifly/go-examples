package main

// time : O(n*n)
func heapSort(nums []int, n int) {
	for i := n/2 - 1; i >= 0; i-- {
		j := i*2 + 1
		if j+1 < n && nums[j+1] > nums[j] {
			j = j + 1
		}
		if nums[j] > nums[i] {
			nums[j], nums[i] = nums[i], nums[j]
		}
	}
	nums[0], nums[n-1] = nums[n-1], nums[0]
}

func HeapSort(nums []int) []int {
	n := len(nums)
	for i := n; i > 0; i-- {
		heapSort(nums, i)
	}
	//fmt.Println(nums)
	return nums
}

func up(vec []int, j int) []int {
	for {
		i := (j - 1) / 2
		if i == j || vec[j] >= vec[i] {
			break
		}
		vec[i], vec[j] = vec[j], vec[i]
		j = i
	}
	return vec
}

//func main() {
//	vec := []int{4,3,2,100,6,7,9}
//	vec = HeapSort(vec)
//	fmt.Println(vec)
//	//vec = append(vec, -1)
//	//fmt.Println(vec)
//	//vec = up(vec, len(vec)-1)
//	//fmt.Println(vec)
//}

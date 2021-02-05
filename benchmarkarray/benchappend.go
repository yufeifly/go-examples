package benchmarkarray

func AppendTail() {
	var arr []int
	for i := 0; i < 100; i++ {
		arr = append(arr, i)
	}
}

func AppendHead() {
	var arr []int
	for i := 0; i < 100; i++ {
		arr = append([]int{i}, arr...)
	}
}

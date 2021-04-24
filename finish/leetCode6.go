package main

import "fmt"

type Testcase struct {
	nums []int
	n    int
}

func shuffle(nums []int, n int) []int {
	newArr := []int{}
	x := []int{}
	y := []int{}
	for i := 0; i < n; i++ {
		x = append(x, nums[i])
	}
	for i := n; i < 2*n; i++ {
		y = append(y, nums[i])
	}
	for i := 0; i < n; i++ {
		newArr = append(newArr, x[i])
		newArr = append(newArr, y[i])
	}
	return newArr
}

func main() {
	case1 := Testcase{
		nums: []int{2, 5, 1, 3, 4, 7},
		n:    3,
	}
	case2 := Testcase{
		nums: []int{1, 2, 3, 4, 4, 3, 2, 1},
		n:    4,
	}
	case3 := Testcase{
		nums: []int{1, 1, 2, 2},
		n:    2,
	}

	result1 := shuffle(case1.nums, case1.n)
	result2 := shuffle(case2.nums, case2.n)
	result3 := shuffle(case3.nums, case3.n)

	fmt.Printf("Output case 1: %v\n", result1)
	fmt.Printf("Output case 2: %v\n", result2)
	fmt.Printf("Output case 3: %v\n", result3)
}

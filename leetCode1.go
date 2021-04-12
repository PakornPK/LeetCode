package main

import "fmt"

func runningSum(nums []int) []int {
	runningSun := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j <= i; j++ {
			runningSun[i] = runningSun[i] + nums[j]
		}
	}
	return runningSun
}

func main() {
	var case1 = []int{1, 2, 3, 4}
	var case2 = []int{1, 1, 1, 1, 1}
	var case3 = []int{3, 1, 2, 10, 1}

	result1 := runningSum(case1)
	result2 := runningSum(case2)
	result3 := runningSum(case3)

	fmt.Printf("Case 1 : %v \n", result1)
	fmt.Printf("Case 2 : %v \n", result2)
	fmt.Printf("Case 3 : %v \n", result3)
}

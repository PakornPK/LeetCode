package main

import (
	"fmt"
	"sort"
)

type Testcase struct {
	numbers []int
	target  int
}

func twoSum(numbers []int, target int) []int {
	sort.Ints(numbers)
	left := 0
	right := len(numbers) - 1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{}
}

func main() {
	case1 := Testcase{
		numbers: []int{2, 7, 11, 15},
		target:  9,
	}
	case2 := Testcase{
		numbers: []int{2, 3, 4},
		target:  6,
	}
	case3 := Testcase{
		numbers: []int{-1, 0},
		target:  -1,
	}

	result1 := twoSum(case1.numbers, case1.target)
	result2 := twoSum(case2.numbers, case2.target)
	result3 := twoSum(case3.numbers, case3.target)

	fmt.Printf("Output case 1 : %v\n", result1)
	fmt.Printf("Output case 2 : %v\n", result2)
	fmt.Printf("Output case 3 : %v\n", result3)

}

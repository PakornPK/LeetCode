package main

import (
	"fmt"
)

type Testcase struct {
	nums   []int
	target int
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		v, ok := m[temp]
		if ok {
			return []int{v, i}
		} else {
			m[nums[i]] = i
		}
	}
	return []int{}
}

func main() {
	case1 := Testcase{
		nums:   []int{2, 7, 11, 15},
		target: 9,
	}
	case2 := Testcase{
		nums:   []int{3, 2, 4},
		target: 6,
	}
	case3 := Testcase{
		nums:   []int{3, 3},
		target: 6,
	}

	result1 := twoSum(case1.nums, case1.target)
	result2 := twoSum(case2.nums, case2.target)
	result3 := twoSum(case3.nums, case3.target)

	fmt.Printf("Output case 1 : %v\n", result1)
	fmt.Printf("Output case 2 : %v\n", result2)
	fmt.Printf("Output case 3 : %v\n", result3)
}

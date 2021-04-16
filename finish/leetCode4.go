package main

import (
	"fmt"
	"sort"
)

func maximumWealth(accounts [][]int) int {
	richest := make([]int, len(accounts))
	for i := 0; i < len(accounts); i++ {
		for _, v := range accounts[i] {
			richest[i] = richest[i] + v
		}
	}
	fmt.Printf("%v\n", richest)
	sort.Ints(richest)
	return richest[len(richest)-1]
}

func main() {
	case1 := [][]int{{1, 2, 3}, {3, 2, 1}}
	case2 := [][]int{{1, 5}, {7, 3}, {3, 5}}
	case3 := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}

	result1 := maximumWealth(case1)
	result2 := maximumWealth(case2)
	result3 := maximumWealth(case3)

	fmt.Printf("Output case 1 : %v\n", result1)
	fmt.Printf("Output case 2 : %v\n", result2)
	fmt.Printf("Output case 3 : %v\n", result3)
}

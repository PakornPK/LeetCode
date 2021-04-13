package main

import (
	"fmt"
	"sort"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	bufArr := make([]int, len(candies))
	for i := 0; i < len(bufArr); i++ {
		bufArr[i] = candies[i]
	}
	sort.Ints(bufArr)
	lowSite := bufArr[len(bufArr)-1] - extraCandies
	for i := 0; i < len(candies); i++ {
		if candies[i] < lowSite {
			result[i] = false
		} else {
			result[i] = true
		}
	}
	return result
}

type testCase struct {
	candies      []int
	extraCandies int
}

func main() {
	case1 := testCase{
		candies:      []int{2, 3, 5, 1, 3},
		extraCandies: 3,
	}
	case2 := testCase{
		candies:      []int{4, 2, 1, 1, 2},
		extraCandies: 1,
	}
	case3 := testCase{
		candies:      []int{12, 1, 12},
		extraCandies: 10,
	}

	result1 := kidsWithCandies(case1.candies, case1.extraCandies)
	result2 := kidsWithCandies(case2.candies, case2.extraCandies)
	result3 := kidsWithCandies(case3.candies, case3.extraCandies)

	fmt.Printf("Result case 1 : %v\n", result1)
	fmt.Printf("Result case 2 : %v\n", result2)
	fmt.Printf("Result case 3 : %v\n", result3)
}

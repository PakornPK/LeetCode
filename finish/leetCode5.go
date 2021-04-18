package main

import (
	"fmt"
)

func countPairs(deliciousness []int) int {
	ans := 0
	maphash := make(map[int]int)
	for i := 0; i < len(deliciousness); i++ {
		maphash[i] = deliciousness[i]
	}
	k := 0
	for j := 1; j <= len(maphash); j++ {
		if j == len(maphash) {
			k = k + 1
			j = k + 1
		}
		if _, ok := maphash[j]; !ok {
			continue
		}
		temp := maphash[k] + maphash[j]
		if temp != 0 && ((temp & (temp - 1)) == 0) {
			ans = ans + 1
		}
	}
	return ans
}

func main() {
	case1 := []int{1, 3, 5, 7, 9}
	case2 := []int{1, 1, 1, 3, 3, 3, 7}
	case3 := []int{2160, 1936, 3, 29, 27, 5, 2503, 1593, 2, 0, 16, 0, 3860, 28908, 6, 2, 15, 49, 6246, 1946, 23, 105, 7996, 196, 0, 2, 55, 457, 5, 3, 924, 7268, 16, 48, 4, 0, 12, 116, 2628, 1468}
	case4 := []int{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64}
	case5 := []int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234}

	result1 := countPairs(case1)
	result2 := countPairs(case2)
	result3 := countPairs(case3)
	result4 := countPairs(case4)
	result5 := countPairs(case5)

	fmt.Printf("Output case 1 : %v\n", result1)
	fmt.Printf("Output case 2 : %v\n", result2)
	fmt.Printf("Output case 3 : %v\n", result3)
	fmt.Printf("Output case 4 : %v\n", result4)
	fmt.Printf("Output case 5 : %v\n", result5)
}

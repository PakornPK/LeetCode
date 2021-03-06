package main

import (
	"fmt"
	"math"
)

func pow2needadd(x int) int {
	ans := 1
	for x > ans {
		ans <<= 1
	}
	return ans - x
}

func countPairs(deliciousness []int) int {
	// const var
	modVal := int(math.Pow10(9) + 7)

	// initial sqrtList (all 2^0 to 2^21)
	var sqrtList [22]int
	for i := 0; i <= 21; i++ {
		sqrtList[i] = int(math.Pow(2, float64(i)))
	}

	// initial deliciousness food as map (key = deliciousness, val = count)
	deliciousnessMap := make(map[int]int)
	for _, v := range deliciousness {
		deliciousnessMap[v]++
	}

	// count combinations
	count := 0

	// loop over and calculate all possible sqrt in map
	for d, v := range deliciousnessMap {
		for _, sq := range sqrtList {
			if d > sq {
				// shortcut break to run a bit faster if delicious already over sq
				continue
			}
			// get possible pair
			pair := sq - d

			// if pair equals itself, then find all their combinations
			if v > 0 && pair == d {
				count = (count + (v * (v - 1) / 2)) % modVal
				continue
			}

			// if already contains food with equals pair, then multiply it
			if pairCount, exists := deliciousnessMap[pair]; exists {
				count = (count + (v * pairCount)) % modVal
				continue
			}
		}
		// remove map so it will not count duplicate pair
		delete(deliciousnessMap, d)
	}

	return count % modVal
}

func main() {
	// case1 := []int{1, 3, 5, 7, 9}
	case2 := []int{1, 1, 1, 3, 3, 3, 7}
	// case3 := []int{2160, 1936, 3, 29, 27, 5, 2503, 1593, 2, 0, 16, 0, 3860, 28908, 6, 2, 15, 49, 6246, 1946, 23, 105, 7996, 196, 0, 2, 55, 457, 5, 3, 924, 7268, 16, 48, 4, 0, 12, 116, 2628, 1468}
	// case4 := []int{64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64, 64}
	// case5 := []int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234}

	// result1 := countPairs(case1)
	result2 := countPairs(case2)
	// result3 := countPairs(case3)
	// result4 := countPairs(case4)
	// result5 := countPairs(case5)

	// fmt.Printf("Output case 1 : %v\n", result1)
	fmt.Printf("Output case 2 : %v\n", result2)
	// fmt.Printf("Output case 3 : %v\n", result3)
	// fmt.Printf("Output case 4 : %v\n", result4)
	// fmt.Printf("Output case 5 : %v\n", result5)
}

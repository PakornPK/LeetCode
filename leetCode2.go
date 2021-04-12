package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	ipSplit := strings.Split(address, ".")
	return fmt.Sprintf("%s[.]%s[.]%s[.]%s",ipSplit[0],ipSplit[1],ipSplit[2],ipSplit[3])
}

func main() {
	case1 := "1.1.1.1"
	case2 := "255.100.50.0"

	result1 := defangIPaddr(case1)
	result2 := defangIPaddr(case2)

	fmt.Printf("Output case 1 : %s\n", result1)
	fmt.Printf("Output case 2 : %s\n", result2)
}

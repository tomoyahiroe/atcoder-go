package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		s   string
		sum int64
	)
	fmt.Scanf("%s", &s)
	numbers := strings.Split(s, "")
	for _, number := range numbers {
		num, _ := strconv.ParseInt(number, 10, 64)
		sum += num
		// fmt.Printf("%s\n", number)
	}
	// fmt.Printf("%v", numbers)
	fmt.Printf("%d\n", sum)

	// switch int(s/100) + int(s%100/10) + int(s%10) {
	// case 0:
	// 	fmt.Printf("0\n")
	// case 1:
	// 	fmt.Printf("1\n")
	// case 2:
	// 	fmt.Printf("2\n")
	// case 3:
	// 	fmt.Printf("3\n")
	// default:
	// 	fmt.Printf("invalid input")
	// }
}

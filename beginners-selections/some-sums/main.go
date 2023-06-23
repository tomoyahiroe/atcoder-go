package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, a, b int
	output := 0
	fmt.Scanf("%v %v %v", &n, &a, &b)
	for i := 1; i < n+1; i++ {
		s := strconv.Itoa(i)
		arr := strings.Split(s, "")
		sum := 0
		for _, v := range arr {
			str, _ := strconv.Atoi(v)
			sum = sum + str
		}
		if a-1 < sum && sum < b+1 {
			output = output + i
		}
	}
	fmt.Println(output)
}

package main

import "fmt"

func main() {
	var a, b, c, x int
	count := 0

	fmt.Scanf("%v\n%v\n%v\n%v", &a, &b, &c, &x)
	for i := 0; i < a+1; i++ {
		for j := 0; j < b+1; j++ {
			for k := 0; k < c+1; k++ {
				if 500*i+100*j+50*k == x {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, alice, bob int
	fmt.Scanf("%d", &N)
	array := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanf("%d", &array[i])
	}
	sort.Slice(array, func(i, j int) bool {
		return array[i] > array[j]
	})
	for i, v := range array {
		if i%2 == 0 {
			alice = alice + v
		} else {
			bob = bob + v
		}
	}
	fmt.Println(alice - bob)
}

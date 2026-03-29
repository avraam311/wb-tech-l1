package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	bSet := make(map[int]struct{})
	for _, num := range b {
		bSet[num] = struct{}{}
	}

	var res []int
	for _, num := range a {
		if _, exists := bSet[num]; exists {
			res = append(res, num)
		}
	}

	fmt.Println(res)
}

package main

import (
	"fmt"
)

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	si := []int{15, 20, 45, 67, 22}
	fmt.Println(Index(si, 22))

	ss := []string{"money", "Dolars", "Soms"}
	fmt.Println(Index(ss, "kaka"))
}

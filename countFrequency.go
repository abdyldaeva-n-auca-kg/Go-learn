package main

import "fmt"

func countFrequency(arr []int) map[int]int {
	frequency := make(map[int]int)
	for _, num := range arr {
		frequency[num]++
	}
	return frequency
}
func main() {
	arr := []int{1, 2, 3, 4, 3, 1, 1}
	frequency := countFrequency(arr)
	for num, count := range frequency {
		fmt.Printf("%d occurs %d times\n", num, count)
	}
}

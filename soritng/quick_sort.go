package main

import (
	"fmt"
)

func main() {
	var arrayTosort_tmp = []int{2, 11, 6, 3, 4, 7, 8, 5}
	fmt.Println(len(arrayTosort_tmp))
	pivot := arrayTosort_tmp[len(arrayTosort_tmp)-1]
	fmt.Println("Pivot Value: ", pivot)
	current_index := 0
	tail_index := 0

	for current_index < len(arrayTosort_tmp)-1 {
		if arrayTosort_tmp[current_index] < pivot {

		}
	}
}

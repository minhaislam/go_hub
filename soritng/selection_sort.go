package main

import (
	"fmt"
)

func main() {

	var arrayTosort = []int{9, 1, 6, 8, 4, 3, 7, 5}
	fmt.Println(arrayTosort)

	for j := 0; j < len(arrayTosort); j++ {
		current_minimum := arrayTosort[j]
		current_item_index := j
		current_item := arrayTosort[j]
		k := j + 1
		for k < len(arrayTosort) {

			if current_item > arrayTosort[k] {
				current_item = arrayTosort[k]
				current_item_index = k
			}
			k++
		}
		fmt.Println("Current Minimum: ", current_minimum)
		fmt.Println("Current Item: ", current_item)
		arrayTosort[current_item_index] = current_minimum
		arrayTosort[j] = current_item
		fmt.Println("After Sort: ", arrayTosort)

	}
}

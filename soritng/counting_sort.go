package main

import (
	"fmt"
)

func main() {

	var arrayTosort = []int{9, 1, 6, 1, 8, 4, 3, 7, 9}
	var max_range = 9
	var min_range = 0
	array_len := len(arrayTosort)
	start_range := min_range
	init_array := make([]int, max_range+1)
	fmt.Println(arrayTosort)
	current_index := 0
	i := start_range
	for i < len(arrayTosort) {

		init_array[arrayTosort[i]] = init_array[arrayTosort[i]] + 1
		i++
	}
	fmt.Println(init_array)

	final_array := make([]int, array_len)
	j := start_range
	fmt.Println("Start Range: ", j)

	for j < len(init_array) {
		fmt.Println("value is: ", init_array[j], " and Index is: ", j)
		k := init_array[j]

		fmt.Println("value of k: ", k)
		for k > 0 && init_array[j] != 0 {
			fmt.Println("intial Value", final_array[k])
			final_array[current_index] = j
			fmt.Println("Assigned Value", final_array[k])
			current_index++
			k--
		}
		fmt.Println("After each sort:", final_array)
		j++
		// break
	}
	fmt.Println("Sorted Array: ", final_array)
}

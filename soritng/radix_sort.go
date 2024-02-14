package main

import (
	"fmt"
)

func main() {

	var arrayTosort_tmp = []int{921, 11, 6, 11, 8, 4, 39, 7, 92}
	max_val := find_max(arrayTosort_tmp)
	fmt.Println(max_val)
	// i := 0
	for exp_data := 1; max_val/exp_data > 0; exp_data = exp_data * 10 {
		arrayTosort_tmp = counting_sort(arrayTosort_tmp, exp_data)
	}
}

func find_max(arrayTosort1 []int) (current_max int) {
	fmt.Println(arrayTosort1)
	current_max = arrayTosort1[0]
	for max_index := 1; max_index < len(arrayTosort1); max_index++ {
		if current_max < arrayTosort1[max_index] {
			current_max = arrayTosort1[max_index]
		}

	}
	fmt.Println("Maximum value:", current_max)
	return
}

func counting_sort(arrayTosort []int, exp_data int) (return_array []int) {
	// var arrayTosort =
	var max_range = 9
	var min_range = 0
	array_len := len(arrayTosort)
	start_range := min_range
	init_array := make([]int, max_range+1)
	current_index := 0
	i := start_range
	for i < len(arrayTosort) {

		init_array[(arrayTosort[i]/exp_data)%10] = init_array[(arrayTosort[i]/exp_data)%10] + 1

		i++
	}
	fmt.Println(init_array)

	final_array := make([]int, array_len)
	j := start_range
	// fmt.Println("Start Range: ", j)

	for j < len(init_array) {

		start_index_val := 0
		for start_index_val < len(arrayTosort) && init_array[j] != 0 {
			// fmt.Println("intial Value", final_array[k])
			if (arrayTosort[start_index_val]/exp_data)%10 == j {
				final_array[current_index] = arrayTosort[start_index_val]
				current_index++
			}
			start_index_val++
		}
		fmt.Println("After each sort:", final_array)
		j++
		// break
	}
	return_array = final_array
	return
	// for sort_index
}

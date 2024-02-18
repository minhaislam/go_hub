package main

import (
	"fmt"
)

func main() {

	var arrayTosort_tmp = []int{2, 11, 6, 3, 921, 3, 4, 7, 8, 1}
	fmt.Println(len(arrayTosort_tmp) / 2)
	div_value := len(arrayTosort_tmp) / 2

	final_sort := merge_sort(arrayTosort_tmp, div_value)
	// merge(array_partition1, array_partition2)
	fmt.Println(final_sort)
}

func merge_sort(arrayTosort_tmp []int, div_value int) (get_array []int) {

	array_partition1 := arrayTosort_tmp[:div_value]
	array_partition2 := arrayTosort_tmp[div_value:]

	first_partition := array_partition1
	second_partition := array_partition2
	fmt.Println(first_partition)
	fmt.Println(second_partition)
	if len(array_partition1) >= 2 && len(array_partition2) >= 2 {
		first_partition = merge_sort(array_partition1, len(array_partition1)/2)
		second_partition = merge_sort(array_partition2, len(array_partition2)/2)

	} else if len(array_partition1) >= 2 && len(array_partition2) < 2 {
		first_partition = merge_sort(array_partition1, len(array_partition1)/2)
		second_partition = array_partition2
	} else if len(array_partition1) < 2 && len(array_partition2) >= 2 {
		first_partition = array_partition1
		second_partition = merge_sort(array_partition2, len(array_partition2)/2)
	}
	fmt.Println("to Merge:", first_partition, "-->", second_partition)
	get_array = merge(first_partition, second_partition)
	fmt.Println(first_partition)
	fmt.Println(second_partition)
	return
}

func merge(array_partition1 []int, array_partition2 []int) (final_array []int) {
	array1_index := 0
	array2_index := 0
	main_index := 0
	init_array := make([]int, len(array_partition1)+len(array_partition2))

	for array1_index < len(array_partition1) && array2_index < len(array_partition2) {

		fmt.Println("---------")
		if array_partition1[array1_index] < array_partition2[array2_index] {
			init_array[main_index] = array_partition1[array1_index]
			array1_index++
			main_index++
		} else {
			init_array[main_index] = array_partition2[array2_index]
			array2_index++
			main_index++
		}

	}

	for array1_index < len(array_partition1) {
		init_array[main_index] = array_partition1[array1_index]
		array1_index++
		main_index++
	}

	for array2_index < len(array_partition2) {
		init_array[main_index] = array_partition2[array2_index]
		array2_index++
		main_index++
	}

	final_array = init_array

	return

}

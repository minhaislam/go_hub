package main

import (
	"fmt"
)

func main() {
	var arrayTosort_tmp = []int{2, 11, 6, 3, 3, 4, 7, 8, 5}
	// var arrayTosort_tmp = []int{2, 5}
	sorted_array := quick_sort(arrayTosort_tmp)
	fmt.Println("Sorted: ", sorted_array)
}

func quick_sort(arrayTosort_tmp1 []int) (get_array []int) {
	var arrayTosort_tmp = arrayTosort_tmp1
	fmt.Println(len(arrayTosort_tmp))
	pivot := arrayTosort_tmp[len(arrayTosort_tmp)-1]
	fmt.Println("Pivot Value: ", pivot)
	current_index := 0
	tail_index := -1

	for current_index < len(arrayTosort_tmp)-1 {
		if arrayTosort_tmp[current_index] < pivot {
			// fmt.Println(arrayTosort_tmp[current_index], "------------>", pivot)
			tail_index = tail_index + 1
			fmt.Println(arrayTosort_tmp[current_index], "------------>", arrayTosort_tmp[tail_index])
			temp_val := arrayTosort_tmp[current_index]
			arrayTosort_tmp[current_index] = arrayTosort_tmp[tail_index]
			arrayTosort_tmp[tail_index] = temp_val
			current_index = current_index + 1
			// tail_index = tail_index + 1
		} else if arrayTosort_tmp[current_index] >= pivot {
			current_index = current_index + 1
		}
	}
	fmt.Println("Current_index", current_index)
	fmt.Println("tail Index", tail_index)
	fmt.Println("Before shuffle: ", arrayTosort_tmp)
	temp_val := arrayTosort_tmp[len(arrayTosort_tmp)-1]
	arrayTosort_tmp[len(arrayTosort_tmp)-1] = arrayTosort_tmp[tail_index+1]
	arrayTosort_tmp[tail_index+1] = temp_val
	fmt.Println("After shuffle: ", arrayTosort_tmp)
	fmt.Println("left division", arrayTosort_tmp[:tail_index+1])
	fmt.Println("Right division", arrayTosort_tmp[tail_index+2:])
	fmt.Println("Right division", len(arrayTosort_tmp[tail_index+2:]))
	first_partition := arrayTosort_tmp[:tail_index+1]
	second_partition := arrayTosort_tmp[tail_index+2:]
	mid_partition := []int{arrayTosort_tmp[tail_index+1]}
	fmt.Println("first_partition: ", first_partition)
	fmt.Println("second_partition: ", second_partition)
	fmt.Println("mid_partition: ", mid_partition)
	if len(arrayTosort_tmp[:tail_index+1]) > 1 && len(arrayTosort_tmp[tail_index+2:]) > 1 {
		first_partition = quick_sort(arrayTosort_tmp[:tail_index+1])
		second_partition = quick_sort(arrayTosort_tmp[tail_index+2:])
	} else if len(arrayTosort_tmp[:tail_index+1]) > 1 && len(arrayTosort_tmp[tail_index+2:]) <= 1 {
		first_partition = quick_sort(arrayTosort_tmp[:tail_index+1])
		second_partition = arrayTosort_tmp[tail_index+2:]
	} else if len(arrayTosort_tmp[:tail_index+1]) <= 1 && len(arrayTosort_tmp[tail_index+2:]) > 1 {
		first_partition = arrayTosort_tmp[:tail_index+1]
		second_partition = quick_sort(arrayTosort_tmp[tail_index+2:])
	}
	get_array = merge(first_partition, mid_partition, second_partition)
	return
}

func merge(array_partition1 []int, array_partition2 []int, array_partition3 []int) (final_array []int) {
	array1_index := 0
	array2_index := 0
	array3_index := 0
	main_index := 0
	init_array := make([]int, len(array_partition1)+len(array_partition2)+len(array_partition3))

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

	for array3_index < len(array_partition3) {
		init_array[main_index] = array_partition3[array3_index]
		array3_index++
		main_index++
	}

	final_array = init_array

	return

}

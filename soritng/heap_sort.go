package main

import (
	"fmt"
)

func main() {
	fmt.Println("Heap Sort")
	arrayTosort_tmp := []int{2, 11, 6, 3, 4, 7, 8, 5}
	// buid_binary_tree(arrayTosort_tmp)
	arr_length := len(arrayTosort_tmp)
	final_array := make([]int, len(arrayTosort_tmp))
	// final_array := build_max_heap(arrayTosort_tmp)
	for arr_length > 0 {
		fmt.Println("Array Loop: ", arrayTosort_tmp[:arr_length])
		build_max_heap(arrayTosort_tmp[:arr_length])
		temp_value := arrayTosort_tmp[0]
		arrayTosort_tmp[0] = arrayTosort_tmp[arr_length-1]
		arrayTosort_tmp[arr_length-1] = temp_value
		final_array[len(arrayTosort_tmp)-arr_length] = arrayTosort_tmp[arr_length-1]

		arr_length--
	}
	// fmt.Println(final_array)
	fmt.Println("Final Array: ", final_array)
}

func buid_binary_tree(arrayTosort_tmp1 []int) (get_array []int) {
	fmt.Println(arrayTosort_tmp1)
	i := 0
	j := i + 1

	for i <= (len(arrayTosort_tmp1)-1)/2 {
		fmt.Println("Parent Node: ", arrayTosort_tmp1[i])
		fmt.Println((len(arrayTosort_tmp1) - 1) / 2)
		fmt.Println("value of I and J and array_length: ", i, j, len(arrayTosort_tmp1)-1)
		k := j
		for j <= len(arrayTosort_tmp1)-1 && j <= k+1 {
			fmt.Println("child Node: ", arrayTosort_tmp1[j])
			j++
		}

		i++
	}
	get_array = arrayTosort_tmp1
	return
}

func build_max_heap(arrayTosort_tmp2 []int) (get_array []int) {
	fmt.Println((len(arrayTosort_tmp2) - 1) / 2)
	start_idx := (len(arrayTosort_tmp2) - 1) / 2

	for start_idx >= 0 {
		fmt.Println("Parent Node: ", arrayTosort_tmp2[start_idx])
		heapify(arrayTosort_tmp2, start_idx)

		start_idx--
	}
	// fmt.Println(arrayTosort_tmp2[8])

	get_array = arrayTosort_tmp2
	return
}

func heapify(arrayTosort_tmp2 []int, index_val int) (get_array []int) {
	fmt.Println("Heapify:")
	left_node := (2 * index_val) + 1
	fmt.Println("Left Node: ", left_node)
	right_node := (2 * index_val) + 2
	fmt.Println("Right Node: ", right_node)
	large_index := index_val
	// fmt.Println("Large Index: ", large_index)
	if left_node < len(arrayTosort_tmp2) && arrayTosort_tmp2[large_index] < arrayTosort_tmp2[left_node] {
		large_index = left_node
		// fmt.Println("Left is Large: ", large_index)
	}
	if right_node < len(arrayTosort_tmp2) && arrayTosort_tmp2[large_index] < arrayTosort_tmp2[right_node] {
		large_index = right_node
		// fmt.Println("Right is Large: ", large_index)
	}

	if large_index != index_val {
		fmt.Println("Not Same: ", large_index)
		fmt.Println("Swap Betwee: ", arrayTosort_tmp2[index_val], " and ", arrayTosort_tmp2[large_index])
		temp_val := arrayTosort_tmp2[index_val]
		arrayTosort_tmp2[index_val] = arrayTosort_tmp2[large_index]
		arrayTosort_tmp2[large_index] = temp_val
		fmt.Println("After Swap: ", arrayTosort_tmp2)
		if (2*large_index)+1 <= len(arrayTosort_tmp2)-1 {
			heapify(arrayTosort_tmp2, large_index)
		}

	}

	get_array = arrayTosort_tmp2
	return

}

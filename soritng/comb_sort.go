package main

import (
	"fmt"
)

func main() {
	fmt.Println("Comb Sort")
	arrayTosort_tmp := []int{2, 11, 6, 3, 9, 2, 7, 1, 8, 4, 6, 5, 55, 78, 43, 21}
	// arrayTosort_tmp := []int{2, 11, 2, 4}
	// i := 0
	array_len := float32(len(arrayTosort_tmp) / 2)
	fmt.Println(array_len)
	// i := array_len

	for array_len >= 1 {
		fmt.Println("First Gap Value is: ", array_len)

		fmt.Println("after division: ", array_len)
		i := int(array_len)
		for i < len(arrayTosort_tmp) {

			cycle_and_replace(arrayTosort_tmp, int(array_len), int(i))
			i++
		}
		// fmt.Println("Final Array: ", arrayTosort_tmp)
		array_len = array_len / float32(1.3)
	}

}

func cycle_and_replace(arrayTosort_tmp []int, array_len_in int, index_in int) {

	// get_val := arrayTosort_tmp[index_val_in]

	if array_len_in < len(arrayTosort_tmp) && arrayTosort_tmp[index_in] < arrayTosort_tmp[index_in-array_len_in] {
		fmt.Println(arrayTosort_tmp[index_in], " is shorter than ", arrayTosort_tmp[index_in-array_len_in])
		temp_val := arrayTosort_tmp[index_in]
		arrayTosort_tmp[index_in] = arrayTosort_tmp[index_in-array_len_in]
		arrayTosort_tmp[index_in-array_len_in] = temp_val
		fmt.Println("Final Array: ", arrayTosort_tmp)
		index_in = index_in - array_len_in
		// array_len_in = index_in - array_len_in
		fmt.Println(index_in)
		fmt.Println(array_len_in)
		// cycle_and_replace(arrayTosort_tmp, array_len_in, index_in)

	}
	// else {
	// fmt.Println(arrayTosort_tmp[index_in], " is greater or equal to ", arrayTosort_tmp[index_in-array_len_in])
	// fmt.Println("Not sorted: ", arrayTosort_tmp)
	// }
}

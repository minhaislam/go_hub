package main

import "fmt"

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

		fmt.Println("after division: ", int(array_len))
		i := int(array_len)
		for i < len(arrayTosort_tmp) {

			// cycle_and_replace(arrayTosort_tmp, int(array_len), int(i))
			i++
		}
		// fmt.Println("Final Array: ", arrayTosort_tmp)
		array_len = array_len / float32(1.3)
	}
}

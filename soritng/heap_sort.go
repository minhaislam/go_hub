package main

import (
	"fmt"
)

func main() {
	fmt.Println("Heap Sort")
	arrayTosort_tmp := []int{2, 11, 6, 3, 4, 7, 8, 5}
	build_max_heap(arrayTosort_tmp)

}

func build_max_heap(arrayTosort_tmp1 []int) (get_array []int) {
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

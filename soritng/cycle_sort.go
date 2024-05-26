package main

import (
	"fmt"
)

func main() {
	fmt.Println("Heap Sort")
	arrayTosort_tmp := []int{2, 11, 6, 3, 4, 9, 2, 7, 1, 8, 5}
	// arrayTosort_tmp := []int{2, 11, 2, 4}
	i := 0
	// count := 0
	fmt.Println(arrayTosort_tmp)
	for i < len(arrayTosort_tmp) {
		// j := i + 1
		temp_val := arrayTosort_tmp[i]
		fmt.Println(arrayTosort_tmp[i])
		cycle_sort(arrayTosort_tmp, i, temp_val)
		i++
		// break
	}
	fmt.Println("Final Array: ", arrayTosort_tmp)

}

func cycle_sort(arrayTosort_tmp []int, index_val int, temp_val int) {
	j := index_val + 1
	count := 0
	fmt.Println("tem Value is:", temp_val)
	fmt.Println("Array is:", arrayTosort_tmp)
	for j < len(arrayTosort_tmp) {
		if temp_val > arrayTosort_tmp[j] {
			fmt.Println(temp_val, " greater than ", arrayTosort_tmp[j])
			count = count + 1
		}
		j++
	}
	fmt.Println("Count is: ", count, " For value : ", temp_val)
	// i++
	// fmt.Println(arrayTosort_tmp2[i], "is learger than ", count, " digits")
	if count > 0 {
		// temp_val = arrayTosort_tmp[index_val+count]
		count = index_val + count
		temp_val = cycle_and_replace(arrayTosort_tmp, temp_val, count)
		cycle_sort(arrayTosort_tmp, index_val, temp_val)
	} else {
		arrayTosort_tmp[index_val] = temp_val
	}

}

func cycle_and_replace(arrayTosort_tmp []int, temp_val_in int, count int) (get_val int) {

	// get_val := arrayTosort_tmp[index_val_in]
	if arrayTosort_tmp[count] != temp_val_in {
		get_val = arrayTosort_tmp[count]
		fmt.Println("assign val: ", get_val)
		fmt.Println("Before Adding ", temp_val_in)
		arrayTosort_tmp[count] = temp_val_in
		fmt.Println("Status ", arrayTosort_tmp)
	} else if arrayTosort_tmp[count] == temp_val_in {
		count = count + 1
		cycle_and_replace(arrayTosort_tmp, temp_val_in, count)
	}

	return
}

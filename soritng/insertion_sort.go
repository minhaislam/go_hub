package main

import (
	"fmt"
)

func main() {

	var arrayTosort = []int{9, 1, 1, 6, 8, 9, 4, 67}
	fmt.Println("Provided Array: ", arrayTosort)

	for j := 1; j < len(arrayTosort); j++ {

		keyVal := arrayTosort[j]
		index_no := j - 1
		for index_no >= 0 && arrayTosort[index_no] > keyVal {
			fmt.Println("Swapping Position of Value:", arrayTosort[index_no])
			arrayTosort[index_no+1] = arrayTosort[index_no]
			// fmt.Println(arrayTosort[index_no+1])
			index_no--
		}
		fmt.Println("Status After each loop: ", arrayTosort)
		arrayTosort[index_no+1] = keyVal

	}
	fmt.Println(arrayTosort)
}

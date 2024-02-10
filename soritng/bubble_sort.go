package main

import (
	"fmt"
)

func main() {
	// fmt.Println("hello")

	var arrayTosort = []int{9, 1, 6, 8, 4, 3, 7}
	fmt.Println(arrayTosort)
	for i := len(arrayTosort) - 1; i >= 0; i-- {
		fmt.Println(arrayTosort[i])
		for j := 0; j <= i-1; j++ {
			fmt.Println("2nd Loop", arrayTosort[j])
			if arrayTosort[j] > arrayTosort[j+1] {
				swap_value := arrayTosort[j]
				arrayTosort[j] = arrayTosort[j+1]
				arrayTosort[j+1] = swap_value

			}
		}

	}
	fmt.Println("After Swap :", arrayTosort)
}

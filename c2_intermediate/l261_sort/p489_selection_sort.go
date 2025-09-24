package main

import (
	"fmt"
)

func selectionSort(arr []int32) []int32 {
	for i, num := range arr {
		minIndex := i
		for j := i; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], num
	}
	return arr
}

func main() {
	fmt.Println(selectionSort([]int32{2, 12, 5, 10, 9, 8}))
}

package main

import (
	"fmt"
)

func insertionSort(arr []int32) []int32 {
	for i := 1; i < len(arr); i++ {
		for j := 0; j <= i-1; j++ {
			if arr[i] < arr[j] {
				tmp := arr[i]
				for k := i - 1; k >= j; k-- {
					arr[k+1] = arr[k]
				}
				arr[j] = tmp
				break
			}
		}
	}
	return arr
}

func main() {
	fmt.Println(insertionSort([]int32{2, 12, 5, 10, 9, 8}))
}

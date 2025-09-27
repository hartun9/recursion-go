package main

func sumOfArray(arr []int32) int32 {
	if len(arr) == 0 {
		return 0
	}
	return sumOfArrayHelper(arr, 0, len(arr)-1)
}

func sumOfArrayHelper(arr []int32, start int, end int) int32 {
	if start == end {
		return arr[start]
	}

	middle := (start + end) / 2

	return sumOfArrayHelper(arr, start, middle) + sumOfArrayHelper(arr, middle+1, end)
}

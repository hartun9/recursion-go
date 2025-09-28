package main

func findPeak(arr []int32) int32 {
	return findPeakHelper(arr, 0, len(arr)-1)
}

func findPeakHelper(arr []int32, start int, end int) int32 {
	if start == end {
		return arr[start]
	}
	middle := (start + end) / 2
	return 0
}

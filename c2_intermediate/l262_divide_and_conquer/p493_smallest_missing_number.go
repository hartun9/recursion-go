package main

func smallestMissingNumber(arr []int32) int32 {
	return smallestMissingNumberHelper(arr, 0, len(arr)-1)
}

func smallestMissingNumberHelper(arr []int32, start int, end int) int32 {
	if start == end {
		if start != int(arr[start]) {
			return arr[start] - 1
		}
		return arr[start] + 1
	}
	mid := (start + end) / 2
	if mid < int(arr[mid]) {
		return smallestMissingNumberHelper(arr, start, mid)
	} else {
		return smallestMissingNumberHelper(arr, mid+1, end)
	}
}

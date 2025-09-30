package main

func smallestMissingNumber(arr []int32) int32 {
	min, _ := smallestMissingNumberHelper(arr, 0, len(arr)-1)
	return min
}

func smallestMissingNumberHelper(arr []int32, start int, end int) (min int32, isPresent bool) {
	if start == end {
		min = arr[start]
		isPresent = false
		return min, isPresent
	} else if start+1 == end {
		if arr[start]+1 == arr[end] {
			min = arr[end]
			isPresent = false
			return min, isPresent
		} else {
			min = arr[start] + 1
			isPresent = true
			return min, isPresent
		}
	}
	middle := (start + end) / 2
	minLeft, isPresentLeft := smallestMissingNumberHelper(arr, start, middle)

	if isPresentLeft {
		return minLeft, true
	} else {
		if arr[middle]+1 != arr[middle+1] {
			return arr[middle] + 1, true
		}
	}

	minRight, isPresentRight := smallestMissingNumberHelper(arr, middle+1, end)
	if isPresentRight {
		return minRight, true
	}
	return minRight + 1, false
}

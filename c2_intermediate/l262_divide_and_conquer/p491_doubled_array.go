package main

func doubledArray(arr []int32) []int32 {
	return doubledArrayHelper(arr, 0, len(arr)-1)
}

func doubledArrayHelper(arr []int32, start int, end int) []int32 {
	if start == end {
		return []int32{arr[start] * 2}
	}
	middle := (start + end) / 2
	left := doubledArrayHelper(arr, start, middle)
	right := doubledArrayHelper(arr, middle+1, end)
	return append(left, right...)
}

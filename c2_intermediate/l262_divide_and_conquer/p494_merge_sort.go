package main

func mergeSort(arr []int32) []int32 {
	return mergeSortHelper(arr, 0, len(arr)-1)
}

func mergeSortHelper(arr []int32, start int, end int) []int32 {
	if start == end {
		return []int32{arr[start]}
	}

	mid := (start + end) / 2

	left := mergeSortHelper(arr, start, mid)
	right := mergeSortHelper(arr, mid+1, end)

	leftIdx := 0
	rightIdx := 0

	combined := []int32{}

	for leftIdx <= len(left)-1 || rightIdx <= len(right)-1 {
		if leftIdx >= len(left) {
			combined = append(combined, right[rightIdx])
			rightIdx++
			continue
		} else if rightIdx >= len(right) {
			combined = append(combined, left[leftIdx])
			leftIdx++
			continue
		}

		if left[leftIdx] <= right[rightIdx] {
			combined = append(combined, left[leftIdx])
			leftIdx++
			continue
		} else {
			combined = append(combined, right[rightIdx])
			rightIdx++
			continue
		}
	}

	return combined
}

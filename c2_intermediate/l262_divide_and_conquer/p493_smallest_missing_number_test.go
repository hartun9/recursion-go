package main

import "testing"

func TestSmallestMissingNumber(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		want int32
	}{
		{"two missing", []int32{0, 2, 3, 4, 6}, 1},
		{"one missing", []int32{0, 1, 2, 3, 5}, 4},
		{"last missing", []int32{0, 1, 2, 3, 4}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestMissingNumber(tt.arr)
			if got != tt.want {
				t.Errorf("smallestMissingNumber(%v) = %d; want %d", tt.arr, got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		want []int32
	}{
		{"case1", []int32{2, 12, 5, 10, 9, 8}, []int32{2, 5, 8, 9, 10, 12}},
		{"case2", []int32{1, 5, 3, 4, 3, 1, 6}, []int32{1, 1, 3, 3, 4, 5, 6}},
		{"case3", []int32{11, 45, 32, 75, 88, 15, 15}, []int32{11, 15, 15, 32, 45, 75, 88}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := quickSort(tt.arr)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mergeSort(%v) = %d; want %d", tt.arr, got, tt.want)
			}
		})
	}
}

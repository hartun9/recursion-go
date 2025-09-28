package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		want []int32
	}{
		{"normal", []int32{2, 4, 1, 5, 9}, []int32{1, 2, 4, 5, 9}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectionSort(tt.arr)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("selectionSort(%v) = %d; want %d", tt.arr, got, tt.want)
			}
		})
	}
}

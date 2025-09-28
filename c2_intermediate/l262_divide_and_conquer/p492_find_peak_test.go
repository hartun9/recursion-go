package main

import "testing"

func TestFindPeak(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		want int32
	}{
		{"simgle element", []int32{1}, 1},
		{"multiple element", []int32{4, 5, 10, 2, 7, 5}, 10},
		{"with negatives", []int32{1, -5, 2}, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeak(tt.arr)
			if got != tt.want {
				t.Errorf("findPeak(%v) = %d; want %d", tt.arr, got, tt.want)
			}
		})
	}
}

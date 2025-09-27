package main

import "testing"

func TestSumOfArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		want int32
	}{
		{"empty array", []int32{}, 0},
		{"single element", []int32{5}, 5},
		{"multiple elements", []int32{1, 2, 3, 4}, 10},
		{"with negatives", []int32{-1, -2, 3}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumOfArray(tt.arr)
			if got != tt.want {
				t.Errorf("sumOfArray(%v) = %d; want %d", tt.arr, got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDoubledArray(t *testing.T) {
	tests := []struct {
		name string
		arr  []int32
		want []int32
	}{
		{"normal", []int32{1, 2, 3, 4}, []int32{2, 4, 6, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := doubledArray(tt.arr)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("doubledArray(%v) = %d; want %d", tt.arr, got, tt.want)
			}
		})
	}
}

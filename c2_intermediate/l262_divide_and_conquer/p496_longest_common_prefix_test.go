package main

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want string
	}{
		{"case1", []string{"ccc", "cbd", "cbc", "c"}, "c"},
		{"case2", []string{"Recurrence", "Relation", "Recursion", "Restart"}, "Re"},
		{"case3", []string{"autopilot", "autobiography", "automobile", "auto"}, "auto"},
		{"case4", []string{"divide", "and", "conquer"}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.arr)
			if got != tt.want {
				t.Errorf("longestCommonPrefix(%v) = %s; want %s", tt.arr, got, tt.want)
			}
		})
	}
}

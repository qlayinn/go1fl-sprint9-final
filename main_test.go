package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 100} {
		data := generateRandomElements(size)
		wantSize := size
		if wantSize < 0 {
			wantSize = 0
		}
		if len(data) != wantSize {
			t.Errorf("size %d: got length %d, want %d", size, len(data), wantSize)
		}
		for _, value := range data {
			if value <= 0 {
				t.Errorf("expected positive number, got %d", value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"one", []int{7}, 7},
		{"first", []int{9, 2, 3}, 9},
		{"last", []int{1, 3, 9}, 9},
		{"middle", []int{1, 9, 3}, 9},
		{"equal", []int{5, 5, 5}, 5},
		{"negative", []int{-3, -1, -5}, -1},
		{"eight", []int{1, 2, 3, 4, 5, 6, 7, 8}, 8},
		{"remainder", []int{1, 2, 3, 4, 5, 6, 7, 8, 99}, 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum: got %d, want %d", got, tt.want)
			}
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks: got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestMaxChunksMatchesMaximum(t *testing.T) {
	for size := 1; size <= 100; size++ {
		for position := 0; position < size; position++ {
			data := make([]int, size)
			data[position] = 100
			if got := maxChunks(data); got != maximum(data) {
				t.Fatalf("size %d, position %d: got %d, want 100", size, position, got)
			}
		}
	}
}

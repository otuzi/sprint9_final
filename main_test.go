package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name    string
		size    int
		wantLen int
	}{
		{name: "zero size", size: 0, wantLen: 0},
		{name: "negative size", size: -5, wantLen: 0},
		{name: "one element", size: 1, wantLen: 1},
		{name: "many elements", size: 1000, wantLen: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.size)

			if len(got) != tt.wantLen {
				t.Fatalf("generateRandomElements(%d) returned %d elements, want %d", tt.size, len(got), tt.wantLen)
			}

			for i, value := range got {
				if value <= 0 {
					t.Errorf("element %d = %d, want positive number", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "max at start", data: []int{9, 1, 2, 3}, want: 9},
		{name: "max in middle", data: []int{1, 7, 3}, want: 7},
		{name: "max at end", data: []int{1, 2, 3, 10}, want: 10},
		{name: "all equal", data: []int{5, 5, 5, 5}, want: 5},
		{name: "duplicated max", data: []int{3, 8, 8, 1}, want: 8},
		{name: "negative numbers", data: []int{-3, -1, -2}, want: -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "fewer elements than chunks", data: []int{1, 5, 3}, want: 5},
		{name: "exactly CHUNKS elements", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "length not divisible by CHUNKS, max in tail", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100}, want: 100},
		{name: "max in first chunk", data: []int{100, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}, want: 100},
		{name: "all equal", data: []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}, want: 7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("maxChunks(%v) = %d, want %d", tt.data, got, tt.want)
			}
		})
	}
}

// TestMaxChunksMatchesMaximum проверяет, что оба способа дают одинаковый результат
// на случайных данных разной длины, в том числе не кратной CHUNKS.
func TestMaxChunksMatchesMaximum(t *testing.T) {
	for _, size := range []int{1, 7, 8, 9, 15, 16, 17, 1000, 100_003} {
		data := generateRandomElements(size)

		want := maximum(data)
		if got := maxChunks(data); got != want {
			t.Errorf("size %d: maxChunks = %d, maximum = %d", size, got, want)
		}
	}
}

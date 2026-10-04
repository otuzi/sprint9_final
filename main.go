package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// maxRandomValue ограничивает сверху генерируемые значения.
const maxRandomValue = 1_000_000_000

// generateRandomElements возвращает слайс из size случайных положительных чисел.
// Если size <= 0, возвращает пустой слайс.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	elements := make([]int, size)
	for i := range elements {
		// +1 гарантирует, что число строго положительное.
		elements[i] = rand.Intn(maxRandomValue) + 1
	}

	return elements
}

// maximum возвращает максимальный элемент слайса data.
// Для пустого слайса возвращает 0.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxValue := data[0]
	for _, value := range data[1:] {
		if value > maxValue {
			maxValue = value
		}
	}

	return maxValue
}

// maxChunks возвращает максимальный элемент слайса data, разбивая его
// на CHUNKS частей и обрабатывая каждую часть в отдельной горутине.
// Для пустого слайса возвращает 0.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	// Если элементов меньше, чем частей, делить нечего — ищем в один поток.
	if chunkSize == 0 {
		return maximum(data)
	}

	chunkMaximums := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for chunkIndex := 0; chunkIndex < CHUNKS; chunkIndex++ {
		start := chunkIndex * chunkSize
		end := start + chunkSize
		// Остаток от деления длины на CHUNKS отдаём последней части,
		// чтобы не потерять хвост слайса.
		if chunkIndex == CHUNKS-1 {
			end = len(data)
		}

		wg.Add(1)
		go func(chunkIndex int, chunk []int) {
			defer wg.Done()
			chunkMaximums[chunkIndex] = maximum(chunk)
		}(chunkIndex, data[start:end])
	}
	wg.Wait()

	return maximum(chunkMaximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	startTime := time.Now()
	maxValue := maximum(data)
	elapsed := time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	startTime = time.Now()
	maxValue = maxChunks(data)
	elapsed = time.Since(startTime).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", maxValue, elapsed)
}

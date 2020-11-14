package main

import (
	"fmt"
	"math/rand"
)

func main() {
	insertionSort(intSlice(10))
	insertionSort(intSlice(20))
	insertionSort(intSlice(30))
}

func insertionSort(data []int) {
	fmt.Printf("before: %v\n", data)
	for i := 1; i < len(data); i++ {
		j := i
		for 0 < j {
			if data[j] < data[j-1] {
				data[j-1], data[j] = data[j], data[j-1]
			}
			j = j - 1
		}
	}
	fmt.Printf("after: %v\n", data)
}

func intSlice(size int) []int {
	const randSize = 100

	s := make([]int, size, size)
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(randSize)
	}

	return s
}

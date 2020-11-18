package main

import "math/rand"

func main() {
	s := intSlice(10)
	quick(s, 0, len(s)-1)
}

func quick(a []int, i, j int) {
	if i == j {
		return
	}

	p := pivot(a, i, j)
	if p != -1 {
	}

}

func pivot(a []int, i, j int) int {
	k := i + 1
	for k <= j && a[i] == a[k] {
		k++
	}
	if k > j {
		return -1
	}

	if a[i] >= a[k] {
		return 1
	}
}

func intSlice(size int) []int {
	const randSize = 100

	s := make([]int, size, size)
	for i := 0; i < size; i++ {
		s[i] = rand.Intn(randSize)
	}

	return s
}

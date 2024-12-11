package main

import "fmt"

// Define a custom constraint that includes types supporting the > operator
type ordered interface {
	~int | ~float64 | ~string
}

// Define a generic function that finds the maximum value in a slice
func Max[T ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	strings := []string{"a", "b", "c", "d", "e"}

	fmt.Println("Max int:", Max(ints))
	fmt.Println("Max float:", Max(floats))
	fmt.Println("Max string:", Max(strings))
}

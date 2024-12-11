package main

import "fmt"

func main() {
	s1 := []int{7, 8, 9}
	fmt.Printf("%v,%v,%v\n", s1, len(s1), cap(s1))

	s2 := make([]int, 3, 4)
	fmt.Printf("%v,%v,%v\n", s2, len(s2), cap(s2))

	s2 = append(s2, 1)
	fmt.Printf("%v,%v,%v\n", s2, len(s2), cap(s2))

	s2 = append(s2, 1)
	fmt.Printf("%v,%v,%v\n", s2, len(s2), cap(s2))

	s3 := make([]int, len(s2))
	copy(s3, s2)
	fmt.Printf("%p,%v,%v\n", s2, len(s2), cap(s2))
	fmt.Printf("%v,%v,%v\n", s3, len(s3), cap(s3))
	fmt.Printf("%p,%v,%v\n", s3, len(s3), cap(s3))

	var s4 = s3
	fmt.Printf("%p,%v,%v\n", s4, len(s4), cap(s4))
}

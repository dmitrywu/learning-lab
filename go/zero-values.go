package main

import "fmt"

func main() {
	var name string
	var age int
	var height float64
	var isStudent bool

	fmt.Printf("name: %q\n", name)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("height: %.1f\n", height)
	fmt.Printf("isStudent: %t\n", isStudent)
}
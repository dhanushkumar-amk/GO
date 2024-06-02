package main

import (
	"fmt"
)

// Function to add two integers
func add(x int, y int) int {
	return x + y
}

func main() {
	// Variable declaration and initialization
	var a int = 5
	b := 10 // Short variable declaration

	// Calling the add function
	sum := add(a, b)
	fmt.Println("Sum:", sum)

	// Loop example: printing numbers from 1 to 5
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Conditional example: checking if a number is even or odd
	if sum%2 == 0 {
		fmt.Println("The sum is even")
	} else {
		fmt.Println("The sum is odd")
	}

	// Array example
	arr := [5]int{1, 2, 3, 4, 5}
	for index, value := range arr {
		fmt.Printf("arr[%d] = %d\n", index, value)
	}

	// Map example
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
	}
	for name, age := range ages {
		fmt.Printf("%s is %d years old\n", name, age)
	}
}

package main

import "fmt"

func slicemanip() {
	values := []int{1, 2, 3};
	fmt.Println("Values:", values, "\nCapacity:", cap(values));

	values = append(values, 4);
	fmt.Println("Values:", values, "\nCapacity:", cap(values));
	values = append(values, 5);
	fmt.Println("Values:", values, "\nCapacity:", cap(values));
	values = append(values, 6);
	fmt.Println("Values:", values, "\nCapacity:", cap(values));

	// amortized capacity manipulation, like an ArrayList data structure
}
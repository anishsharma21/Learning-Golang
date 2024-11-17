package main

import "fmt"

func newmake() {
	numP := new(int);
	*numP = 3;
	fmt.Println("Pointer to new int:", numP, "| Value of new int:", *numP);

	values := make([]int, 0, 4);
	for i := 0; i < 5; i++ {
		values = append(values, i+1);
		fmt.Println("Values:", values);
		fmt.Printf("Length of values: %d, Capacity of values: %d\n", len(values), cap(values));
	}
	fmt.Println("Values:", values);

	pInt := new(int);
	*pInt = 4;
	fmt.Printf("Address of int: %p\tValue of int: %d\n", pInt, *pInt);
}
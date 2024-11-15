package main

import "fmt"

func slicecopy() {
	slice1 := []int{1, 2, 3};
	slice2 := []int{3, 4, 5};

	fmt.Println("Slice1:", slice1);
	fmt.Println("Slice2:", slice2);

	copy(slice2, slice1);
	slice1[0] = 3;

	fmt.Println("Slice1:", slice1);
	fmt.Println("Slice2:", slice2);
}
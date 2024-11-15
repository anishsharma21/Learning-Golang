package main

import "fmt"

func dangerousop() {
	fmt.Println("Calling dangerous operation...");
	dangerousOperation();
	fmt.Println("Program recovers and continues to proceed.");
}

func dangerousOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from error:", r);
		}
	}();
	fmt.Println("Starting dangerous operation...");
	panic("Something went wrong, panicking!");
	// fmt.Println("This line will not execute.");
}
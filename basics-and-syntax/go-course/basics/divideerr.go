package main

import (
	"errors"
	"fmt"
)

func divideerr() {
	result, err := divide(5, 3);
	if err != nil {
		fmt.Println("Error:", err);
	} else {
		fmt.Printf("Result of division is %.2f\n", result);
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero");
	}
	return a / b, nil;
}

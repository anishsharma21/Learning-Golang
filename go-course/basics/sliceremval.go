package main

import (
	"errors"
	"fmt"
)

func sliceremval() {
	var values []float64 = []float64{1, 2, 3};
	var valueToRemove float64;

	fmt.Println("Initial array:", values);
	fmt.Println();

	fmt.Print("Value to remove: ");
	fmt.Scan(&valueToRemove);

	values, err := arrRemove(valueToRemove, values);
	if err != nil {
		fmt.Println("Error:", err);
		return;
	}
	fmt.Println("Updated array:", values);
}

func arrRemove(target float64, values []float64) ([]float64, error) {
	for i := 0; i < len(values); i++ {
		if values[i] == target {
			values = append(values[:i], values[i+1:]...);
			return values, nil;
		}
	}
	return values, errors.New("target value not found in array");
}
package main

import (
	"fmt"
)

func averagecalc() {
	values := []float64{1, 2, 3}; // better to use fixed array, rather than slice though
	valuesArr := [4]float64{1, 2, 3, 0};
	var userValue float64;

	fmt.Println("Current values slice:", values);
	fmt.Println("Current values arr:", valuesArr);
	fmt.Print("Number to add to array/slice: ");
	fmt.Scan(&userValue);

	values = append(values, userValue);
	valuesArr[len(values) - 1] = userValue;

	fmt.Println("Updated slice:", values);
	fmt.Println("Updated array:", valuesArr);
	fmt.Println("Calculating average...");
	average := arrAverage(values);
	averageArr := arrAverage(values);
	fmt.Println("Average of values in slice:", average);
	fmt.Println("Average of values in array:", averageArr);
}

func arrAverage(values []float64) float64 {
	var sum float64;
	for i := 0; i < len(values); i++ {
		sum += values[i];
	}
	return float64(sum) / float64(len(values));
}
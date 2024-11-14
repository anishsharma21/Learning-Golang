package main

import "fmt"

func calculator() {
	var a, b int;
	var operator string;

	fmt.Print("First number: ");
	fmt.Scan(&a);
	fmt.Print("Second number: ");
	fmt.Scan(&b);
	fmt.Print("Operation (+ - * /): ");
	fmt.Scan(&operator);

	var result int;
	switch operator {
	case "+":
		result = a + b;
	case "-":
		result = a - b;
	case "*":
		result = a * b;
	case "/":
		if (b != 0) {
			result = a / b;
		}
	default:
		result = 0;
	}
	fmt.Println("Result:", result);
}
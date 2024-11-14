package main

import "fmt"

func sumfirstn() {
	var n int;
	fmt.Print("Number: ");
	fmt.Scan(&n);

	var sum int = 0;
	for i := n; i > 0; i-- {
		sum += i;
	}

	fmt.Println("Sum of first", n, "numbers:", sum);
}
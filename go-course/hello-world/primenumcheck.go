package main

import (
	"fmt"
	"math"
)

func main() {
	var num int;

	fmt.Print("Number to check: ");
	fmt.Scan(&num);

	var isPrime bool = true;

	if (num == 2) {
		isPrime = true;
	} else if (num % 2 == 0) {
		isPrime = false;
	} else {
		sqrt := int(math.Sqrt(float64(num)));
		for i := 3; i <= sqrt; i+=2 {
			if num % i == 0 {
				isPrime = false;
				break;
			}
		}
	}

	if isPrime {
		fmt.Println(num, "is prime.");
	} else {
		fmt.Println(num, "is not prime.");
	}
}
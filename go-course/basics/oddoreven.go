package main

import "fmt"

func oddoreven() {
	var num int = 1;

	for num != 0 {
		fmt.Print("Number: ");
		fmt.Scan(&num);
		if num % 2 == 0 {
			fmt.Println(num, "is even.");
		} else {
			fmt.Println(num, "is odd.");
		}
	}
}
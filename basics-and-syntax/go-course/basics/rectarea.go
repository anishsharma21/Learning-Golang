package main

import "fmt"

func rectarea() {
	var h, w int;
	fmt.Print("Width: ");
	fmt.Scan(&w);
	fmt.Print("Height: ");
	fmt.Scan(&h);

	area := rectangleArea(h, w);
	fmt.Printf("Area of rectangle is %d\n", area);
}

func rectangleArea(height int, width int) (area int) {
	return height * width;
}
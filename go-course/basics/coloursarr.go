package main

import "fmt"

func coloursarr() {
	colours := [3]string{"red", "green", "blue"};
	coloursSlice := colours;

	for i := 0; i < len(coloursSlice); i++ {
		fmt.Printf("Colour %d: %s\n", i + 1, coloursSlice[i]);
	}
}
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt");
	if err != nil {
		fmt.Println("Error:", err);
		return;
	}
	fmt.Println("File created!");
	defer file.Close();

	file, err = os.Open("example.txt");
	if err != nil {
		fmt.Println("Error:", err);
		return;
	}
	fmt.Println("File opened!");
	defer file.Close();

	err = os.Remove("example.txt");
	if err != nil {
		fmt.Println("Error:", err);
		return;
	}
	fmt.Println("File removed!");
}
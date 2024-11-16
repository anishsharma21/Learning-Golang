package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "go.mod";
	fileInfo, err := os.Stat(filename);
	if err != nil {
		fmt.Println("Error opening file:", err);
		return;
	}
	fmt.Println("File name:", fileInfo.Name());
	fmt.Println("File size:", fileInfo.Size());
	fmt.Println("Is directory:", fileInfo.IsDir());
}
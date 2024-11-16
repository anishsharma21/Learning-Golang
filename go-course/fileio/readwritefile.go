package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func readwritefile() {
	filename := "example.txt";
	file, err := os.Create(filename);
	if err != nil {
		fmt.Println("Error creating file:", err);
		return;
	}

	_, err = file.WriteString("Hello, this is Go file handling!\n");
	if err != nil {
		fmt.Println("Error writing to file:", err);
	} else {
		fmt.Println("Successfully wrote to file!");
	}
	file.Close();

	file, err = os.Open(filename);
	if err != nil {
		fmt.Println("Error opening file:", err);
		return;
	}

	fmt.Println("Reading file with io:");
	content := make([]byte, 64);
	for {
		n, err := file.Read(content);
		if err == io.EOF {
			break;
		}
		if err != nil {
			fmt.Println("Error reading file:", err);
			return;
		}
		fmt.Print(string(content[:n]));
	}
	fmt.Println();
	file.Close();

	file, err = os.Open(filename);
	if err != nil {
		fmt.Println("Error opening file:", err);
		return;
	}

	fmt.Println("Reading file with bufio");
	scanner := bufio.NewScanner(file);
	for scanner.Scan() {
		fmt.Println(scanner.Text());
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file with bufio:", err);
	}
	file.Close();

	err = os.Remove(filename);
	if err != nil {
		fmt.Println("Error removing file:", err);
	}
}

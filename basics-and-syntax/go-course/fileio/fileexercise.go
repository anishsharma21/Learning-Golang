package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var filename string;
	reader := bufio.NewReader(os.Stdin);

	fmt.Print("Name of file to create: ");
	fmt.Scan(&filename);

	fmt.Print("Line to add to file: ");
	userstring, _ := reader.ReadString('\n');

	file, err := os.Create(filename);
	if err != nil {
		fmt.Println("Error creating file:", err);
		return;
	}
	
	_, err = file.WriteString(userstring);
	if err != nil {
		fmt.Println("Error writing to file:", err);
		return;
	}
	fmt.Println("Successfully wrote to file!");
	file.Close();

	file, err = os.Open(filename);
	if err != nil {
		fmt.Println("Error opening file:", err);
		return;
	}
	scanner := bufio.NewScanner(file);
	fmt.Println("Contents of file:");
	for scanner.Scan() {
		fmt.Println(scanner.Text());
	}
	file.Close();

	// copy contents to a newly created file
	fileCopy, err := os.Create("copy.txt");
	if err != nil {
		fmt.Println("Error creating new copy file:", err);
		return;
	}
	fileCopy.Close();

	fileOriginal, err := os.Open(filename);
	if err != nil {
		fmt.Println("Error opening original file:", err);
		return;
	}
	defer fileOriginal.Close();
	
	fileCopy, err = os.Open("copy.txt");
	if err != nil {
		fmt.Println("Error opening copy file:", err);
		return;
	}
	defer fileCopy.Close();

	scanner = bufio.NewScanner(fileOriginal);
	writer := bufio.NewWriter(fileCopy);

	for scanner.Scan() {
		content := scanner.Text()
		fmt.Println(content)
		_, err := writer.WriteString(content + "\n")
		if err != nil {
			fmt.Println("Error writing to copy file:", err)
			return
		}
	}

	if err := writer.Flush(); err != nil {
	    fmt.Println("Error flushing writer:", err)
	}

	// count words in file
}
package main

import "fmt"

func grading() {
	var grade int;

	fmt.Print("Grade: ");
	fmt.Scan(&grade);

	var gradeSymbol string;
	switch {
	case grade >= 90 && grade <= 100:
		gradeSymbol = "A";
	case grade >= 80:
		gradeSymbol = "B";
	case grade >= 70:
		gradeSymbol = "C";
	case grade >= 60:
		gradeSymbol = "D";
	default:
		gradeSymbol = "F";
	}
	fmt.Println("Grade for", grade, "is", gradeSymbol);
}
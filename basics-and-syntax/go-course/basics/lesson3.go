package main

import "fmt"

func lesson3() {
	const a int = 3;
	const b int = 10;
	add := a + b;
	subtract := a - b;
	multiply := a * b;
	divide := a / b;
	var divideop float32 = float32(b / a);
	rem1 := a % b;
	var rem2 int = b % a;

	fmt.Println(add, subtract, multiply, divide, divideop, rem1, rem2);
	fmt.Println();

	basicio();
}

func basicio() {
	fmt.Print("Hello");
	fmt.Print(", World");

	const age int = 20;
	const name string = "Anish";

	fmt.Printf("My name is %s and I am %d years old.\n\n", name, age);

	var yourname string;
	fmt.Print("What is your name: ");
	fmt.Scan(&yourname);
	fmt.Printf("Hello, %s\n", yourname);
}
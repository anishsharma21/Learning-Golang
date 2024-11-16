package main

import "fmt"

func pointersgoprog() {
	var num int;
	p := &num;
	fmt.Println("Before:", num);
	*p = 5;
	fmt.Println("After:", num);

	var pointer *int;
	fmt.Println("Uninitialised pointer:", pointer);
	pointer = new(int);
	*pointer = 10;
	fmt.Printf("Pointer (address): %p, Dereferenced value: %d\n", pointer, *pointer);

	doubleByPointer(pointer);
	fmt.Printf("Pointer (address): %p, After doubling: %d\n", pointer, *pointer);
}

func doubleByPointer(intP *int) {
	*intP *= 2;
}
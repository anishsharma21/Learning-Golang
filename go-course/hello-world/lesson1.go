package main

import (
	"fmt"
	"math"
)

func lesson1() {
    greet();
    fmt.Println(math.Pi);
    farewell();

    // var is simplest way to declare variables
    // it also includes implicit typing, as is seen with height
    var age int = 21;
    var name string = "John";
    var height = 5.9;
    var height2 float32 = 5.10;

    // following are inferred when I used the := symbol combination
    count := 0;
    message := "hi there!";

    // can also declare multiple variables in one go
    var x, y int = 2, 3;
    newname, newage := "Bob", 34;

    // constants too
    const pi = 3.1415;
    const greeting = "Hello world!";

    fmt.Println(age, name, height, height2, count, message, newname, newage, pi, greeting, x, y);

    // type conversion
    var intHeight = int(height);
    fmt.Println("Int version of height", intHeight);
}

func greet() {
    fmt.Println("Hi there!");
}

func farewell() {
    fmt.Println("Bye now!");
}
package main

import (
	"fmt"
	"math"
)

func main() {
    greet();
    fmt.Println(math.Pi);
    farewell();
}

func greet() {
    fmt.Println("Hi there!");
}

func farewell() {
    fmt.Println("Bye now!");
}
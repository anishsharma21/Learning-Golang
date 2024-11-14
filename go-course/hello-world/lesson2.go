package main

import (
	"fmt"
	"math"
)

func lesson2() {
    var temperature float32 = 36.6;
    const language string = "Go";

    fmt.Println("Temperature:", temperature);
    countstatus();

    var radius int = 7;
    var radiusF float64 = float64(radius);
    var area float64 = math.Pi * math.Pow(radiusF, 2);
    fmt.Println("Area of circle with radius", radius, "is:", area);
}

func countstatus() {
    count := 10;
    status := "active";
    fmt.Println("Count:", count, "\nStatus:", status);
}
package main

import "fmt"

func celsiusToFahr() {
    var cel float64;
    fmt.Print("Temperature in C˚: ");
    fmt.Scan(&cel);

    var fahr float64 = cel * (9.0 / 5.0) + 32;

    fmt.Printf("Temperature in F˚: %.2f\n", fahr);
}
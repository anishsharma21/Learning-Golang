package main

import "fmt"

type Car struct {
	Brand string
	MaxSpeed float64
}

func main() {
	car := Car{
		Brand: "Toyota",
		MaxSpeed: 234,
	};
	displayCarDetails(&car);
	increaseSpeed(&car, 10);
	displayCarDetails(&car);
}

func displayCarDetails(carP *Car) {
	fmt.Printf("Car brand: %s\tCar max speed: %.2f\n", carP.Brand, carP.MaxSpeed);
}

func increaseSpeed(carP *Car, increment float64) {
	carP.MaxSpeed += increment;
}
package main

import (
	"errors"
	"fmt"
)

var balance float64 = 1000;

func withdrawfrombank() {
	var withdrawalAmount float64;
	fmt.Print("Amount to withdraw: ");
	fmt.Scan(&withdrawalAmount);

	cash, err := withdraw(withdrawalAmount);
	if err != nil {
		fmt.Println("Error:", err);
		return;
	}

	fmt.Println("New balance:", balance);
	fmt.Println("Cash taken out:", cash);
}

func withdraw(amount float64) (float64, error) {
	if amount > balance {
		return 0, errors.New("cannot withdraw with negative balance");
	}
	balance -= amount;
	return amount, nil;
}
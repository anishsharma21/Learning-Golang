package main

import "fmt"

func mapstructs() {
	// maps
	scores := make(map[string]int);
	scores["Game 1"] = 5;
	scores["Game 2"] = 12;
	fmt.Println("Scores:", scores);

	ages := map[string]int{
		"John": 22,
		"Anish": 31,
	};
	fmt.Println("Ages:", ages);

	score, exists := scores["Game 1"];
	if exists {
		fmt.Println("Game 1 score:", score);
	} else {
		fmt.Println("Score does not exist.");
	}

	score, exists = scores["Game 2"];
	if exists {
		delete(scores, "Game 2");
	} else {
		fmt.Println("Score does not exist.");
	}

	type Address struct {
		City string
		Zipcode string
	}

	type Person struct {
		Name string
		Age int
		Address Address
	}
	person := Person{Name: "Ash", Age: 20};
	fmt.Println("Person struct:", person);

	personDetails := Person{
		Name: "John",
		Age: 32,
		Address: Address{
			City: "Sydney",
			Zipcode: "1991",
		},
	};
	fmt.Println("Details of another person:", personDetails);
	fmt.Println("City of other person:", personDetails.Address.City);
}
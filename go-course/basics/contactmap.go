package main

import "fmt"

var contacts map[string]string = make(map[string]string);

func contactmap() {
	var userChoice string;

	for {
		fmt.Println("d - Display contacts");
		fmt.Println("a - Add contact");
		fmt.Println("r - Remove contact");
		fmt.Println("f - Find contact");
		fmt.Println("q - Quit");
		fmt.Println();
		fmt.Print("Choose an option from above: ");
		fmt.Scan(&userChoice);
		fmt.Println();

		switch userChoice {
		case "d":
			displayContacts();
		case "a":
			addContact();
		case "r":
			removeContact();
		case "f":
			findContact();
		case "q":
			return;
		default:
			fmt.Println("Invalid user choice. Please select from d, a, r, f, or q.");
		}
		fmt.Println();
	}
}

func displayContacts() {
	for key, value := range contacts {
		fmt.Printf("Name: %s, Phone Number: %s\n", key, value);
	}
}

func addContact() {
	var newContactName string;
	var newContactPhoneNumber string;
	fmt.Print("Name of new contact: ");
	fmt.Scan(&newContactName);
	fmt.Print("Phone number of new contact: ");
	fmt.Scan(&newContactPhoneNumber);

	contacts[newContactName] = newContactPhoneNumber;
	fmt.Println("Successfully added new contact.");
}

func removeContact() {
	var contactNameToRemove string;
	fmt.Print("Name of contact to remove: ");
	fmt.Scan(&contactNameToRemove);

	_, exists := contacts[contactNameToRemove];
	if exists {
		delete(contacts, contactNameToRemove);
		fmt.Println("Name", contactNameToRemove, "successfully deleted.");
	} else {
		fmt.Println("Name", contactNameToRemove, "does not exist in contacts.");
	}
}

func findContact() {
	var contactName string;
	fmt.Print("Name of contact to find: ");
	fmt.Scan(&contactName);

	_, exists := contacts[contactName];
	if exists {
		fmt.Printf("Name: %s, Phone number: %s\n", contactName, contacts[contactName]);
	} else {
		fmt.Println("Name", contactName, "does not exist in contacts.");
	}
}
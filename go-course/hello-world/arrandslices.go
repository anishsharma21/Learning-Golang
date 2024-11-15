package main

import (
	"fmt"
)

func arrandslices() {
	var values [5]int;
	values[0] = 1;
	values[1] = 2;

	for i := 0; i < len(values); i++ {
		fmt.Printf("Value at index %d: %d\n", i, values[i]);
	}

	letters := [3]string{"a", "b", "c"};
	fmt.Println(letters);

	sliceOfValues := letters[0:3];
	fmt.Println(sliceOfValues);

	numbers := []int{1, 2, 3};
	fmt.Println("Numbers slice:", numbers);

	characters := make([]string, 3, 5);
	fmt.Println("Length of characters slice:", len(characters));
	fmt.Println("Capacity of characters slice:", cap(characters));
	fmt.Println("characters slice:", characters);

	arr := [5]int{1, 2, 3, 4, 5};
	sli := arr[1:4];

	fmt.Println("Arr:", arr);
	fmt.Println("Sli:", sli);

	sli = append(sli, 6);
	
	fmt.Println("Arr:", arr);
	fmt.Println("Sli:", sli);
}
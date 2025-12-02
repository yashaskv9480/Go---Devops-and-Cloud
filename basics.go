package main

import (
	"fmt"
)

func add(a int, b int) int { // decalring function
	return a + b
}

func Result() {
	var x int = 10 //int
	y := 20
	const pi = 3.14 //float

	if x < y {
		fmt.Println("x is smaller")
	} else {
		fmt.Println("y is smaller")
	}

	for i := 1; i <= 3; i++ { //for lloop
		fmt.Println("Loop:", i)
	}

	nums := [4]int{1, 2, 3, 4} //array
	fmt.Println("Array:", nums)

	slice := []int{10, 20, 30} //Slice
	fmt.Println("Slice:", slice)

	m := map[string]int{"a": 1, "b": 2} //map
	fmt.Println("Map:", m)

	result := add(5, 7)
	fmt.Println("Add Result:", result)

	fmt.Println("PI:", pi)
}

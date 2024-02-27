package main

import "fmt"

// Cheak a Largest Number

func main() {
	var num1 int
	var num2 int
	var num3 int

	fmt.Println("Enter Three Numbers")
	fmt.Scanf("%d %d %d", &num1, &num2, &num3)

	if num1 > num2 {
		fmt.Printf(" num1 %d is a Largest Number", num1)

	} else if num2 > num3 {
		fmt.Printf("num2 %d is a Largest Number", num2)
	} else if num3 > num1 && num3 > num2 {
		fmt.Printf("num3 %d is a Largest Number", num3)
	}
}

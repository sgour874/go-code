package main

import "fmt"

func main() {
	var num uint64
	fact := uint64(1)
	fmt.Println("Enter Number")
	fmt.Scanf("%d", &num)

	for i := uint64(1); i <= num; i++ {
		fact *= i
	}
	fmt.Printf("Factorial = %d", fact)
}

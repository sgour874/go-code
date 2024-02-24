package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter Number")
	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		for k := 1; k <= i; k++ {
			fmt.Printf("  ")
		}
		for j := num - i; j >= i; j-- {
			fmt.Printf("* ")
		}
		fmt.Printf("\n")
	}

}

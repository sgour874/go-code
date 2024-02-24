// write a programm which print a number only divisible by 7 and not be multiple of 5
package main

import "fmt"

func main() {
	var lowerLimit int
	var upperLimit int

	fmt.Println("Enter Lower & Upper Limits")
	fmt.Scanf("%d %d", &lowerLimit, &upperLimit)

	for i := lowerLimit; i <= upperLimit; i++ {
		if isDivisibleBy7(i) && !isDivisibleBy5(i) {
			fmt.Printf("%d ", i)
		}
	}
}

func isDivisibleBy7(n int) bool {
	return n%7 == 0
}

func isDivisibleBy5(n int) bool {
	return n%5 == 0
}

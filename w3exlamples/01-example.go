package main

// Check a Leap Year

import "fmt"

func main() {
	var year int
	fmt.Println("Enter a Year")
	fmt.Scanf("%d", &year)

	if check_leap_year(year) {
		fmt.Println("It is leap Year")
	} else {
		fmt.Println("It is not leap Year")
	}
}

func check_leap_year(n int) bool {
	if n%4 == 0 || n%100 == 0 || n%400 == 0 {
		return true
	}
	return false
}

package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter number")
	fmt.Scanf("%d", &num)

	if isPrime(num) {
		fmt.Println("Prime Number")
	} else {
		fmt.Println("Not Prime Number")
	}
}
func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true

}

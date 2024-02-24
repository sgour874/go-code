package main

import "fmt"

func main() {
	var num int
	var result []int
	fmt.Println("Enter Number")
	fmt.Scanf("%d", &num)

	for i := 1; i <= num; i++ {
		if isPrime(i) {
			result = append(result, i)
		}
	}

	fmt.Println(result)

}

func isPrime(n int) bool {
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

package main

// Two Interval Prime Number

import "fmt"

func main() {
	var lower int
	var upper int

	fmt.Println("Enter Lower & Upper Value")
	fmt.Scanf("%d %d", &lower, &upper)

	for i := lower; i < upper; i++ {
		if isprimeNum(i) {
			fmt.Printf("%d\n", i)
		}
	}
}

func isprimeNum(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

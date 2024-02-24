package main

import "fmt"

func main() {
	square := make(map[uint]uint)

	var lowerLimit uint
	var upperLimit uint

	fmt.Println("Enter Lower & Upper Limits")
	fmt.Scanf("%d %d", &lowerLimit, &upperLimit)

	for i := lowerLimit; i <= upperLimit; i++ {
		square[i] = i * i
	}
	fmt.Printf("\n\n%v", square)
}

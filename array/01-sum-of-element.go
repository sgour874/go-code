package main

import "fmt"

func main() {

	var a [10]int
	sum := 0

	for i := 0; i < 10; i++ {
		fmt.Scanf("%d", &a[i])

	}
	for i := 0; i < 10; i++ {
		sum += a[i]
	}

	fmt.Printf("%d", sum)

}

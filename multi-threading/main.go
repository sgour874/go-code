package main

import (
	"fmt"
	"sgour/primes"
)

func main() {
	fmt.Println("Enter a range")
	var lowerLimit, upperLimit int
	fmt.Scanf("%d %d", &lowerLimit, &upperLimit)
	p := primes.GetPrimes(lowerLimit, upperLimit)
	p.Run()
	p.Wait()
}

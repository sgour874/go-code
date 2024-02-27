package main

// Make a Simple Calculater

import "fmt"

func main() {
	var a float64
	var b float64

	fmt.Println("Enter a & b ")
	fmt.Scanf("%f %f", &a, &b)

	addition := addition(a, b)
	substraction := substraction(a, b)
	maltipication := maltipication(a, b)
	division := division(a, b)

	fmt.Printf("Addition = %f\n", addition)
	fmt.Printf("Substraction = %f\n", substraction)
	fmt.Printf("Maltipication = %f\n", maltipication)
	fmt.Printf("Division = %f\n", division)
}

func addition(n float64, m float64) float64 {
	return n + m
}

func substraction(n float64, m float64) float64 {
	return n - m
}

func maltipication(n float64, m float64) float64 {
	return n * m
}

func division(n float64, m float64) float64 {
	return n / m
}

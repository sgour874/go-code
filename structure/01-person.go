package main

import "fmt"

type Complex struct {
	real float32
	img  float32
}

func main() {
	var a Complex
	var b Complex
	var c Complex
	fmt.Println("Enter number a, b")
	fmt.Scanf("%f %f %f %f", &a.real, &a.img, &b.real, &b.img)

	c.real = a.real + b.real
	c.img = a.img + b.img

	fmt.Printf(" Ans = %f %fi ", c.real, c.img)

}

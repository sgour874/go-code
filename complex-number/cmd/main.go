package main

import (
	"sgour/complexnumber"
)

func main() {
	a := complexnumber.GetComplex(23, 45)

	/*
		add := complexnumber.Add(a, b)
		mul := complexnumber.Multiply(a, b)
		div := complexnumber.Divisionn(a, b)
		sub := complexnumber.Substract(a, b)

		complexnumber.Print(add)
		complexnumber.Print(mul)
		complexnumber.Print(div)
		complexnumber.Print(sub)
	*/

	b := complexnumber.GetComplex(a.Real()+4, a.Img())

	a.Add(b)
	a.Print()
	b.Print()

}

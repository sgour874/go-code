package main

import "fmt"

type Address struct {
	name    string
	city    string
	state   string
	pincode int
}

func main() {
	a1 := &Address{"Priti", "Raipur", "Chhatisgad", 123456}

	fmt.Println("Address1: ", a1)

	a2 := Address{"Shubham", "pune", "Maharastra", 324354}

	fmt.Println("Address: ", a2)

	a3 := Address{name: "Taani", city: "Amravati", pincode: 878679}

	fmt.Println("Address: ", a3)

	fmt.Println("a1name: ", (*a1).name)

	fmt.Println("a2city: ", a2.city)

	a3.state = "Maharastra"

	fmt.Println("Address3: ", a3)

}

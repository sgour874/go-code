package main

import "fmt"

func main() {
	for i, j := range "xabcdefgKA" {
		fmt.Printf("The Index Number Of %U if %d\n", j, i)
	}
}

package main

import "fmt"

func main() {
	colour := [3]string{"pink", "orange", "yellow"}
	fruite := [4]string{"stawberry", "papaya", "bannana", "apple"}

	for i := 0; i < len(colour); i++ {
		for j := 0; j < len(fruite); j++ {
			fmt.Println(colour[i], fruite[j])
		}
	}
}

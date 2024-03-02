package main

import "fmt"

func main() {
	mmap := map[int]string{
		1: "how",
		2: "are",
		3: "you",
	}
	for x, y := range mmap {
		fmt.Println(x, y)
	}

}

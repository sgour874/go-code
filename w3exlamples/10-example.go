package main

import "fmt"

func main() {
	chanl := make(chan int)
	go func() {
		chanl <- 10
		chanl <- 100
		chanl <- 1000
		chanl <- 10000
		close(chanl)
	}()

	for i := range chanl {
		fmt.Println(i)
	}
}

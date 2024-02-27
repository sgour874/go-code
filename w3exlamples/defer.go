package main

import (
	"fmt"
	"sync"
)

func main() {

	printInParallel()

}

func printInParallel() {
	var wg sync.WaitGroup

	for i := 0; i < 15; i++ {

		wg.Add(1)
		go func() {
			defer wg.Done()
			printNumbertill10()
		}()
	}

	wg.Wait()
	fmt.Println()
	fmt.Println("All thread ended")

}

func printNumbertill10() {

	for i := 1; i <= 100; i++ {
		fmt.Printf("%d ", i)
	}
}

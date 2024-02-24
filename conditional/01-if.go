package main

import "fmt"

func main() {
	var mark int
	fmt.Println("Enter marks")
	fmt.Scanf("%d", &mark)

	if mark > 45 {
		fmt.Println("Result is Pass")
	} else {
		fmt.Println("Result is Fail")
	}
}

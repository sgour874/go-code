package main

// Switch Case

import "fmt"

func main() {
	var day int

	fmt.Println("Enter Number")
	fmt.Scanf("%d", &day)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wensday")
	case 4:
		fmt.Println("Thusday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Week End")
	default:
		fmt.Println(" No More Days")
	}

}

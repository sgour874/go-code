package main

import "fmt"

func main() {
	flowers := map[string]string{
		"rose":     "red",
		"lily":     "pink",
		"lavender": "white",
	}
	for key, value := range flowers {
		fmt.Println(key, value)
	}

}

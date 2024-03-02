package main

import "fmt"

func main() {

	arr := [3][3]string{{"Shero", "Pikachu", "Taddy"},
		{"Sinchan", "Ais", "Mr.Ben"},
		{"Chocochip", "Riseball", "Pop"},
	}
	fmt.Println("Elements Of Arry1")

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println(arr[x][y])
		}
	}

	var arr1 [2][2]int
	arr1[0][0] = 10
	arr1[0][1] = 20
	arr1[1][0] = 30
	arr1[1][1] = 40

	fmt.Println("\nElements of Arry2")

	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			fmt.Println(arr1[x][y])
		}
	}
	// find the length of array
	var arr2 = []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println("Length of Arr2 : ", len(arr2))
}

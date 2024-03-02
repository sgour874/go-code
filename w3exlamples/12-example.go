package main

import "fmt"

type Student struct {
	name   string
	branch string
	Std    int
}

type Teacher struct {
	name      string
	subject   string
	experiece int
	details   Student
}

func main() {
	result := Teacher{
		name:      "Anisha",
		subject:   "Maths",
		experiece: 8,
		details:   Student{"Shivani", "CBSE", 8},
	}
	fmt.Println("Details of Teacher")
	fmt.Println("Teacher Name: ", result.name)
	fmt.Println("Subject: ", result.subject)
	fmt.Println("Experiece: ", result.experiece)

	fmt.Println("\nStudent Name: ", result.details.name)
	fmt.Println("Branch: ", result.details.branch)
	fmt.Println("Standard: ", result.details.Std)
}

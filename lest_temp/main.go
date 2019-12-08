package main

import "fmt"

var (
	marks = 90
	grade = "D"
)

func main() {

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println("Your grade is:", grade)
}

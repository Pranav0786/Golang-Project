package main

import (
	"fmt"
	features "student-db/Features"
)

func Exit() {
	fmt.Println("Exiting the program.")
	// You can perform any cleanup or additional actions before exiting if needed.
	// For now, let's just end the program.
	panic("Program terminated.")
}

func displayChoices() {
	fmt.Println("Select one of the choices")
	fmt.Println("1. Add a new Student")
	fmt.Println("2. Show all Students")
	fmt.Println("3. Update Student")
	fmt.Println("4. Delete Student")
	fmt.Println("5. Get Start")
	fmt.Println("6. Exit")
}

func main() {
	fmt.Println("Welcome to student DB")
	var class features.Class
	class.NewClass()

	for {
		displayChoices()
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			class.AddStudent()
		case 2:
			class.ShowStudents()
		case 3:
			class.UpdateStudent()
		case 4:
			class.DeleteStudent()
		case 5:
			class.GetStart()
		case 6:
			Exit()
		default:
			fmt.Println("Invalid choice")
		}
	}

}

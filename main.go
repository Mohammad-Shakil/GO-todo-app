package main

import (
	"fmt"
)

func main() {

	fmt.Println("---TO DO---")

	for {
		var option string

		fmt.Println("\nAdd task--->1")
		fmt.Println("Show task--->2")
		fmt.Println("Delete task--->3")
		fmt.Println("Completed task--->4")
		fmt.Println("Enter Q to exit")
		fmt.Print("Choose:")
		fmt.Scanln(&option)

		switch option {
		case "1", "2", "3", "4":
		case "Q", "q":
			return
		default:
			fmt.Println("\nInvalid operation")
			continue
		}

	}

}

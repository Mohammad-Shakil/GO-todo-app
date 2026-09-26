package logics

import "fmt"

type Task struct {
	Task      string
	Completed bool
}

func Check(option string) {
	switch option {
	case "1", "2", "3", "4":
	case "Q", "q":
		return
	default:
		fmt.Println("\nInvalid operation")
		continue
	}
}

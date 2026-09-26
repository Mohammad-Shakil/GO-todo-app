package main

import (
	"fmt"
	"todo/logics"
)

func main() {
	var Tasks []logics.Work

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
		case "1":
			var task string
			fmt.Print("\n Enter task:")
			fmt.Scanln(&task)
			Tasks = append(Tasks, logics.Work{Work: task, Completed: false})
		case "2":
			fmt.Println()
			for i, t := range Tasks {
				fmt.Printf("%d. %s (Done: %t)\n", i+1, t.Work, t.Completed)
			}
		case "3":
			var deltask string
			fmt.Print("Enter task to delete:")
			fmt.Scanln(&deltask)
			index := -1
			for i := 0; i < len(Tasks); i++ {
				if deltask == Tasks[i].Work {
					index = i
				}
			}

			if index == -1 {
				fmt.Println("NO task found")
				continue
			}

			for i := index; i < len(Tasks)-1; i++ {

				Tasks[i] = Tasks[i+1]

			}
			Tasks = Tasks[:len(Tasks)-1]
			fmt.Println("Success task deleted")

		case "Q", "q":
			return

		}

	}

}

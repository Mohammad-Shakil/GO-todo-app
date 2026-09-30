package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"todo/logics"
)

func main() {
	var Tasks []logics.Work
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("---TO DO---")

	for {
		var option string

		fmt.Println("\nAdd task--->1")
		fmt.Println("Show task--->2")
		fmt.Println("Delete task--->3")
		fmt.Println("Completed task--->4")
		fmt.Println("Edit task-->5")
		fmt.Println("Search task keyword--->6")
		fmt.Println(" Q to exit")
		fmt.Print("Choose:")
		fmt.Scanln(&option)

		switch option {
		case "1":

			fmt.Print("\n Enter task:")
			task, _ := reader.ReadString('\n')
			task = strings.TrimSpace(task)

			Tasks = append(Tasks, logics.Work{Work: task, Completed: false})

		case "2":
			fmt.Println()
			fmt.Println("Total tasks:", len(Tasks))
			Completed := 0
			Pending := 0
			for i := 0; i < len(Tasks); i++ {
				if Tasks[i].Completed == true {
					Completed++
				} else {
					Pending++
				}
			}
			fmt.Println("Completed:", Completed)
			fmt.Println("Pending:", Pending)
			for i, t := range Tasks {
				fmt.Printf("%d. %s (Done: %t)\n", i+1, t.Work, t.Completed)
			}
		case "3":

			fmt.Print("Enter task to delete:")
			deltask, _ := reader.ReadString('\n')
			deltask = strings.TrimSpace(deltask)

			index := -1
			for i := 0; i < len(Tasks); i++ {
				if deltask == Tasks[i].Work {
					index = i
				}
			}

			if index == -1 {
				err := errors.New("NO item found")
				if err != nil {
					fmt.Println("Item not found")
				}
				continue
			}

			for i := index; i < len(Tasks)-1; i++ {

				Tasks[i] = Tasks[i+1]

			}
			Tasks = Tasks[:len(Tasks)-1]
			fmt.Println("Success task deleted")

		case "Q", "q":
			return
		case "5":
			var oldTask string
			fmt.Print("\nEnter task name to update:")
			fmt.Scanln(&oldTask)

			index := -1
			for i := 0; i < len(Tasks); i++ {
				if Tasks[i].Work == oldTask {
					index = i
					break
				}

			}

			if index == -1 {
				fmt.Println("No task found")
				continue
			}
			var newtask string
			fmt.Print("Enter new name:")
			fmt.Scanln(&newtask)

			Tasks[index].Work = newtask
			fmt.Println("Task updated")

		case "4":
			var op string
			fmt.Print("\nEnter completed task name:")
			fmt.Scanln(&op)
			index := -1
			for i := 0; i < len(Tasks); i++ {
				if op == Tasks[i].Work {
					index = i
					break
				}
			}
			if index == -1 {
				fmt.Println("\nNo task found")
				continue
			}

			Tasks[index].Completed = true
			fmt.Println("\nDONE")

		case "6":
			var keyWord string
			fmt.Printf("\nEnter task: ")
			fmt.Scanln(&keyWord)
			found := false
			for i := 0; i < len(Tasks); i++ {

				if strings.Contains(Tasks[i].Work, keyWord) {
					fmt.Println(Tasks[i].Work)
					found = true
				}

			}
			if found == false {
				fmt.Println("NO task found")
			}

		default:
			fmt.Println("\n Invalid operation")
		}

	}

}

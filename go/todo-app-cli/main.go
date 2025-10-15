package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct { //define object using struct
	ID    int
	Title string
	Done  bool
}

var tasks []Task // in memory storage
var nextId int = 1

/*
scanner is a bufio.Scanner object, which reads input from a source (in this case, os.Stdin which is the terminal input).

scanner.Scan()
This method waits for and reads the next line of input from the user. It returns true if it successfully read input, or false if there was an error or end of input.

scanner.Text()
After calling Scan(), this method retrieves the text of the line that was just read as a string.
*/

func main() {

	loadTask() //task saved in the file

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nVChoose following options for TODO")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Completed task")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit()")
		scanner.Scan()           //wait for scan the user input text
		choice := scanner.Text() //get the user input value
		switch choice {
		case "1":
			fmt.Println("Enter the Task Title: ")
			scanner.Scan()
			title := scanner.Text()
			addTask(title)
		case "2":
			fmt.Println("View the Task")
			viewListTask()
		case "3":
			fmt.Println("Enter the task id to complete the task: ")
		case "4":
			fmt.Println("Enter the task id that you want to delete: ")
		case "5":
			fmt.Println("Thank you! For Choosing ToDo ApP")
			saveAllTasks()
			return
		default:
			fmt.Println("Invalid Input")
			return
		}

	}

}

func addTask(title string) {
	task := Task{ID: nextId, Title: title, Done: false}
	tasks = append(tasks, task)
	nextId++
	fmt.Println("Task Added")

}

func viewListTask() {
	fmt.Println("Task List:")
	for _, v := range tasks {
		status := ""
		if v.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", v.ID, status, v.Title)
	}
}

func saveAllTasks() {
	file, err := os.Create("task.json")
	if err != nil {
		fmt.Println("Error! Unable to save new task", err)
		return
	}
	defer file.Close()
	json.NewEncoder(file).Encode(tasks)

}

func loadTask() {
	file, err := os.Open("task.json")
	if err != nil {
		return //no task saved
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&tasks)

	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	nextId = maxID + 1

}

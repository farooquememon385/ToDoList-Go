package main

import (
	"fmt"
	"bufio"
	"os"
)

//the to do list is implemented with file handling
const todoFile = "todo.txt"

func main(){
	fmt.Println("#######################################")
	fmt.Println("******** Welcome to To-Do List ********")
	fmt.Println("#######################################")


	// Read user input and perform corresponding actions
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a command (add/remove/show/quit): ")
	scanner.Scan()
	command := scanner.Text()


}

// Load existing to-do items from file
func loadTodos() []string {
	var todos []string

	file, err := os.Open(todoFileName)
	if err != nil {
		fmt.Println("No existing to-do list found. Starting with an empty list.")
		return todos
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		todo := scanner.Text()
		todos = append(todos, todo)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading to-do list:", err)
	}

	fmt.Println("To-do list loaded from file.")
	return todos
}
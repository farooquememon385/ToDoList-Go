package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

//the to do list is implemented with file handling
const todoFileName = "todo.txt"

func main(){
	fmt.Println("#######################################")
	fmt.Println("******** Welcome to To-Do List ********")
	fmt.Println("#######################################")

	// Load existing to-do items from file
	todos := loadTodos()

	// Display the current to-do list
	fmt.Println("Current To-Do List:")
	displayTodos(todos)


	// Read user input and perform corresponding actions
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter a command (add/remove/show/quit): ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "add":
			fmt.Print("Enter the to-do item: ")
			scanner.Scan()
			item := scanner.Text()
			todos = append(todos, item)
			fmt.Printf("Added '%s' to the to-do list.\n", item)
		case "remove":
			fmt.Print("Enter the index of the item to remove: ")
			scanner.Scan()
			indexStr := scanner.Text()
			index := getIndex(indexStr)

			if index >= 0 && index < len(todos) {
				removedItem := todos[index]
				todos = removeItem(todos, index)
				fmt.Printf("Removed '%s' from the to-do list.\n", removedItem)
			} else {
				fmt.Println("Invalid index.")
			}
		case "show":
			// Display the current to-do list
			fmt.Println("Current To-Do List:")
			displayTodos(todos)
		case "quit":
			// Save the to-do items to file before quitting
			saveTodos(todos)
			fmt.Println("To-do list saved. Exiting...")
			return
		default:
			fmt.Println("Invalid command.")
		}

		// Display the updated to-do list
		displayTodos(todos)
	}

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

func saveTodos(todos []string) {
	file, err := os.Create(todoFileName)
	if err != nil {
		fmt.Println("Error saving to-do list:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, todo := range todos {
		_, err := writer.WriteString(todo + "\n")
		if err != nil {
			fmt.Println("Error saving to-do list:", err)
			return
		}
	}

	writer.Flush()
	fmt.Println("To-do list saved to file.")
}

// Display the current to-do list
func displayTodos(todos []string) {
	fmt.Println("--------")
	for i, todo := range todos {
		fmt.Printf("%d. %s\n", i, todo)
	}
	fmt.Println("--------")
}

// Remove an item from the to-do list
func removeItem(todos []string, index int) []string {
	return append(todos[:index], todos[index+1:]...)
}

// Get the index from a string input
func getIndex(indexStr string) int {
	index, err := strconv.Atoi(indexStr)
	if err != nil {
		return -1
	}
	return index
}

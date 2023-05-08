package main

import (
	"fmt"
	"bufio"
	"os"
)

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
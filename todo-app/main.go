package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	Title       string
	Completed   bool
	CompletedAt *time.Time
	CreatedAt   time.Time
}

// Create a slice
type Todos []Todo

// ➕ Add todos
func (todos *Todos) add(title string) Todo {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
	return todo
}

func main() {
	fmt.Println("Todo App 🚀")
	fmt.Println("Type 'add' to add a Todo")
	fmt.Println("Type 'remove' to remove a Todo")

	var todo Todos
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		command := scanner.Text()

		switch command {
		case "add":
			fmt.Println("Add a new Task")

			if scanner.Scan() {
				title := scanner.Text()
				newTodo := todo.add(title)
				fmt.Printf("Added new task: %s\n", newTodo.Title)
			}
		case "remove":
			fmt.Println("Your task is deleting...")
		default:
			fmt.Println("❓ Unknown command. Try 'add' or 'remove'.")

		}

	}

}

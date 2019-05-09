package main

import "fmt"

var currentID int

var todos Todos

// init provieds seed data.
func init() {
	createTodo(Todo{Name: "Write presentation"})
	createTodo(Todo{Name: "Host meetup"})
}

func findTodo(id int) Todo {
	for _, t := range todos {
		if t.ID == id {
			return t
		}
	}

	return Todo{}
}

func createTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todos = append(todos, t)
	return t
}

func deleteTodo(id int) error {
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with `id` of %d to delete", id)
}

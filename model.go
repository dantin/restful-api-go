package main

import "time"

// Todo servers as a mock entity.
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos is a collections of `Todo`.
type Todos []Todo

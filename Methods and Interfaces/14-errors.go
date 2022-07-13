package main

import (
	"fmt"
	"time"
)

type MyCustomError struct {
	When time.Time
	What string
}

func (e *MyCustomError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyCustomError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

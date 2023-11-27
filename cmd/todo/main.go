package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/christsantiris/cli-todo"
)

const (
	todoFile = ".todos.json"
)

func main() {
	// Pass the flags to the cli as arguments. See readme.
	add := flag.Bool("add", false, "Add a new To do")
	complete := flag.Int("complete", 0, "Mark a to do as completed")
	delete := flag.Int("delete", 0, "delete a to do")

	flag.Parse()

	todos := &todo.Todos{}

	// Check if file exists. If not, generate it with empty add command. If exists, grab it
	if err := todos.LoadItems(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		todos.AddItem("Sample to do")
		err := todos.StoreAddedItem(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *complete > 0:
		err := todos.CompleteItem(*complete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.StoreAddedItem(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	case *delete > 0:
		err := todos.DeleteItem(*delete)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.StoreAddedItem(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}
}

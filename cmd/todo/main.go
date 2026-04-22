package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	todo "github.com/christsantiris/cli-todo"
)

func resolveFilePath() (string, error) {
	if path := os.Getenv("TODO_FILE"); path != "" {
		return path, nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + "/.todos.json", nil
}

func main() {
	// Pass the flags to the cli as arguments. See readme.
	add := flag.Bool("add", false, "Add a new To do")
	complete := flag.Int("complete", 0, "Mark a to do as completed")
	delete := flag.Int("delete", 0, "delete a to do")
	edit := flag.Int("edit", 0, "Edit a to do by index")
	list := flag.Bool("list", false, "list all todos")

	flag.Parse()

	todoFile, err := resolveFilePath()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	todos := &todo.Todos{}

	// Check if file exists. If not, generate it with empty add command. If exists, grab it
	if err := todos.LoadItems(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.AddItem(task)
		err = todos.StoreAddedItem(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.PrintToDos("")
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
		todos.PrintToDos("")
	case *edit > 0:
		task, err := getInput(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.EditItem(*edit, task)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		err = todos.StoreAddedItem(todoFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		todos.PrintToDos("")
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
		todos.PrintToDos("")
	case *list:
		todos.PrintToDos(strings.Join(flag.Args(), " "))
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(1)
	}
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("a to do value is required")
	}

	return scanner.Text(), nil
}

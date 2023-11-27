package todo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type Item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []Item

func (t *Todos) AddItem(task string) {
	todo := Item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo) // Append newly created todo item to the list of To Dos
}

func (t *Todos) CompleteItem(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	ls[index-1].CompletedAt = time.Now()
	ls[index-1].Done = true

	return nil
}

func (t *Todos) DeleteItem(index int) error {
	ls := *t
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}

	*t = append(ls[:index-1], ls[index:]...) // Remove the item at this index from list

	return nil
}

func (t *Todos) LoadItems(filename string) error { // Load items from json file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}

	return nil
}

func (t *Todos) StoreAddedItem(filename string) error { // Write added todo back to json file
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, data, 0644)
}

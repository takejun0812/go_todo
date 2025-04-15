package todo

import (
	"encoding/json"
	"os"
	"time"
)

type Todo struct {
	Task string    `json:"task"`
	Done bool      `json:"done"`
	Due  time.Time `json:"due"`
}

const todoFile = "todos.json"

func Load() ([]Todo, error) {
	var todos []Todo
	data, err := os.ReadFile(todoFile)
	if err != nil {
		if os.IsNotExist(err) {
			return todos, nil
		}
		return nil, err
	}
	err = json.Unmarshal(data, &todos)
	return todos, err
}

func Save(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, data, 0644)
}

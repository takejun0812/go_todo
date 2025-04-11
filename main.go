package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Todo struct {
	Task string `json:"task"`
	Done bool   `json:"done"`
}

const todoFile = "todos.json"

func loadTodos() ([]Todo, error) {
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

func saveTodos(todos []Todo) error {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, data, 0644)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("使い方: todo [add|list|done] [内容]")
		return
	}

	command := args[1]
	todos, err := loadTodos()
	if err != nil {
		fmt.Println("読み込みエラー:", err)
		return
	}

	switch command {
	case "add":
		if len(args) < 3 {
			fmt.Println("ToDoの内容が必要です")
			return
		}
		newTodo := Todo{Task: args[2], Done: false}
		todos = append(todos, newTodo)
		fmt.Printf("追加: %s\n", newTodo.Task)

	case "list":
		if len(todos) == 0 {
			fmt.Println("ToDoはありません")
			return
		}
		fmt.Println("ToDo一覧:")
		for i, t := range todos {
			status := " "
			if t.Done {
				status = "✓"
			}
			fmt.Printf("%d: [%s] %s\n", i+1, status, t.Task)
		}

	case "done":
		if len(args) < 3 {
			fmt.Println("完了するToDo番号を指定してください")
			return
		}
		index, err := strconv.Atoi(args[2])
		if err != nil || index < 1 || index > len(todos) {
			fmt.Println("無効な番号です")
			return
		}
		todos[index-1].Done = true
		fmt.Printf("完了にしました: %s\n", todos[index-1].Task)

	case "delete":
		if len(args) < 3 {
			fmt.Println("削除するTodo番号を指定してください")
			return
		}
		index, err := strconv.Atoi(args[2])
		if err != nil || index < 1 || index > len(todos) {
			fmt.Println("無効な番号です")
			return
		}
		deleted := todos[index-1]
		todos = append(todos[:index-1], todos[index:]...)
		fmt.Printf("削除しました: %s\n", deleted.Task)

	default:
		fmt.Println("不明なコマンドです: add, list, doneが使えます")
	}

	err = saveTodos(todos)
	if err != nil {
		fmt.Println("保存エラー:", err)
	}
}

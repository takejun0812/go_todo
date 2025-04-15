/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"studyGo/todo/todo"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "新しいToDoを追加する",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todos, _ := todo.Load()
		var err error
		var due time.Time
		if dueStr != "" {
			due, err = time.Parse("2006-01-02", dueStr)
			if err != nil {
				fmt.Println("日付の形式が正しくありません（例： 2025-04-15）")
				return
			}
		}
		newTask := todo.Todo{Task: args[0], Done: false, Due: due}
		todos = append(todos, newTask)
		todo.Save(todos)
		fmt.Printf("追加: %s (期限: %s) \n", newTask.Task, newTask.Due.Format("2006-01-02"))
	},
}

var dueStr string

func init() {
	addCmd.Flags().StringVarP(&dueStr, "due", "d", "", "期限(例: 2025-04-15)")
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

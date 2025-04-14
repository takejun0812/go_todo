/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"studyGo/todo/todo"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "ToDoリストを表示する",
	Run: func(cmd *cobra.Command, args []string) {
		todos, _ := todo.Load()
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
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

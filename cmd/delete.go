/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"studyGo/todo/todo"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task number]",
	Short: "指定したタスクを削除する",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("番号は数字で指定してください")
			return
		}

		todos, _ := todo.Load()
		if index < 1 || index > len(todos) {
			fmt.Println("無効な数字です")
			return
		}

		deleted := todos[index-1]
		todos = append(todos[:index-1], todos[index:]...)
		todo.Save(todos)
		fmt.Printf("削除しました: %s\n", deleted.Task)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

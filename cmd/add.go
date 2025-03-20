/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/JonathanGeorgiou/astrolog/internal/database"
	"github.com/JonathanGeorgiou/astrolog/internal/database/models"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task or note.",
	Long: `adds a new task or note to the database.
  Uasge:
  astrolog add task "task title"
  Flags:
  astrolog add task "task title" --description "a longer task description for my task"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskTitle := args[0]
		err := addTask(taskTitle)
		if err != nil {
			log.Fatalf("Failed to add task: %v", err)
		}
		fmt.Printf("✅ Task added: %s\n", taskTitle)
	},
}

func addTask(taskTitle string) error {
	task := models.Task{
		Title:     taskTitle,
		Completed: false,
	}
	result := database.DB.Create(&task)
	return result.Error
}

func init() {
	rootCmd.AddCommand(addCmd)
}

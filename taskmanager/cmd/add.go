package cmd

import (
	"fmt"
	"strings"
	"taskmanger/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ") // join args with space
		_, err := db.CreateTask(task)   // create task
		if err != nil {                 // if error is not nil
			fmt.Println("Something went wrong:", err) // print error
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task) // print message
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

package cmd

import (
	"fmt"
	"os"
	"taskmanger/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := db.AllTasks() // get all tasks
		if err != nil {             // if error is not nil
			fmt.Println("Something went wrong:", err) // print error
			os.Exit(1) // exit with code 1
		}

		if len(tasks) == 0 { // if length of tasks is 0
			fmt.Println("You have no tasks to complete! Why not take a vacation? 🏖") // print message
			return
		}

		fmt.Println("You have the following tasks:") // print message
		for i, task := range tasks {                  // loop through tasks
			fmt.Printf("%d. %s, Key=%d\n", i+1, task.Value, task.Key) // print task
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

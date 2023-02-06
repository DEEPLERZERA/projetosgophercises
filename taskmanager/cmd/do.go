package cmd

import (
	"fmt"
	"strconv"
	"taskmanger/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{ // create do command
	Use:   "do", // use do
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) { // run function
		var ids []int // create ids slice of int
		for _, arg := range args { // loop through args
			id, err := strconv.Atoi(arg) // convert arg to int
			if err != nil { // if error is not nil
				fmt.Printf("Failed to parse the argument: %s", arg) // print error
			} else { // if no error
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks() // get all tasks
		if err != nil { 		   // if error is not nil
			fmt.Println("Something went wrong:", err)
			return
		}

		for _, id := range ids { // loop through ids
			if id <= 0 || id > len(tasks) { // if id is less than or equal to 0 or greater than length of tasks
				fmt.Printf("Invalid task number: %d", id) // print error
				continue
			}

			task := tasks[id-1] // get task
			err := db.DeleteTask(task.Key) // delete task
			if err != nil { // if error is not nil
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s", id, err) // print error
			} else { // if no error
				fmt.Printf("Marked \"%d\" as completed.", id) // print message
			}
		}
		fmt.Println(ids) // print ids
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}

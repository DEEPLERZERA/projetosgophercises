package main

import (
	"fmt"
	"gophercises/taskmanager/cmd"
	"os"
	"path/filepath"
	"taskmanger/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {  // main function
	home, _ := homedir.Dir() // get home directory
	dbPath := filepath.Join(home, "tasks.db") // join home directory with tasks.db
	must(db.Init(dbPath)) // initialize database
	must(cmd.RootCmd.Execute()) // execute root command
}

func must(err error) { // must function
	if err != nil { // if error is not nil
		fmt.Println(err.Error()) // print error
		os.Exit(1) // exit with code 1
	}
}

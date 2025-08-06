package main

import (
	"code-fox/commands"
	"code-fox/database"
	"fmt"
	"log"
	"os"
)

func main() {
	err := database.InitializeDatabaseObjects()
	if err != nil {
		log.Fatal(err)
	}

	if err := commands.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

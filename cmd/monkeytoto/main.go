package main

import (
	"fmt"
	"os/user"

	"github.com/apastori/monkey-toto-interpreter-go/internal/repl"
)

func main() {
	fmt.Println("Main Application from Monkey Toto Interpreter Started")
	var currentUser *user.User
	var err error
	currentUser, err = user.Current()
	if err != nil {
		fmt.Println("Error getting current user:", err)
		panic(err)
	}
	fmt.Println("Starting REPL for user:", currentUser.Username)
	repl.StartREPL(currentUser.Username)
}

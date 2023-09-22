package main

import (
	"fmt"
	"interpreter-mod/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to the InterLang programming language REPL\n", user.Username)
	fmt.Printf("Type in commands to start\n")
	repl.Start(os.Stdin, os.Stdout)
}

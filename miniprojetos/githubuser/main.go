package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Username not provided")
		os.Exit(0)
	}

	username := args[0]
	user := getUser(username)

	fmt.Printf("%#v", user)
}

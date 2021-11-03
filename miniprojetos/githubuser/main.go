package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Username not provided")
		os.Exit(0)
	}

	username := args[0]

	user, err := getUser(username)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%#v", user)
}

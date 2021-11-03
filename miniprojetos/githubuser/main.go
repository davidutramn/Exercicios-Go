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

	printUser(user)
}

func printUser(user *GitHubUser) {
	fmt.Println("Login", user.Login)
	fmt.Println("ID", user.ID)
	fmt.Println("NodeID", user.NodeID)
	fmt.Println("AvatarURL", user.AvatarURL)
	fmt.Println("GravatarID", user.GravatarID)
	fmt.Println("Type", user.Type)
	fmt.Println("SiteAdmin", user.SiteAdmin)
	fmt.Println("Name", user.Name)
	fmt.Println("Company", user.Company)
	fmt.Println("Blog", user.Blog)
	fmt.Println("Location", user.Location)
	fmt.Println("Email", user.Email)
	fmt.Println("Hireable", user.Hireable)
	fmt.Println("Bio", user.Bio)
	fmt.Println("TwitterUsername", user.TwitterUsername)
	fmt.Println("PublicRepos", user.PublicRepos)
	fmt.Println("PublicGists", user.PublicGists)
	fmt.Println("Followers", user.Followers)
	fmt.Println("Following", user.Following)
	fmt.Println("CreatedAt", user.CreatedAt)
	fmt.Println("UpdatedAt", user.UpdatedAt)
}

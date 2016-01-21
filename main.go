package main

import (
	"fmt"
	"log"

	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)
	fmt.Println(avatarUrl("mswift42", client))
}

func avatarUrl(usr string, client *github.Client) string {
	user, _, err := client.Users.Get(usr)
	if err != nil {
		log.Fatal(err)
	}
	return *user.AvatarURL
}

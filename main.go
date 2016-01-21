package main

import (
	"fmt"
	"log"

	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)
	fmt.Println(avatarURL("mswift42", client))
}

func avatarURL(usr string, client *github.Client) string {
	user, _, err := client.Users.Get(usr)
	if err != nil {
		log.Fatal(err)
	}
	return *user.AvatarURL
}

package main

import (
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"launchpad.net/go-unityscopes/v2"
)

const searchCategoryTemplate = `
  "schema-version": 1,
  "template": {
    "category-layout": "grid",
    "card-size": "small"
  },
  "components": {
    "title": "title",
    "art":  "art",
    "subtitle": "username"
  }
}`

func main() {
	client := github.NewClient(nil)
	fmt.Println(avatarURL("mswift42", client))
}

type GithubScope struct {
	base *scopes.ScopeBase
}

func avatarURL(usr string, client *github.Client) string {
	user, _, err := client.Users.Get(usr)
	if err != nil {
		log.Fatal(err)
	}
	return *user.AvatarURL
}

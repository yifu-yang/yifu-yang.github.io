package main

import (
	"fmt"
	"log"
	"net/url"

	"HackerNewsBot"
)

func main() {
	c := new(HackerNewsBot.BotClient)
	var collection *HackerNewsBot.Stories
	var story *HackerNewsBot.Item
	var user *HackerNewsBot.User
	var err error
	collection, err = c.GetStors(HackerNewsBot.Job)
	if err == nil {
		log.Println(collection)
	}
	story, err = c.GetItem((*collection)[0])
	if err == nil {
		log.Println(story.URL)
	}
	user, err = c.GetUser(story.By)
	if err == nil {
		log.Println(url.QueryUnescape(user.ID))
	}
	fmt.Print("Hello world!")
}

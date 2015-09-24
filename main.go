package main

import (
	"os"

	"github.com/go-chat-bot/bot/slack"
	_ "github.com/go-chat-bot/plugins/catfacts"
	_ "github.com/go-chat-bot/plugins/catgif"
	_ "github.com/go-chat-bot/plugins/chucknorris"

	// Import all the commands you wish to use
)

func main() {
	slack.Run(os.Getenv("SLACK_TOKEN"))
}

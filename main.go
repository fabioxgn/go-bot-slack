package main

import (
	"os"

	"github.com/fabioxgn/go-bot"
	_ "github.com/fabioxgn/go-bot/commands/catfacts"
	_ "github.com/fabioxgn/go-bot/commands/catgif"
	_ "github.com/fabioxgn/go-bot/commands/chucknorris"

	// Import all the commands you wish to use
)

func main() {
	bot.RunSlack(os.Getenv("SLACK_TOKEN"))
}

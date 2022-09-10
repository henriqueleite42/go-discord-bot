package main

import (
	"monetizzer/src/commands"
	"monetizzer/src/events"
	"os"

	"github.com/andersfylling/disgord"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("could not load .env file")
	}

	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("DISCORD_BOT_TOKEN"),
		Intents:  disgord.IntentDirectMessages | disgord.IntentGuildMessages | disgord.IntentGuilds | disgord.IntentGuildMembers,
	})
	defer client.Gateway().StayConnectedUntilInterrupted()

	commands.Register(client)

	events.Register(client)
}

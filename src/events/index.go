package events

import (
	buttonclick "monetizzer/src/events/button-click"

	"github.com/andersfylling/disgord"
)

func Register(client *disgord.Client) {
	client.Gateway().BotReady(BotReady(client))

	client.Gateway().InteractionCreate(buttonclick.Listen)
}

package events

import (
	"log"
	"os"
	"time"

	"monetizzer/src/config/colors"
	"monetizzer/src/config/ids"

	"github.com/andersfylling/disgord"
)

func BotReady(client *disgord.Client) func() {
	return func() {
		if os.Getenv("ENV") != "dev" {
			modChannel := client.Channel(ids.StaffBotsChannelId)

			modChannel.CreateMessage(&disgord.CreateMessage{
				Embeds: []*disgord.Embed{
					{
						Title: "New deploy complete!",
						Timestamp: disgord.Time{time.Now()},
						Color: colors.Monetizze(),
					},
				},
			})
		}

		log.Printf("Bot is ready!")
	}
}

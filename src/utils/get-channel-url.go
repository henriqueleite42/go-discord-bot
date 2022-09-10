package utils

import (
	"fmt"

	"github.com/andersfylling/disgord"
)

func GetChannelUrl(channel *disgord.Channel) string {
	return fmt.Sprintf("https://discord.com/channels/%v/%v", channel.GuildID, channel.ID)
}

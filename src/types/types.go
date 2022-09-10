package types

import "github.com/andersfylling/disgord"

type SlashCommand struct {
	ID                disgord.Snowflake
	GuildID           disgord.Snowflake
	Name              string
	Description       string
	Options           []*disgord.ApplicationCommandOption
	DefaultPermission bool
	Function 					func(s disgord.Session, h *disgord.InteractionCreate)
}

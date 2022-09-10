package buttonclick

import (
	"context"
	"monetizzer/src/config/colors"

	"github.com/andersfylling/disgord"
)

func DeleteChannel(s disgord.Session, h *disgord.InteractionCreate, channelID disgord.Snowflake) {
	s.SendInteractionResponse(context.Background(), h, &disgord.CreateInteractionResponse{
		Type: 5,
		Data: &disgord.CreateInteractionResponseData{
			Flags: 1 << 6,
		},
	})

	s.Channel(channelID).WithContext(context.Background()).Delete()

	if h.ChannelID != channelID {
		s.EditInteractionResponse(context.Background(), h, &disgord.Message{
			Embeds: []*disgord.Embed{
				{
					Title: "Canal deletado!",
					Color: colors.Success(),
				},
			},
			Flags: 1 << 6,
		})
	}
}

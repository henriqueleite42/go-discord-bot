package buttonclick

import (
	"context"
	"fmt"
	"monetizzer/src/config/colors"
	"monetizzer/src/config/ids"
	"monetizzer/src/utils"

	"github.com/andersfylling/disgord"
)

func CreateTicket(s disgord.Session, h *disgord.InteractionCreate) {
	s.SendInteractionResponse(context.Background(), h, &disgord.CreateInteractionResponse{
		Type: 5,
		Data: &disgord.CreateInteractionResponseData{
			Flags: 1 << 6,
		},
	})

	channelName := fmt.Sprintf(
		"%v - %v",
		h.Member.User.Username,
		h.Member.User.Discriminator,
	)

	channel, err := s.Guild(ids.MonetizzerGuildId).CreateChannel(channelName, &disgord.CreateGuildChannel{
		Name: channelName,
		Type: 0,
		ParentID: ids.TicketsCategoryId,
		PermissionOverwrites: []disgord.PermissionOverwrite{
			{
				ID: ids.MonetizzerGuildId,
				Type: 0,
				Deny: 1 << 10,
			},
			{
				ID: h.Member.User.ID,
				Type: 1,
				Allow: 1 << 10,
			},
			{
				ID: ids.StaffRoleId,
				Type: 0,
				Allow: 1 << 10,
			},
		},
	})

	if err != nil {
		s.EditInteractionResponse(context.Background(), h, &disgord.Message{
			Embeds: []*disgord.Embed{
				{
					Title: "Falha ao criar ticket",
					Description: err.Error(),
					Color: colors.Error(),
				},
			},
			Flags: 1 << 6,
		})

		return
	}

	s.Channel(channel.ID).WithContext(context.Background()).CreateMessage(
		&disgord.CreateMessage{
			Content: fmt.Sprintf(`<@&%v> <@%v>`, ids.StaffRoleId, h.Member.User.ID),
			Embeds: []*disgord.Embed{
				{
					Title: "Novo Ticket!",
					Description: "Por favor, descreva seu problema com o maior número de detalhes possivel, e nossos moderadores tentaram te ajudar o mais rápido possivel.",
					Color: colors.Monetizze(),
				},
			},
			Components: []*disgord.MessageComponent{
				{
					Type: 1,
					Components: []*disgord.MessageComponent{
						{
							Type: 2,
							CustomID: fmt.Sprintf("DELETE_CHANNEL/%v", channel.ID),
							Emoji: &disgord.Emoji{
								Name: "🗑️",
							},
							Label: "Encerrar suporte",
							Style: 4,
						},
					},
				},
			},
		},
	)

	s.EditInteractionResponse(context.Background(), h, &disgord.Message{
		Embeds: []*disgord.Embed{
			{
				Title: "Ticket criado!",
				Color: colors.Success(),
			},
		},
		Components: []*disgord.MessageComponent{
			{
				Type: 1,
				Components: []*disgord.MessageComponent{
					{
						Type: 2,
						Label: "Ir para o canal",
						Url: utils.GetChannelUrl(channel),
						Style: 5,
					},
				},
			},
		},
		Flags: 1 << 6,
	})
}

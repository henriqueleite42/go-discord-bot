package commands

import (
	"context"
	"log"
	"os"

	_ "monetizzer/src/commands/adm"
	"monetizzer/src/config/colors"
	"monetizzer/src/config/ids"
	"monetizzer/src/types"
	"monetizzer/src/utils"

	"github.com/andersfylling/disgord"
)

func listen(client *disgord.Client) {
	client.Gateway().InteractionCreate(func(s disgord.Session, h *disgord.InteractionCreate) {
		if (h.Type != 2) {
			return
		}

		command, err := utils.Find(utils.GetSlashCommands(), func(c types.SlashCommand) bool {
			return c.Name == h.Data.Name
		})

		if err != nil {
			s.SendInteractionResponse(context.Background(), h, &disgord.CreateInteractionResponse{
				Type: 4,
				Data: &disgord.CreateInteractionResponseData{
					Embeds: []*disgord.Embed{
						{
							Title: "Command not found",
							Color: colors.Error(),
						},
					},
					Flags: 1 << 6,
				},
			})

			return
		}

		command.Function(s, h)
	})
}

func register(client *disgord.Client) {
	for _, command := range utils.GetSlashCommands() {
		slashCommand := &disgord.CreateApplicationCommand{
			Name:              command.Name,
			Description:       command.Description,
			Options:           command.Options,
			DefaultPermission: command.DefaultPermission,
			Type: 1,
		}

		var err error

		if (command.DefaultPermission) {
			err = client.
				ApplicationCommand(
					disgord.ParseSnowflakeString(os.Getenv("DISCORD_APPLICATION_ID")),
				).
				Global().Create(slashCommand)
		} else {
			err = client.
				ApplicationCommand(
					disgord.ParseSnowflakeString(os.Getenv("DISCORD_APPLICATION_ID")),
				).
				Guild(ids.MonetizzerGuildId).Create(slashCommand)
		}

		if err == nil {
			log.Printf("Command '%v' registered\n", command.Name)
		} else {
			log.Fatal(err)
		}
	}
}

func Register(client *disgord.Client) {
	register(client)

	listen(client)
}

package buttonclick

import (
	"strings"

	"github.com/andersfylling/disgord"
)

func Listen(s disgord.Session, h *disgord.InteractionCreate) {
	if (h.Type != 3) {
		return
	}

	// Do Nothing

	if (strings.HasPrefix(h.Data.CustomID, "DO_NOTHING")) {
		return
	}

	// Create Ticket

	if h.Data.CustomID == "CREATE_TICKET" {
		CreateTicket(s, h)
		return
	}

	// Delete Channel

	if (strings.HasPrefix(h.Data.CustomID, "DELETE_CHANNEL/")) {
		channelId := strings.ReplaceAll(h.Data.CustomID, "DELETE_CHANNEL/", "")
		DeleteChannel(s, h, disgord.ParseSnowflakeString(channelId))
		return
	}
}

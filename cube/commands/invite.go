package commands

import (
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/messages"
)

func init() {
	commandprocessor.Commands["invite"] = &commandprocessor.Command{
		Description: "Gives you an invite to the bot.",
		Category:    categories.INFORMATIONAL,
		Function: func(Args *commandprocessor.CommandArgs) {
			InviteURL := "https://discordapp.com/oauth2/authorize?client_id=" + Args.Session.State.User.ID + "&scope=bot&permissions=1546775744"
			messages.GenericText(
				Args.Channel, "Invite:",
				"You can invite the bot [here]("+InviteURL+").",
				nil, Args.Session,
			)
		},
	}
}

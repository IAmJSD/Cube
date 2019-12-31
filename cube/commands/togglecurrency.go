package commands

import (
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/permissions"
)

func init() {
	commandprocessor.Commands["togglecurrency"] = &commandprocessor.Command{
		Description:      "Used to toggle currency on and off in this guild.",
		Category:         categories.CURRENCY,
		PermissionsCheck: permissions.ADMINISTRATOR,
		Function: func(Args *commandprocessor.CommandArgs) {
			cur := currency.GetCurrency(Args.Message.GuildID)
			if cur.Enabled {
				cur.Enabled = false
				currency.SaveCurrency(Args.Message.GuildID, cur)
				messages.Error(Args.Channel, "Currency Disabled:", "Your currency has been disabled.", Args.Session)
			} else {
				cur.Enabled = true
				currency.SaveCurrency(Args.Message.GuildID, cur)
				messages.GenericText(Args.Channel, "Currency Enabled:", "Your currency has been enabled.", nil, Args.Session)
			}
		},
	}
}

package commands

import (
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/drops"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/utils"
)

func init() {
	commandprocessor.Commands["pick"] = &commandprocessor.Command{
		Description:      "Allows you to pick up some currency by its drop ID!",
		Usage:            "<id>",
		Category:         categories.CURRENCY,
		PermissionsCheck: currency.CurrencyEnabled,
		Function: func(Args *commandprocessor.CommandArgs) {
			// Parse the arguments.
			split := utils.SpaceSplit(Args.RawArgs)
			if 1 > len(split) || split[0] == "" {
				messages.NotEnoughArgs(Args.Channel, Args.Session)
				return
			}

			// Get the currency.
			cur := (*Args.Shared)["currency"].(*currency.Currency)

			// Call the pick up function.
			drops.PickUp(split[0], Args.Channel.ID, Args.Message.ID, Args.Channel.GuildID,
				Args.Message.Author.ID, Args.Message.Author.Mention(), cur, Args.Session)
		},
	}
}

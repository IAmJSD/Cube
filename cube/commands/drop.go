package commands

import (
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/drops"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/utils"
	"github.com/jakemakesstuff/Cube/cube/wallets"
	"strconv"
)

func init() {
	commandprocessor.Commands["drop"] = &commandprocessor.Command{
		Description:      "Allows you to drop some currency for others to pick up!",
		Usage:            "<amount>",
		Category:         categories.CURRENCY,
		PermissionsCheck: currency.CurrencyEnabled,
		Function: func(Args *commandprocessor.CommandArgs) {
			// Parse the arguments.
			split := utils.SpaceSplit(Args.RawArgs)
			if 1 > len(split) || split[0] == "" {
				messages.NotEnoughArgs(Args.Channel, Args.Session)
				return
			}
			i, err := strconv.Atoi(split[0])
			if err != nil {
				messages.NotAnInteger(Args.Channel, Args.Session, "amount")
				return
			}

			// Get the currency.
			cur := (*Args.Shared)["currency"].(*currency.Currency)

			// If the number is 0 or negative, throw an error.
			if 0 >= i {
				messages.NotAPositiveInteger(Args.Channel, Args.Session, "amount")
				return
			}

			// Delete the message if possible.
			_ = Args.Session.ChannelMessageDelete(Args.Channel.ID, Args.Message.ID)

			// Check if the person can afford to drop this amount of money.
			b := wallets.GetBalance(Args.Message.Author.ID, Args.Message.GuildID)
			if i > b {
				messages.OutOfFunds(Args.Channel, Args.Session, *cur.Emoji, b)
				return
			}

			// Subtract the amount from the balance.
			err = wallets.SubtractFromBalance(Args.Message.Author.ID, Args.Message.GuildID, int64(i))
			if err == nil {
				// Drop the currency.
				drops.ExecuteDrop(Args.Channel.ID, Args.Session, i, Args.Prefix, Args.Message.Author.Mention()+" has dropped "+*cur.Emoji+" "+strconv.Itoa(i)+".", cur.DropsImage)
			}
		},
	}
}

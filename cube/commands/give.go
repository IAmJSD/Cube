package commands

import (
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/utils"
	"github.com/jakemakesstuff/Cube/cube/wallets"
	"strconv"
)

func init() {
	commandprocessor.Commands["give"] = &commandprocessor.Command{
		Description:      "Allows users to give other users money.",
		Usage:            "<user> <amount>",
		Category:         categories.CURRENCY,
		PermissionsCheck: currency.CurrencyEnabled,
		Function: func(Args *commandprocessor.CommandArgs) {
			// Gets the currency/split arguments.
			cur := (*Args.Shared)["currency"].(*currency.Currency)
			split := utils.SpaceSplit(Args.RawArgs)

			// Processes all needed arguments.
			if 2 > len(split) {
				messages.NotEnoughArgs(Args.Channel, Args.Session)
				return
			}
			amount, err := strconv.Atoi(split[1])
			if err != nil {
				messages.NotAnInteger(Args.Channel, Args.Session, "amount")
				return
			}
			user := utils.ProcessMemberArg(Args.Channel, Args.Session, split[0])
			if user == nil {
				return
			}

			// If the number is 0 or negative, throw an error.
			if 0 >= amount {
				messages.NotAPositiveInteger(Args.Channel, Args.Session, "amount")
				return
			}

			// Check if the person can afford to drop this amount of money.
			b := wallets.GetBalance(Args.Message.Author.ID, Args.Message.GuildID)
			if amount > b {
				messages.OutOfFunds(Args.Channel, Args.Session, *cur.Emoji, b)
				return
			}

			// Subtract the amount from the balance.
			err = wallets.SubtractFromBalance(Args.Message.Author.ID, Args.Message.GuildID, int64(amount))
			if err == nil {
				// Add the currency to the other user.
				_ = wallets.AddToBalance(user.User.ID, Args.Message.GuildID, int64(amount))

				// Send a confirmation.
				messages.GenericText(Args.Channel,
					"Money transferred:", strconv.Itoa(amount)+" "+*cur.Emoji+" has been transferred to "+user.Mention(), nil, Args.Session)
			}
		},
	}
}

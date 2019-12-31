package commands

import (
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/permissions"
	"github.com/jakemakesstuff/Cube/cube/utils"
	"github.com/jakemakesstuff/Cube/cube/wallets"
	"strconv"
)

func init() {
	commandprocessor.Commands["setbal"] = &commandprocessor.Command{
		Description:      "Allows administrators to set the balance of other users. Can break economies if misused!",
		Usage:            "<user> <balance>",
		Category:         categories.CURRENCY,
		PermissionsCheck: permissions.All(currency.CurrencyEnabled, permissions.ADMINISTRATOR),
		Function: func(Args *commandprocessor.CommandArgs) {
			// Gets the currency/split arguments.
			cur := (*Args.Shared)["currency"].(*currency.Currency)
			split := utils.SpaceSplit(Args.RawArgs)

			// Processes all needed arguments.
			if 2 > len(split) {
				messages.NotEnoughArgs(Args.Channel, Args.Session)
				return
			}
			bal, err := strconv.Atoi(split[1])
			if err != nil {
				messages.NotAnInteger(Args.Channel, Args.Session, "balance")
				return
			}
			user := utils.ProcessMemberArg(Args.Channel, Args.Session, split[0])
			if user == nil {
				return
			}

			// Handle the balance setting.
			wallets.SetBalance(user.User.ID, Args.Channel.GuildID, int64(bal))

			// Send a confirmation.
			messages.GenericText(Args.Channel,
				"Balance set:", user.Mention() + "'s balance has been set to " + strconv.Itoa(bal) + *cur.Emoji, nil, Args.Session)
		},
	}
}

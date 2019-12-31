package commands

import (
	"fmt"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/wallets"
)

func init() {
	commandprocessor.Commands["balance"] = &commandprocessor.Command{
		Description:      "Gets your balance in this guild.",
		Category:         categories.CURRENCY,
		PermissionsCheck: currency.CurrencyEnabled,
		Function: func(Args *commandprocessor.CommandArgs) {
			cur := (*Args.Shared)["currency"].(*currency.Currency)
			bal := wallets.GetBalance(Args.Message.Author.ID, Args.Message.GuildID)
			if cur.Emoji == nil {
				Emoji := "💵"
				cur.Emoji = &Emoji
			}
			messages.GenericText(Args.Channel, "Balance:", fmt.Sprintf("You have %v %s in your wallet!",
				bal, *cur.Emoji), nil, Args.Session)
		},
	}
}

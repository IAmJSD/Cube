package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/styles"
	"github.com/jakemakesstuff/Cube/cube/utils"
	"github.com/jakemakesstuff/Cube/cube/wallets"
	"math/rand"
	"strconv"
	"strings"
)

func init() {
	commandprocessor.Commands["flip"] = &commandprocessor.Command{
		Description:      "Allows the user to flip a coin.",
		Usage:            "<amount> <h/t>",
		Category:         categories.GAMBLING,
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
			amount, err := strconv.Atoi(split[0])
			if err != nil {
				messages.NotAnInteger(Args.Channel, Args.Session, "amount")
				return
			}
			HTArg := strings.ToLower(split[1])
			if HTArg != "h" && HTArg != "t" {
				messages.Error(Args.Channel, "Need heads or tales:", "The second argument needs to be h for heads or t for tails.", Args.Session)
			}
			UserHeads := HTArg == "h"

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
				// Flip a coin.
				IsHeads := rand.Uint64()&(1<<63) == 0
				LossWinText := "You lost!"
				if IsHeads == UserHeads {
					// The user wins!
					LossWinText = "You win double!"
					_ = wallets.AddToBalance(Args.Message.Author.ID, Args.Message.GuildID, int64(amount*2))
				}
				Description := "The coin flipped to "
				Image := &discordgo.MessageEmbedImage{}
				if IsHeads {
					Description += "heads. "
					Image.URL = "https://i.imgur.com/8hpMNW4.png"
				} else {
					Description += "tails. "
					Image.URL = "https://i.imgur.com/x2QauOa.png"
				}
				Description += LossWinText
				_, _ = Args.Session.ChannelMessageSendComplex(Args.Channel.ID, &discordgo.MessageSend{
					Embed: &discordgo.MessageEmbed{
						Description: Description,
						Color:       styles.Generic,
						Image:       Image,
					},
				})
			}
		},
	}
}

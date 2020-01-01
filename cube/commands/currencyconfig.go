package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/embed_menus"
	"github.com/jakemakesstuff/Cube/cube/permissions"
	"github.com/jakemakesstuff/Cube/cube/styles"
	messagewaiter "github.com/jakemakesstuff/Cube/cube/wait_for_message"
	"github.com/jakemakesstuff/Cube/cube/wallets"
	"strings"
)

// CreateCurrencyMenu is used to create the currency config menu.
func CreateCurrencyMenu(MenuID string, GuildID string, msg *discordgo.Message, cur *currency.Currency) *embedmenus.EmbedMenu {
	// Creates the embed menu.
	Menu := embedmenus.NewEmbedMenu(
		discordgo.MessageEmbed{
			Title:       "Cube Currency Configuration",
			Description: "Using this menu, you can configure Cube's currency.",
			Color:       styles.Generic,
		}, &embedmenus.MenuInfo{
			MenuID: MenuID,
			Author: msg.Author.ID,
			Info:   []string{},
		},
	)

	// Add a toggle for the currency.
	var EnableDisableCurrency string
	if cur.Enabled {
		EnableDisableCurrency = "Disable Currency"
	} else {
		EnableDisableCurrency = "Enable Currency"
	}
	Menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       "🔐",
			Name:        EnableDisableCurrency,
			Description: "Toggles the currency in this guild.",
		},
		Function: func(ChannelID string, MessageID string, menu *embedmenus.EmbedMenu, client *discordgo.Session) {
			// Preform the action.
			cur.Enabled = !cur.Enabled

			// Redraw the embed.
			_ = client.MessageReactionsRemoveAll(ChannelID, MessageID)
			currency.SaveCurrency(GuildID, cur)
			CreateCurrencyMenu(MenuID, GuildID, msg, cur).Display(ChannelID, MessageID, client)
		},
	})

	// Used to purge all guild wallets.
	Menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       "💥",
			Name:        "Purge ALL guild wallets",
			Description: "Can be used to start fresh with a currency. Here be dragons, you will be asked to confirm!",
		},
		Function: func(ChannelID string, MessageID string, menu *embedmenus.EmbedMenu, client *discordgo.Session) {
			// Remove all reactions.
			_ = client.MessageReactionsRemoveAll(ChannelID, MessageID)

			// Show the warning.
			_, _ = client.ChannelMessageEditComplex(&discordgo.MessageEdit{
				Embed:   &discordgo.MessageEmbed{
					Title:       "Here be dragons!",
					Description: "This will delete __**EVERY SINGLE GUILD WALLET**__ and cannot be undone. Type \"yes\" if you wish to continue or anything else to not continue (**this is not case sensitive, pranking fellow admins with this isn't smart!**).",
					Color:       styles.Error,
				},
				ID:      MessageID,
				Channel: ChannelID,
			})

			// Wait for the "yes" message or something else.
			confirmationOrNot := messagewaiter.WaitForMessage(ChannelID, msg.Author.ID, 0)
			if strings.ToLower(strings.Trim(confirmationOrNot.Content, " ")) == "yes" {
				wallets.PurgeGuildWallets(msg.GuildID)
			}

			// Try and clean up the message.
			_ = client.ChannelMessageDelete(msg.ChannelID, confirmationOrNot.ID)

			// Redraw the embed.
			CreateCurrencyMenu(MenuID, GuildID, msg, cur).Display(ChannelID, MessageID, client)
		},
	})

	// Returns the embed menu.
	return &Menu
}

func init() {
	commandprocessor.Commands["currencyconfig"] = &commandprocessor.Command{
		Description:      "Used to configure the currency in this guild.",
		Category:         categories.CURRENCY,
		PermissionsCheck: permissions.ADMINISTRATOR,
		Function: func(Args *commandprocessor.CommandArgs) {
			// Gets the currency.
			cur := currency.GetCurrency(Args.Message.GuildID)

			// Gets the menu ID.
			MenuID := uuid.Must(uuid.NewRandom()).String()

			// Create the message.
			m, err := Args.Session.ChannelMessageSendComplex(Args.Message.ChannelID, &discordgo.MessageSend{
				Embed: &discordgo.MessageEmbed{
					Title: "Loading...",
				},
			})
			if err != nil {
				return
			}

			// Show the embed.
			CreateCurrencyMenu(MenuID, Args.Message.GuildID, Args.Message, cur).Display(Args.Channel.ID, m.ID, Args.Session)
		},
	}
}

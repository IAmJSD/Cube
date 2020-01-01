package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/embed_menus"
	"github.com/jakemakesstuff/Cube/cube/message_waiter"
	"github.com/jakemakesstuff/Cube/cube/permissions"
	"github.com/jakemakesstuff/Cube/cube/reaction_waiter"
	"github.com/jakemakesstuff/Cube/cube/styles"
	"github.com/jakemakesstuff/Cube/cube/wallets"
	"strconv"
	"strings"
)

// CurrencyDropsMenu is the menu used for currency drops.
func CurrencyDropsMenu(Parent embedmenus.EmbedMenu, msg *discordgo.Message, MessageID string, client *discordgo.Session, cur *currency.Currency) {
	// Creates the menu.
	Menu := embedmenus.NewEmbedMenu(
		discordgo.MessageEmbed{
			Title:       "Currency Drops",
			Description: "This menu allows you to configure currency drops.",
			Color:       styles.Generic,
		}, &embedmenus.MenuInfo{
			MenuID: Parent.MenuInfo.MenuID,
			Author: msg.Author.ID,
			Info:   []string{},
		},
	)
	Menu.AddParentMenu(&Parent)

	// Adds a back button.
	Menu.AddBackButton()

	// Add a toggle for currency drops.
	var EnableDisableCurrencyDrops string
	var EnabledEmoji string
	if cur.DropsEnabled {
		EnableDisableCurrencyDrops = "Disable Currency Drops"
		EnabledEmoji = "❗"
	} else {
		EnableDisableCurrencyDrops = "Enable Currency Drops"
		EnabledEmoji = "✅"
	}
	Menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       EnabledEmoji,
			Name:        EnableDisableCurrencyDrops,
			Description: "Toggles the currency drops in this guild.",
		},
		Function: func(ChannelID string, MessageID string, menu *embedmenus.EmbedMenu, client *discordgo.Session) {
			// Preform the action.
			cur.DropsEnabled = !cur.DropsEnabled

			// Redraw the embed.
			_ = client.MessageReactionsRemoveAll(ChannelID, MessageID)
			currency.SaveCurrency(msg.GuildID, cur)
			CurrencyDropsMenu(Parent, msg, MessageID, client, cur)
		},
	})

	// Handles whether an image should be shown with the drop.
	Menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       "🌄",
			Name:        "Drop Image",
			Description: "Allows you to configure showing an image URL with the drop or turning it off.",
		},
		Function: func(ChannelID string, MessageID string, menu *embedmenus.EmbedMenu, client *discordgo.Session) {
			// Clear the reactions.
			_ = client.MessageReactionsRemoveAll(ChannelID, MessageID)

			// Shows the box which prompts the user to insert an image URL or "disable".
			_, _ = client.ChannelMessageEditComplex(&discordgo.MessageEdit{
				Embed: &discordgo.MessageEmbed{
					Title:       "Type your image URL or \"disable\":",
					Description: "Please enter your image URL. If you wish to disable the image embed, please type \"disable\".",
					Color:       styles.Generic,
				},
				ID:      MessageID,
				Channel: ChannelID,
			})
			usrMessage := messagewaiter.WaitForMessage(ChannelID, msg.Author.ID, 0)

			// Check if we should disable the embed.
			if strings.ToLower(strings.Trim(usrMessage.Content, " ")) == "disable" {
				// Yes we should.
				cur.DropsImage = nil
			} else {
				// No we shouldn't.
				cur.DropsImage = &usrMessage.Content
			}

			// Try and clean up the message.
			_ = client.ChannelMessageDelete(msg.ChannelID, usrMessage.ID)

			// Redraw the embed.
			currency.SaveCurrency(msg.GuildID, cur)
			CurrencyDropsMenu(Parent, msg, MessageID, client, cur)
		},
	})

	// Handles the drop amount.
	DropsAmount := "100 (default)"
	if cur.DropsAmount != nil {
		DropsAmount = strconv.Itoa(*cur.DropsAmount)
	}
	Menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       "💵",
			Name:        "Drop Amount",
			Description: "Allows you to select how much will be dropped.\n**Current Amount:** " + DropsAmount,
		},
		Function: func(ChannelID string, MessageID string, menu *embedmenus.EmbedMenu, client *discordgo.Session) {
			// Clear the reactions.
			_ = client.MessageReactionsRemoveAll(ChannelID, MessageID)

			// Shows the box which prompts the user to insert an image URL or "disable".
			_, _ = client.ChannelMessageEditComplex(&discordgo.MessageEdit{
				Embed: &discordgo.MessageEmbed{
					Title:       "Please enter the amount:",
					Description: "Please enter the amount which you would like to drop.",
					Color:       styles.Generic,
				},
				ID:      MessageID,
				Channel: ChannelID,
			})
			usrMessage := messagewaiter.WaitForMessage(ChannelID, msg.Author.ID, 0)

			// Gets the int value.
			i, err := strconv.Atoi(strings.Trim(usrMessage.Content, " "))
			if err == nil {
				cur.DropsAmount = &i
			}

			// Try and clean up the message.
			_ = client.ChannelMessageDelete(msg.ChannelID, usrMessage.ID)

			// Redraw the embed.
			currency.SaveCurrency(msg.GuildID, cur)
			CurrencyDropsMenu(Parent, msg, MessageID, client, cur)
		},
	})

	// Displays the menu.
	Menu.Display(msg.ChannelID, MessageID, client)
}

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
	var EnabledEmoji string
	if cur.Enabled {
		EnableDisableCurrency = "Disable Currency"
		EnabledEmoji = "❗"
	} else {
		EnableDisableCurrency = "Enable Currency"
		EnabledEmoji = "✅"
	}
	Menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       EnabledEmoji,
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

	// Used to set the reaction used for the currency.
	var Emoji string
	if cur.Emoji != nil {
		Emoji = *cur.Emoji
	} else {
		Emoji = "💵 (default)"
	}
	Menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       "💵",
			Name:        "Currency Emoji",
			Description: "Allows you to set the emoji used for currency.\n**Current Emoji:** " + Emoji,
		},
		Function: func(ChannelID string, MessageID string, menu *embedmenus.EmbedMenu, client *discordgo.Session) {
			// Remove all reactions.
			_ = client.MessageReactionsRemoveAll(ChannelID, MessageID)

			// Show the emoji popup.
			_, _ = client.ChannelMessageEditComplex(&discordgo.MessageEdit{
				Embed: &discordgo.MessageEmbed{
					Title:       "React with the emoji:",
					Description: "Please react to this embed with the emoji you want for the currency.",
					Color:       styles.Generic,
				},
				ID:      MessageID,
				Channel: ChannelID,
			})

			// Wait for the emoji.
			emoji := reactionwaiter.WaitForReaction(MessageID, msg.Author.ID, 0)

			// Handles saving the emoji.
			e := emoji.MessageFormat()
			cur.Emoji = &e
			currency.SaveCurrency(GuildID, cur)

			// Clean up the emoji.
			_ = client.MessageReactionsRemoveAll(ChannelID, MessageID)

			// Redraw the embed.
			CreateCurrencyMenu(MenuID, GuildID, msg, cur).Display(ChannelID, MessageID, client)
		},
	})

	// Used to handle currency drops.
	Menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       "💸",
			Name:        "Currency Drops",
			Description: "Allows you to configure currency drops in your guild.",
		},
		Function: func(ChannelID string, MessageID string, _ *embedmenus.EmbedMenu, client *discordgo.Session) {
			// Remove all reactions.
			_ = client.MessageReactionsRemoveAll(ChannelID, MessageID)

			// Draw the embed.
			CurrencyDropsMenu(Menu, msg, MessageID, client, cur)
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
				Embed: &discordgo.MessageEmbed{
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

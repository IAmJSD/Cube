package roleshop

// TODO: Implement admin panel.

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/embed_menus"
	"github.com/jakemakesstuff/Cube/cube/styles"
	"github.com/jakemakesstuff/Cube/cube/wallets"
	"math"
	"strconv"
	"time"
)

// MaxElements are the max elements per page.
var MaxElements = 20

// ShowRoleShop is designed to show the role shop/role shop config.
func ShowRoleShop(
	Page int, Parent *embedmenus.EmbedMenu, Currency *currency.Currency,
	Config bool, MenuID string, msg *discordgo.Message, session *discordgo.Session,
	MessageID string,
) {
	// Get the total amount of pages.
	TotalPages := int(math.Ceil(float64(len(Currency.RoleShop)) / float64(MaxElements)))
	if TotalPages == 0 {
		TotalPages = 1
	}

	// Do a sanity check of the current page.
	if Page > TotalPages {
		Page = TotalPages
	}

	// Create the menu.
	Description := "Role Shop "
	if Config {
		Description += "Configuration "
	}
	Description += "(Page " + strconv.Itoa(Page) + "/" + strconv.Itoa(TotalPages) + ")"
	menu := embedmenus.NewEmbedMenu(discordgo.MessageEmbed{
		Description: Description,
	}, &embedmenus.MenuInfo{
		MenuID: MenuID,
		Author: msg.Author.ID,
		Info:   []string{},
	})

	// Add the parent menu if it exists.
	if Parent != nil {
		menu.AddParentMenu(Parent)
		menu.AddBackButton()
	}

	// Handle the role.
	var HandleRole func(role *discordgo.Role, index int, TrialAllowed bool, Cost int)
	HandleRole = func(role *discordgo.Role, index int, TrialAllowed bool, Cost int) {
		// Create the menu.
		RoleMenu := embedmenus.NewEmbedMenu(discordgo.MessageEmbed{
			Description: role.Name,
		}, &embedmenus.MenuInfo{
			MenuID: MenuID,
			Author: msg.Author.ID,
			Info:   []string{},
		})

		// Allows the user to return back to the main menu.
		RoleMenu.AddParentMenu(&menu)
		RoleMenu.AddBackButton()

		if Config {
			// This is a config.

			// Remove the role.
			RoleMenu.Reactions.Add(embedmenus.MenuReaction{
				Button: embedmenus.MenuButton{
					Emoji:       "🪓",
					Name:        "Remove Role",
					Description: "Removes the role from the store.",
				},
				Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
					// Remove all reactions.
					_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

					// Remove the role.
					copy(Currency.RoleShop[index:], Currency.RoleShop[index+1:])
					Currency.RoleShop = Currency.RoleShop[:len(Currency.RoleShop)-1]
					currency.SaveCurrency(msg.GuildID, Currency)

					// Redraw the parent embed.
					ShowRoleShop(Page, Parent, Currency, Config, MenuID, msg, session, MessageID)
				},
			})

			// Toggle trials.
			var EnableDisableTrials string
			var EnabledEmoji string
			if TrialAllowed {
				EnableDisableTrials = "Disable Role Trials"
				EnabledEmoji = "❗"
			} else {
				EnableDisableTrials = "Enable Role Trials"
				EnabledEmoji = "✅"
			}
			RoleMenu.Reactions.Add(embedmenus.MenuReaction{
				Button: embedmenus.MenuButton{
					Emoji:       EnabledEmoji,
					Name:        EnableDisableTrials,
					Description: "Toggle role trials on this role.",
				},
				Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
					// Remove all reactions.
					_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

					// Toggle role trials.
					TrialAllowed = !TrialAllowed
					Currency.RoleShop[index].TrialAllowed = TrialAllowed
					currency.SaveCurrency(msg.GuildID, Currency)

					// Redraw this embed.
					HandleRole(role, index, TrialAllowed, Cost)
				},
			})
		} else {
			// This is a user thinking about buying the role.

			// Buy the role.
			RoleMenu.Reactions.Add(embedmenus.MenuReaction{
				Button: embedmenus.MenuButton{
					Emoji:       "💸",
					Name:        "Buy Role",
					Description: "Buys the role from the store.",
				},
				Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
					// Remove all reactions.
					_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

					// Ensure that the user has the right amount of money for the role.
					bal := wallets.GetBalance(msg.Author.ID, msg.GuildID)
					if Cost > bal {
						// Display a error.
						_, _ = session.ChannelMessageEditComplex(&discordgo.MessageEdit{
							Content: nil,
							Embed: &discordgo.MessageEmbed{
								Title:       "Not enough funds:",
								Description: "You do not have enough funds. Returning to the role information.",
								Color:       styles.Error,
							},
							ID:      MessageID,
							Channel: msg.ChannelID,
						})
						time.Sleep(time.Second * 3)
					} else {
						// Apply the role to the user then take the money.
						err := session.GuildMemberRoleAdd(msg.GuildID, msg.Author.ID, role.ID)
						if err == nil {
							// Take the money.
							_ = wallets.SubtractFromBalance(msg.Author.ID, msg.GuildID, int64(Cost))

							// Remove the role from listings but DO NOT SAVE IT.
							copy(Currency.RoleShop[index:], Currency.RoleShop[index+1:])
							Currency.RoleShop = Currency.RoleShop[:len(Currency.RoleShop)-1]
						}
					}

					// Redraw the parent embed.
					ShowRoleShop(Page, Parent, Currency, Config, MenuID, msg, session, MessageID)
				},
			})
			// TODO: Implement trials.
		}

		// Show the menu.
		RoleMenu.Display(msg.ChannelID, MessageID, session)
	}

	// Defines all of the role emojis.
	RoleEmojis := []string{
		"🇦", "🇧", "🇨", "🇩", "🇪", "🇫", "🇬", "🇭", "🇮", "🇯", "🇰", "🇱",
		"🇲", "🇳", "🇴", "🇵", "🇶", "🇷", "🇸", "🇹", "🇺", "🇻", "🇼", "🇽", "🇾", "🇿",
	}

	// Add all the elements to the page.
	ElementsFrom := (Page - 1) * MaxElements
	ElementsDone := 0
	ElementsSkipped := 0
	for i, v := range Currency.RoleShop {
		// Check if the role exists.
		Role, err := session.State.Role(msg.GuildID, v.RoleID)
		if err != nil {
			ElementsSkipped++
			continue
		}

		// Check the role paging order.
		if ElementsFrom-ElementsSkipped > i {
			continue
		}
		if ElementsDone == MaxElements {
			break
		}

		// Set the technical information for the role.
		TechnicalInfo := "**Buy or trial this role by clicking the reaction.**"
		if Config {
			TechnicalInfo = "**Configure the role by clicking the reaction.**"
		} else if !v.TrialAllowed {
			TechnicalInfo = "**Buy this role by clicking the reaction.**"
		}

		// Get the role description.
		Description := v.Description + "\n" + TechnicalInfo

		// Set the role button.
		menu.Reactions.Add(embedmenus.MenuReaction{
			Button: embedmenus.MenuButton{
				Emoji:       RoleEmojis[i-ElementsSkipped],
				Name:        Role.Name,
				Description: Description,
			},
			Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
				HandleRole(Role, i, v.TrialAllowed, v.Amount)
			},
		})

		// Add one to done elements.
		ElementsDone++
	}

	// Displays the menu.
	menu.Display(msg.ChannelID, MessageID, session)
}

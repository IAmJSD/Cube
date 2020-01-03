package roleshop

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/embed_menus"
	"github.com/jakemakesstuff/Cube/cube/redis"
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

	// Add back/forward buttons.
	if Page != 1 {
		menu.Reactions.Add(embedmenus.MenuReaction{
			Button: embedmenus.MenuButton{
				Emoji:       "◀️️",
				Name:        "Page Back",
				Description: "Goes back a page.",
			},
			Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
				// Remove all reactions.
				_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

				// Redraw the parent embed.
				ShowRoleShop(Page-1, Parent, Currency, Config, MenuID, msg, session, MessageID)
			},
		})
	}
	if Page != TotalPages {
		menu.Reactions.Add(embedmenus.MenuReaction{
			Button: embedmenus.MenuButton{
				Emoji:       "️▶️",
				Name:        "Page Forward",
				Description: "Goes forward a page.",
			},
			Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
				// Remove all reactions.
				_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

				// Redraw the parent embed.
				ShowRoleShop(Page+1, Parent, Currency, Config, MenuID, msg, session, MessageID)
			},
		})
	}
	if Config {
		menu.Reactions.Add(embedmenus.MenuReaction{
			Button: embedmenus.MenuButton{
				Emoji:       "⛏️",
				Name:        "Add Role",
				Description: "Allows you to add a role into the bot.",
			},
			Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
				// Remove all reactions.
				_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

				// Show the "Add role" screen.
				addRole(msg, MessageID, MenuID, session, Currency, &roleShopConfig{}, &menu, Parent, Page)
			},
		})
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

							// Handle trials.
							handleBuyTrial(role.ID, msg.Author.ID)

							// Remove the role from listings but DO NOT SAVE IT.
							copy(Currency.RoleShop[index:], Currency.RoleShop[index+1:])
							Currency.RoleShop = Currency.RoleShop[:len(Currency.RoleShop)-1]
						}
					}

					// Redraw the parent embed.
					ShowRoleShop(Page, Parent, Currency, Config, MenuID, msg, session, MessageID)
				},
			})

			// Manage the trials.
			if TrialAllowed {
				RoleMenu.Reactions.Add(embedmenus.MenuReaction{
					Button: embedmenus.MenuButton{
						Emoji:       "💳",
						Name:        "Trial Role",
						Description: "Allows you to trial the role for 1 minute from the store.",
					},
					Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
						// Remove all reactions.
						_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

						// Wait for the trial to be over.
						_, _ = session.ChannelMessageEditComplex(&discordgo.MessageEdit{
							Content: nil,
							Embed: &discordgo.MessageEmbed{
								Title:       "You are trialing the role:",
								Description: "You now have 1 minute with the role. Go ahead and try the role!",
								Color:       styles.Generic,
							},
							ID:      MessageID,
							Channel: msg.ChannelID,
						})
						RemoveRole := waitForTrial(role.ID, msg.Author.ID)
						redis.Client.SAdd("t:"+role.ID, msg.Author.ID)
						if RemoveRole {
							// Remove the role if possible.
							_ = session.GuildMemberRoleRemove(msg.GuildID, msg.Author.ID, role.ID)

							// Redraw this embed with trials off.
							TrialAllowed = false
							HandleRole(role, index, TrialAllowed, Cost)
						} else {
							// Redraw the parent embed.
							ShowRoleShop(Page, Parent, Currency, Config, MenuID, msg, session, MessageID)
						}
					},
				})
			}
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

		// Don't offer the role to the user if they have it already.
		HasRole := false
		for _, RoleID := range msg.Member.Roles {
			if RoleID == v.RoleID {
				// User has this role.
				HasRole = true
				break
			}
		}
		if HasRole {
			// The user has this role already.
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
		if redis.Client.SIsMember("t:"+Role.ID, msg.Author.ID).Val() {
			// Ensure people cannot infinitely use trials.
			v.TrialAllowed = false
		}
		TechnicalInfo := "**Buy or trial this role by clicking the reaction.**"
		if Config {
			TechnicalInfo = "**Configure the role by clicking the reaction.**"
		} else if !v.TrialAllowed {
			TechnicalInfo = "**Buy this role by clicking the reaction.**"
		}

		// Get the role description.
		Emoji := "💵"
		if Currency.Emoji != nil {
			Emoji = *Currency.Emoji
		}
		Description := v.Description + "\n" + TechnicalInfo + "\n**Cost:** " + Emoji + strconv.Itoa(v.Amount)

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

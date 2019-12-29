package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"sort"
)

func init() {
	commandprocessor.Commands["help"] = &commandprocessor.Command{
		Description:      "Provides help for the bot.",
		Category:          categories.INFORMATIONAL,
		Function: func(Args *commandprocessor.CommandArgs) {
			// The struct for this function.
			type TmpCommand struct {
				Name string
				Cmd *commandprocessor.Command
			}

			// Gets all of the categories.
			Categories := make([]*categories.Category, 0)
			CategoryMap := map[*categories.Category][]*TmpCommand{}
			for name, v := range commandprocessor.Commands {
				if v.Category != nil {
					Exists := false
					for _, c := range Categories {
						if c == v.Category {
							Exists = true
						}
					}
					if !Exists {
						Categories = append(Categories, v.Category)
					}
					CategoryMap[v.Category] = append(CategoryMap[v.Category], &TmpCommand{Name:name, Cmd:v})
				}
			}
			CatStrings := make([]string, len(Categories))
			for i, v := range Categories {
				CatStrings[i] = v.Name
			}
			sort.Strings(CatStrings)
			OrderedCategories := make([]*categories.Category, 0)
			for _, v := range CatStrings {
				for _, cat := range Categories {
					if cat.Name == v {
						// This is first.
						OrderedCategories = append(OrderedCategories, cat)
						break
					}
				}
			}

			// Handles building all of the embeds.
			Embeds := make([]*discordgo.MessageEmbed, 0)
			for _, v := range OrderedCategories {
				Commands := CategoryMap[v]
				i := 0
				Embed := &discordgo.MessageEmbed{
					Title:v.Name + " Commands",
					Description:v.Description,
				}
				for _, cmd := range Commands {
					if cmd.Cmd.PermissionsCheck != nil {
						check, _ := cmd.Cmd.PermissionsCheck(Args)
						if !check {
							continue
						}
					}
					if i == 25 {
						Embeds = append(Embeds, Embed)
						Embed = &discordgo.MessageEmbed{
							Title:v.Name,
							Description:v.Description,
						}
					}
					i++
					Embed.Fields = append(Embed.Fields, &discordgo.MessageEmbedField{
						Name:   Args.Prefix + cmd.Name + " " + cmd.Cmd.Usage,
						Value:  cmd.Cmd.Description,
						Inline: false,
					})
				}
				Embeds = append(Embeds, Embed)
			}

			// Sends the embeds.
			c, err := Args.Session.UserChannelCreate(Args.Message.Author.ID)
			if err != nil {
				messages.Error(Args.Channel, "Failed to DM:", "Failed to DM you! Do you have DM's off or me blocked?", Args.Session)
				return
			}
			for _, v := range Embeds {
				_, err := Args.Session.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{Embed:v})
				if err != nil {
					messages.Error(Args.Channel, "Failed to DM:", "Failed to DM you! Do you have DM's off or me blocked?", Args.Session)
					return
				}
			}
			messages.GenericText(Args.Channel, "DM'd help:", "I have DM'd you help!", nil, Args.Session)
		},
	}
}

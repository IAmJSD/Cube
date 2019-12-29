package commands

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"io/ioutil"
)

// CreditFields defines all of the fields in the credits screen.
var CreditFields []*discordgo.MessageEmbedField

// Initialises the fields.
func init() {
	f, err := ioutil.ReadFile("./credits.json")
	if err != nil {
		panic(err)
	}
	var Interface map[string]interface{}
	err = json.Unmarshal(f, &Interface)
	if err != nil {
		panic(err)
	}
	Interfaces := Interface["credits"].([]interface{})
	CreditFields = make([]*discordgo.MessageEmbedField, len(Interfaces))
	for i, v := range Interfaces {
		CreditFields[i] = &discordgo.MessageEmbedField{
			Name:   v.(map[string]interface{})["username"].(string),
			Value:  v.(map[string]interface{})["description"].(string),
			Inline: true,
		}
	}
}

func init() {
	commandprocessor.Commands["credits"] = &commandprocessor.Command{
		Description: "Shows the credits for the bot.",
		Category:    categories.INFORMATIONAL,
		Function: func(Args *commandprocessor.CommandArgs) {
			messages.GenericText(
				Args.Channel,
				"Credits:",
				"All of the people listed contributed to Cube. Say thanks to them!",
				CreditFields, Args.Session,
			)
		},
	}
}

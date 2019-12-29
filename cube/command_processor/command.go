package commandprocessor

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"time"
)

// CommandArgs are the arguments given to commands.
type CommandArgs struct {
	RawArgs string
	Prefix string
	Message *discordgo.Message
	Channel *discordgo.Channel
	Session *discordgo.Session
	StartTime *time.Time
}

// Command is the structure which all commands will follow.
type Command struct {
	Description string `json:"description"`
	Usage string `json:"usage"`
	Category *categories.Category `json:"category"`
	PermissionsCheck func(Args *CommandArgs) (bool, string) `json:"-"`
	Function func(Args *CommandArgs) `json:"-"`
}

// Commands defines all of the commands.
var Commands = map[string]*Command{}

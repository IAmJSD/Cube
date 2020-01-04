package commandprocessor

import (
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/sentry-go"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/redis"
	"github.com/jakemakesstuff/Cube/cube/utils"
	"os"
	"strings"
	"time"
)

// DefaultPrefix defines the default prefix which the bot uses.
var DefaultPrefix = os.Getenv("DEFAULT_PREFIX")

// Processor is used to process messages.
func Processor(Message *discordgo.Message, Channel *discordgo.Channel, Session *discordgo.Session, StartTime *time.Time) {
	// Gets the prefix.
	Prefix := DefaultPrefix
	r, err := redis.Client.Get("p:" + Message.GuildID).Result()
	if err == nil {
		// Set the new prefix.
		Prefix = r
	} else if err != redis.Nil {
		sentry.CaptureException(err)
		panic(err)
	}

	// Defines the length to trim from the message.
	PrefixLen := 0

	// Verifies the prefix.
	if strings.HasPrefix(Message.Content, Prefix) {
		// The prefix was used!
		PrefixLen = len(Prefix)
	} else {
		// Check if the bot was mentioned. If not, return.
		PrefixLen = utils.CheckMention(Message.Content, &Session.State.User.ID).Len
		if PrefixLen == 0 {
			return
		}
	}

	// Trim the prefix from the content.
	Content := Message.Content[PrefixLen:]

	// Get the command name.
	CommandName := ""
	CommandLen := 0
	for _, v := range Content {
		CommandLen++
		if v == ' ' {
			// Ignore whitespace if this is the beginning. If not, break.
			if CommandName == "" {
				// Ignore this.
				continue
			} else {
				// Break here.
				break
			}
		} else {
			// Add to the command name.
			CommandName += string(v)
		}
	}

	// Re-trim the content to remove the command name.
	Content = Content[CommandLen:]

	// Get the command from the map.
	cmdlower := strings.ToLower(CommandName)
	cmd, ok := Commands[cmdlower]
	if !ok {
		// This is not a command. Check aliases.
		found := false
		Aliases := redis.Client.SMembers("a:" + Message.GuildID).Val()
		for _, v := range Aliases {
			ssplit := strings.Split(v, " ")
			if ssplit[0] == cmdlower {
				// This is an alias.
				cmdlower = ssplit[1]
				cmd, ok = Commands[cmdlower]
				if !ok {
					// This isn't a valid command.
					return
				}
				found = true
				break
			}
		}
		if !found {
			// No alias found!
			return
		}
	}

	// Defines the main command args struct.
	Args := &CommandArgs{
		RawArgs:   Content,
		Message:   Message,
		Channel:   Channel,
		Session:   Session,
		StartTime: StartTime,
		Prefix:    Message.Content[:PrefixLen],
		Shared:    &map[string]interface{}{},
	}

	// Do the permissions check specified.
	if cmd.PermissionsCheck != nil {
		r, msg := cmd.PermissionsCheck(Args)
		if !r {
			messages.Error(Channel, "Incorrect Permissions", msg, Session)
			return
		}
	}

	// Run the command.
	cmd.Function(Args)
}

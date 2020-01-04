package commands

import (
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/permissions"
	"github.com/jakemakesstuff/Cube/cube/redis"
	"github.com/jakemakesstuff/Cube/cube/wallets"
	"github.com/jakemakesstuff/structuredhttp"
	"strings"
	"sync"
	"time"
)

// aliasFormatter is used to format an alias so it is oven ready.
func aliasFormatter(Data string) string {
	s := strings.Split(Data, " ")
	if len(s) == 0 {
		s[0] = ""
	}
	return strings.ToLower(s[0])
}

func init() {
	commandprocessor.Commands["loadconfig"] = &commandprocessor.Command{
		Description:      "Allows you to load a guilds configuration. Useful for loading in changes made from the dumpconfig command. __**THIS WILL OVERWRITE YOUR CURRENT CONFIG!**__",
		Category:         categories.ADMINISTRATOR,
		PermissionsCheck: permissions.ADMINISTRATOR,
		Function: func(Args *commandprocessor.CommandArgs) {
			// Check if the JSON file is there.
			if len(Args.Message.Attachments) != 1 {
				messages.Error(Args.Channel, "File missing:", "The JSON file appears to be missing. Please attach it with the command.", Args.Session)
				return
			}

			// Get the attachment.
			Attachment := Args.Message.Attachments[0]

			// Check the file size.
			if Attachment.Size > 10000000 {
				// This is over 10MB, lol no.
				messages.Error(Args.Channel, "File over 10MB:", "The file is too large to be processed.", Args.Session)
				return
			}

			// Parse the JSON.
			r, err := structuredhttp.GET(Attachment.URL).Timeout(time.Second * 10).Run()
			if err != nil {
				sentry.CaptureException(err)
				return
			}
			err = r.RaiseForStatus()
			if err != nil {
				sentry.CaptureException(err)
				return
			}
			d := json.NewDecoder(r.RawResponse.Body)
			var dump dumpedConfig
			err = d.Decode(&dump)
			if err != nil {
				messages.Error(Args.Channel, "Failed to read config:", "Failed to read the JSON file.", Args.Session)
				return
			}

			// Sets the prefix.
			if dump.Prefix == nil {
				redis.Client.Del("p:" + Args.Message.GuildID)
			} else {
				redis.Client.Set("p:"+Args.Message.GuildID, *dump.Prefix, 0)
			}

			// Sets the currency config.
			CurrencyConfig, err := json.Marshal(dump.CurrencyConfig)
			if err != nil {
				sentry.CaptureException(err)
				return
			}
			redis.Client.Set("C:"+Args.Message.GuildID, CurrencyConfig, 0)

			// Handles setting all the wallets.
			wg := sync.WaitGroup{}
			wg.Add(len(dump.Wallets))
			for k, v := range dump.Wallets {
				Value := int64(v)
				go func(Key string, x int64) {
					wallets.SetBalance(Key, Args.Message.GuildID, x)
					wg.Done()
				}(k, Value)
			}
			wg.Wait()

			// Handles loading in all of the aliases.
			redis.Client.Del("a:" + Args.Message.GuildID)
			for k, v := range dump.Aliases {
				k = aliasFormatter(k)
				v = aliasFormatter(v)
				redis.Client.SAdd("a:"+Args.Message.GuildID, k+" "+v)
			}

			// Send a success message.
			messages.GenericText(Args.Channel, "Config loaded:", "Successfully loaded the sent config.", nil, Args.Session)
		},
	}
}

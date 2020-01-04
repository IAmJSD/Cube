package aliases

import (
	"github.com/jakemakesstuff/Cube/cube/redis"
	"strings"
)

// GetAliases is used to get all of the bots aliases.
func GetAliases(GuildID string) map[string]string {
	m := map[string]string{}
	Aliases := redis.Client.SMembers("a:" + GuildID).Val()
	for _, v := range Aliases {
		split := strings.Split(v, " ")
		m[split[0]] = split[1]
	}
	return m
}

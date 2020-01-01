package wallets

import (
	"github.com/jakemakesstuff/Cube/cube/redis"
)

// PurgeGuildWallets is used to purge all wallets belonging to a guild.
func PurgeGuildWallets(GuildID string) {
	Wallets := redis.Client.SMembers("w:" + GuildID).Val()
	for _, v := range Wallets {
		Key := "b:" + GuildID + ":" + v
		Lock := getBalanceLock(Key)
		Lock.Lock()
		redis.Client.Set(Key, "0", 0)
		Lock.Unlock()
	}
}

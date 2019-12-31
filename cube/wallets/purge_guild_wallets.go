package wallets

import (
	"github.com/jakemakesstuff/Cube/cube/redis"
	"sync"
)

// PurgeGuildWallets is used to purge all wallets belonging to a guild.
func PurgeGuildWallets(GuildID string) {
	Wallets := redis.Client.SMembers("w:" + GuildID).Val()
	Locks := make([]*sync.Mutex, len(Wallets))
	for i, v := range Wallets {
		Key := "b:" + GuildID + ":" + v
		Wallets[i] = Key
		Locks[i] = getBalanceLock(Key)
	}
	for _, v := range Locks {
		v.Lock()
	}
	redis.Client.Del(Wallets...)
	for _, v := range Locks {
		v.Unlock()
	}
}

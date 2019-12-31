package currency

import (
	"encoding/json"
	"github.com/jakemakesstuff/Cube/cube/redis"
)

// GetCurrency is used to get the currency from Redis.
func GetCurrency(GuildID string) *Currency {
	c, err := redis.Client.Get("C:" + GuildID).Bytes()
	if err != nil {
		var obj Currency
		_ = json.Unmarshal(c, &obj)
		return &obj
	}
	currency := Currency{}
	j, err := json.Marshal(&currency)
	if err != nil {
		return &currency
	}
	redis.Client.Set("C:"+GuildID, j, 0)
	return &currency
}

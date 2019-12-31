package currency

import "github.com/jakemakesstuff/Cube/cube/command_processor"

// CurrencyEnabled is a check to see if the currency is enabled and then set it to a shared variable.
func CurrencyEnabled(Args *commandprocessor.CommandArgs) (bool, string) {
	Currency := GetCurrency(Args.Channel.GuildID)
	if !Currency.Enabled {
		return false, "This servers currency is currently disabled."
	}
	if Currency.Emoji == nil {
		Emoji := "💵"
		Currency.Emoji = &Emoji
	}
	(*Args.Shared)["currency"] = Currency
	return true, ""
}

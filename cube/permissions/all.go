package permissions

import commandprocessor "github.com/jakemakesstuff/Cube/cube/command_processor"

// All is used to check all permissions in the array.
func All(Permissions ...func(Args *commandprocessor.CommandArgs) (bool, string)) func(Args *commandprocessor.CommandArgs) (bool, string) {
	return func(Args *commandprocessor.CommandArgs) (b bool, s string) {
		for _, v := range Permissions {
			p, e := v(Args)
			if !p {
				return p, e
			}
		}
		return true, ""
	}
}

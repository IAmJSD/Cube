package permissions

import "github.com/jakemakesstuff/Cube/cube/command_processor"

// PermissionsWrapper is used to wrap the permissions.
func permissionsWrapper(PermissionName string, PermissionsHex int) func(Args *commandprocessor.CommandArgs) (bool, string)  {
	return func(Args *commandprocessor.CommandArgs) (bool, string) {
		perms, err := Args.Session.UserChannelPermissions(Args.Message.Author.ID, Args.Channel.ID)
		if err != nil {
			return false, ""
		}
		return (perms & PermissionsHex) == PermissionsHex, "You must have the  \"" + PermissionName + "\" permission to run this command."
	}
}

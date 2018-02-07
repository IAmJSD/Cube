# Editing Plugins
Whether you want to make your own plugins for a private instance of Cube or you want to make pull requests for the official instance/GitHub of Cube, creating and editing plugins is an essential part of the bot. **Please note that for this tutorial, you will need fairly advanced knowledge of Python.**

Each plugin is very simply laid out and uses definitions and attributes.

`app` contains all the main functions (such as but not limited to `dclient`, `mysql_connection`, `config` and `logger`).

If you want something that will be initially executed on the load/reload you can use `execute_on_init` in a regular definition:

```py
def execute_on_init(app):
  app.logger.info("This will print on load/reload.")
```

The rest of the definitions are meant to be async and will throw a exception if they are not. The bots core automatically seperates discord.py callbacks (`["on_message", "on_ready", "on_message_delete", "on_reaction_add", "on_reaction_remove", "on_channel_delete", "on_channel_create", "on_channel_update", "on_member_join", "on_member_remove", "on_member_update", "on_server_join", "on_server_remove", "on_server_update"]`) into one dictionary and the rest are treated as commands. For instance, this is treated as a callback:

```py
async def on_ready(app):
  app.logger.info("This will be loaded when discord.py is loaded.")
```

This should work for any of the other callbacks, just make sure you are aware some may have additional arguments (`on_message` does not because it is treated like a command minus the arguments).

This means that commands can be written like this:

```py
# The name of the function is the command name.
async def hi(app):
  await app.say("Hi! The following arguments were given: {}".format(app.args))
```

Commands have special functions which can be executed in `app`:
- `app.say` (async) - Basically a shortcut of send_message but it sends to the main channel.
- `app.whisper` (async) - Basically a shortcut of send_message but it sends to the message author in DM's.
- `app.args` - A list of arguments in the message.
- `app.message` - The initial message.

You can also set the following attributes:
- `description` - The description of the plugin for help.
- `requires_staff` - When the attribute is set to True, will force the user to have a role containing "staff" (not case sensitive).
- `requires_management` - When the attribute is set to True, will force the user to have a role containing "managers" or "management" (not case sensitive).
- `requires_bot_admin` - When the attribute is set to True, requires the users ID to be in a MySQL database.

These can be set like this (we are using `reload.py` as an example):
```py
async def reload(app):
    app.mysql_connection.commit()
    app.logger.info("Reload called.")
    app.load()
    await app.say("Reloaded!")
# Reloads the bot.

reload.requires_bot_admin = True
# Sets the attribute to say that "reload" needs bot admin.

reload.description = "Reloads the bot."
# Sets a description for "reload".
```

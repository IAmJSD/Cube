# Editing Plugins
Whether you want to make your own plugins for a private instance of Cube or you want to make pull requests for the official instance/GitHub of Cube, creating and editing plugins is an essential part of the bot. **Please note that for this tutorial, you will need fairly advanced knowledge of Python.**

Each plugin is very simply laid out and uses definitions and attributes.

`app` contains all the main functions (such as but not limited to `dclient`, `mysql_connection`, `config` and `logger`).

This is all wrapped around the "Plugin" definition.

```py
def Plugin(app):
    # CODE GOES HERE.
```

The events use a decorator inside this definition.

```py
@app.event
async def on_ready(app):
  app.logger.info("This will be loaded when discord.py is loaded.")
```

This should work for any of the other callbacks, just make sure you are aware some may have additional arguments (`on_message` does not because it is treated like a command minus the arguments).

This means that commands can be written like this:

```py
# The name of the function is the command name unless set otherwise.
@app.command("Description goes here.")
async def hi(app):
    await app.say("Hi! The following arguments were given: {}".format(app.args))
```

Commands have special functions which can be executed in `app`:
- `app.say` (async) - Basically a shortcut of send_message but it sends to the main channel.
- `app.whisper` (async) - Basically a shortcut of send_message but it sends to the message author in DM's.
- `app.args` - A list of arguments in the message.
- `app.message` - The initial message.

You can also set the following other keyword arguments to `app.command`:
- `requires_staff` - When the attribute is set to True, will force the user to have a role containing "staff" (not case sensitive).
- `requires_management` - When the attribute is set to True, will force the user to have a role containing "managers" or "management" (not case sensitive).
- `requires_bot_admin` - When the attribute is set to True, requires the users ID to be in a MySQL database.
- `name` - The actual name of the command. Will use the function name by default.
- `usage` - Shows how to use the arguments of the command.

These can be set like this (we are using `reload.py` as an example):
```py
def Plugin(app):

    @app.command("Reloads the bot.", requires_bot_admin=True)
    async def reload(app):
        app.load_plugins()
        app.config = json.load(
            open(os.path.join(app.cube_root, "config.json"), "r")
        )
        await app.say("Reloaded!")
    # Reloads the bot.
```

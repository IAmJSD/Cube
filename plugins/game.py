import discord
# Imports go here.

async def game(app):
    await app.dclient.change_presence(game=discord.Game(name=' '.join(app.args)))
    await app.say("Game set to {}".format(' '.join(app.args)))
# Allows bot admins to change the game.

game.requires_bot_admin = True
# Sets the attribute to say that "game" needs bot admin.

game.description = "Allows bot admins to change the game."
# Sets a description for "game".

async def on_ready(app):
    await app.dclient.change_presence(game=discord.Game(name=app.config["game"]))
# Sets the game on bot boot.

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

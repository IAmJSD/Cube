import requests, discord
# Imports go here.

async def help(app):

    help_str = ""

    is_management = False
    for role in app.message.author.roles:
        if "managers" in role.name.lower() or "management" in role.name.lower():
            is_management = True

    is_staff = False
    for role in app.message.author.roles:
        if "staff" in role.name.lower():
            is_staff = True

    is_bot_admin = False
    app.mysql_connection.commit()
    sql = "SELECT * FROM `bot_admins` WHERE `user_id` = %s"
    with app.mysql_connection.cursor() as cursor:
        cursor.execute(sql, (app.message.author.id, ))
        is_bot_admin = not cursor.fetchone() is None
        cursor.close()

    prefix = app.get_prefix(app.message.server.id)

    for cmd in app.discord_commands:
        show = True
        try:
            if app.discord_commands[cmd].requires_staff and not is_staff:
                show = False
        except:
            pass
        try:
            if app.discord_commands[cmd].requires_management and not is_management:
                show = False
        except:
            pass
        try:
            if app.discord_commands[cmd].requires_bot_admin and not is_bot_admin:
                show = False
        except:
            pass
        if show:
            try:
                desc = app.discord_commands[cmd].description
            except:
                desc = "No description found."
            help_str = help_str + prefix + cmd + " - " + desc + "\n"

    embed=discord.Embed(title="{} Help".format(app.config["bot_name"]), description="```{}```".format(help_str))
    embed.set_footer(text=app.premade_ver)

    try:
        await app.whisper(embed=embed)
        await app.say(embed=discord.Embed(title="📬 Check DM's"))
    except:
        await app.say(embed=discord.Embed(title="Could not DM.", color=0xff0000))
# The help command.

help.description = "Gets help for the bot."
# Sets a description for "help".

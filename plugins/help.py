import discord
# Imports go here.

async def help(app):

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
    x = [discord.Embed(title="Cube Help:")]
    f = 0

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
            if f == 10:
                x.append(discord.Embed())
                f = 0
            x[-1].add_field(name=prefix + cmd, value=desc, inline=False)
            f = f + 1

    try:
        for y in x:
            await app.whisper(embed=y)
        await app.say(embed=discord.Embed(title="📬 Check DM's"))
    except:
        await app.say(embed=discord.Embed(title="Could not DM.", color=0xff0000))
# The help command.

help.description = "Gets help for the bot."
# Sets a description for "help".

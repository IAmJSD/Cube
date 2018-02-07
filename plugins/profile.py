import discord
# Imports go here.

async def profile(app):
    if len(app.args) != 0 and app.pass_user(app.args[0], app.message.server) == None:
        embed=discord.Embed(title="I could not find that user.",
        description="Try tagging the user again or if you want yourself just run the command with no arguments.",
        color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        if app.args == []:
            user = app.message.author
        else:
            user = app.pass_user(app.args[0], app.message.server)
        embed=discord.Embed(title=user.name + "'s Profile:", colour=user.colour)
        embed.add_field(name="Discriminator:", value=str(user.discriminator), inline=True)
        embed.add_field(name="ID:", value=str(user.id), inline=True)
        embed.add_field(name="Bot Account:", value=str(user.bot), inline=True)
        embed.add_field(name="Highest Role:", value=user.top_role.name, inline=True)
        embed.add_field(name="User:", value=user.mention, inline=True)
        if not user.game is None:
            embed.add_field(name="Game:", value=user.game, inline=True)
        embed.add_field(name="Created:", value=str(user.created_at), inline=True)
        embed.add_field(name="Joined:", value=str(user.joined_at), inline=True)
        strike_count_sql = "SELECT COUNT(*) AS COUNT FROM strikes WHERE user_id = %s AND server_id = %s"
        merit_count_sql = "SELECT COUNT(*) AS COUNT FROM merits WHERE user_id = %s AND server_id = %s"
        kick_count_sql = "SELECT COUNT(*) AS COUNT FROM kicks WHERE user_id = %s AND server_id = %s"
        ban_count_sql = "SELECT COUNT(*) AS COUNT FROM bans WHERE user_id = %s AND server_id = %s"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(strike_count_sql, (user.id, app.message.server.id, ))
            strike_count = cursor.fetchone()["COUNT"]
            cursor.execute(merit_count_sql, (user.id, app.message.server.id, ))
            merit_count = cursor.fetchone()["COUNT"]
            cursor.execute(kick_count_sql, (user.id, app.message.server.id, ))
            kick_count = cursor.fetchone()["COUNT"]
            cursor.execute(ban_count_sql, (user.id, app.message.server.id, ))
            ban_count = cursor.fetchone()["COUNT"]
            cursor.close()
        embed.add_field(name="Strikes:", value=str(strike_count), inline=True)
        embed.add_field(name="Merits:", value=str(merit_count), inline=True)
        embed.add_field(name="Kicks:", value=str(kick_count), inline=True)
        embed.add_field(name="Bans:", value=str(ban_count), inline=True)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
# Allows you to get information about a Discord account.

profile.description = "Allows you to get information about a Discord account."
# Sets a description for "profile".

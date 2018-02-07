import discord
# Imports go here.

async def merits(app):
    if len(app.args) != 0 and app.pass_user(app.args[0], app.message.server) == None:
        embed=discord.Embed(title="I could not find that user.", 
        description="Try tagging the user again or if you want yourself just run the command with no arguments.",
        color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        if len(app.args) == 0:
            user = app.message.author
        else:
            user = app.pass_user(app.args[0], app.message.server)
        merit_count_sql = "SELECT COUNT(*) AS COUNT FROM merits WHERE user_id = %s AND server_id = %s"
        get_merits_sql = "SELECT * FROM merits WHERE user_id = %s AND server_id = %s"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(merit_count_sql, (user.id, app.message.server.id, ))
            merit_count = cursor.fetchone()["COUNT"]
            cursor.execute(get_merits_sql, (user.id, app.message.server.id, ))
            merit_list = cursor.fetchall()
            cursor.close()
        embed=discord.Embed(title="Merits for {}:".format(user.name), 
        description="{} has {} merit(s) in {}.".format(user.name, merit_count, app.message.server.name),
        colour=user.colour)
        embed.set_footer(text=app.premade_ver)
        for merit in merit_list:
            if merit["merit_reason"] == "":
                merit["merit_reason"] = "[No merit reason]"
            staff = app.pass_user(merit["staff_id"], app.message.server)
            if staff is None:
                staff_formatted = "{} [No longer in the server]".format(merit["staff_id"])
            else:
                staff_formatted = "{}".format(staff.name)
            embed.add_field(name="{} (Added by {}):".format(merit["merit_id"], staff_formatted),
            value=merit["merit_reason"], inline=False)
        await app.say(embed=embed)
# Allows you to list merits to a user.

merits.description = "Allows you to list merits to a user."
# Sets a description for "merits".

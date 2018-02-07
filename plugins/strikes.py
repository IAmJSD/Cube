import discord
# Imports go here.

async def strikes(app):
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
        strike_count_sql = "SELECT COUNT(*) AS COUNT FROM strikes WHERE user_id = %s AND server_id = %s"
        get_strikes_sql = "SELECT * FROM strikes WHERE user_id = %s AND server_id = %s"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(strike_count_sql, (user.id, app.message.server.id, ))
            strike_count = cursor.fetchone()["COUNT"]
            cursor.execute(get_strikes_sql, (user.id, app.message.server.id, ))
            strike_list = cursor.fetchall()
            cursor.close()
        embed=discord.Embed(title="Strikes for {}:".format(user.name), 
        description="{} has {} strike(s) in {}.".format(user.name, strike_count, app.message.server.name),
        colour=user.colour)
        embed.set_footer(text=app.premade_ver)
        for strike in strike_list:
            if strike["strike_reason"] == "":
                strike["strike_reason"] = "[No strike reason]"
            staff = app.pass_user(strike["staff_id"], app.message.server)
            if staff is None:
                staff_formatted = "{} [No longer in the server]".format(strike["staff_id"])
            else:
                staff_formatted = "{}".format(staff.name)
            embed.add_field(name="{} (Added by {}):".format(strike["strike_id"], staff_formatted),
            value=strike["strike_reason"], inline=False)
        await app.say(embed=embed)
# Allows you to list strikes to a user.

strikes.description = "Allows you to list strikes to a user."
# Sets a description for "strikes".

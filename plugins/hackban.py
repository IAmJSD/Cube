import discord
# Imports go here.

async def hackban(app):
    if app.args == []:
        embed=discord.Embed(title="No user found.", description="Please provide the user ID and the ban reason as arguments.", color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        user = app.args[0]
        if len(app.args) == 1:
            reason = "No reason found."
        else:
            del app.args[0]
            reason = ' '.join(app.args)
        await app.say("Are you sure you want to hackban a user with the ID `{}` for `{}`? Please say yes in order to run this action.".format(user, reason))
        msg2 = await app.dclient.wait_for_message(author=app.message.author, timeout=30)
        if msg2 is None or not "yes" in msg2.content.lower():
            await app.say("Ban cancelled.")
        else:
            try:
                duser = discord.Object(id=user)
                duser.server = app.message.server
                await app.dclient.ban(duser)
                sql = "INSERT INTO bans(user_id, staff_id, server_id, reason) VALUES(%s, %s, %s, %s)"
                with app.mysql_connection.cursor() as cursor:
                    cursor.execute(sql, (user, app.message.author.id, app.message.server.id, reason, ))
                    cursor.close()
                generic_desc = "A user with the ID `{}` has been banned for `{}`".format(user, reason)
                embed=discord.Embed(title="User hackbanned:", description="{}.".format(generic_desc), color=0xff0000)
                embed.set_footer(text=app.premade_ver)
                await app.say(embed=embed)
                log_embed=discord.Embed(title="User hackbanned:", description="{} by `{}`.".format(generic_desc, app.message.author.name))
                await app.attempt_log(app.message.server.id, log_embed)
            except:
                await app.say("Could not ban the specified user.")
# Allows you to hackban users.

hackban.description = "Allows you to hackban users."
# Sets a description for "hackban".

hackban.requires_staff = True
# Set that this script requires staff.

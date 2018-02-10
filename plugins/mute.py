import discord
# Imports go here.

async def mute(app):
    if app.args == []:
        embed=discord.Embed(title="No user found.", description="Please provide the user to mute and a mute reason.", color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        user = app.pass_user(app.args[0], app.message.server)
        if user == None:
            embed=discord.Embed(title="No user found.", description="The first argument could not be found as a user.", color=0xff0000)
            embed.set_footer(text=app.premade_ver)
            await app.say(embed=embed)
        elif "muted" in str([r.name for r in user.roles]).lower():
            await app.say("User already muted.")
        else:
            if len(app.args) == 1:
                reason = "No reason found."
            else:
                del app.args[0]
                reason = ' '.join(app.args)
            await app.say("Are you sure you want to mute `{}` for `{}`? Please say yes in order to run this action.".format(user.name, reason))
            msg2 = await app.dclient.wait_for_message(author=app.message.author, timeout=30)
            if msg2 is None or not "yes" in msg2.content.lower():
                await app.say("Mute cancelled.")
            else:
                member_name_count_sql = "SELECT COUNT(*) AS COUNT FROM member_roles WHERE server_id = %s"
                with app.mysql_connection.cursor() as cursor:
                    cursor.execute(member_name_count_sql, (app.message.server.id, ))
                    member_name_count = cursor.fetchone()["COUNT"]
                    cursor.close()
                if member_name_count == 0:
                    role_part = ""
                else:
                    get_role = "SELECT * FROM member_roles WHERE server_id = %s"
                    with app.mysql_connection.cursor() as cursor:
                        cursor.execute(get_role, (app.message.server.id, ))
                        role_part = cursor.fetchone()["part_of_role"]
                        cursor.close()
                y = False
                for r in app.message.server.roles:
                    if "muted" in r.name.lower():
                        try:
                            await app.dclient.add_roles(user, r)
                            y = True
                        except:
                            pass
                    elif role_part in r.name.lower() and y and not role_part == "":
                        try:
                            await app.dclient.remove_roles(user, r)
                        except:
                            pass
                if not y:
                    await app.say('Could not mute this user. Either you do not have a role with "muted" in its name (not case sensitive) or the bot did not have permission to apply it to this user.')
                else:
                    generic_desc = "`{}` was muted for `{}`".format(user.name, reason)
                    embed=discord.Embed(title="User muted:", description="{}.".format(generic_desc), color=0xff0000)
                    embed.set_footer(text=app.premade_ver)
                    await app.say(embed=embed)
                    log_embed=discord.Embed(title="User muted:", description="{} by `{}`.".format(generic_desc, app.message.author.name))
                    await app.attempt_log(app.message.server.id, log_embed)
                    sql = "INSERT INTO mutes(user_id, staff_id, server_id, reason) VALUES(%s, %s, %s, %s)"
                    with app.mysql_connection.cursor() as cursor:
                        cursor.execute(sql, (user.id, app.message.author.id, app.message.server.id, reason, ))
                        cursor.close()
# Allows you to mute a user.

mute.description = 'Allows you to mute a user (Looks for a case-insensitive role with "muted" in and adds it to the user).'
# Sets a description for "mute".

mute.requires_staff = True
# Set that this script requires staff.

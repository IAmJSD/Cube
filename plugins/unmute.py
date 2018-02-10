import discord
# Imports go here.

async def unmute(app):
    if app.args == []:
        embed=discord.Embed(title="No user found.", description="Please provide the user to unmute.", color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        user = app.pass_user(app.args[0], app.message.server)
        if user is None:
            embed=discord.Embed(title="Rest in pepperonis.", description="I couldn't find that user. Please make sure you put the user as the first argument.", color=0xff0000)
            embed.set_footer(text=app.premade_ver)
            await app.say(embed=embed)
        elif not "muted" in str([r.name for r in user.roles]).lower():
            await app.say("User is not muted.")
        else:
            y = False
            for r in user.roles:
                if "muted" in r.name.lower():
                    try:
                        await app.dclient.remove_roles(user, r)
                        y = True
                    except:
                        pass
            member_name_count_sql = "SELECT COUNT(*) AS COUNT FROM member_roles WHERE server_id = %s"
            with app.mysql_connection.cursor() as cursor:
                cursor.execute(member_name_count_sql, (app.message.server.id, ))
                member_name_count = cursor.fetchone()["COUNT"]
                cursor.close()
            if member_name_count != 0:
                get_role = "SELECT * FROM member_roles WHERE server_id = %s"
                with app.mysql_connection.cursor() as cursor:
                    cursor.execute(get_role, (app.message.server.id, ))
                    role_part = cursor.fetchone()["part_of_role"]
                    cursor.close()
                x = False
                for role in app.message.server.roles:
                    if not x:
                        if role_part in role.name.lower():
                            try:
                                await app.dclient.add_roles(user, role)
                                x = True
                            except:
                                pass
                if y:
                    generic_desc = "`{}` has been unmuted".format(user.name)
                    embed=discord.Embed(title="User unmuted:", description="{}.".format(generic_desc), color=0x00ff00)
                    embed.set_footer(text=app.premade_ver)
                    await app.say(embed=embed)
                    log_embed=discord.Embed(title="User unmuted:", description="{} by `{}`.".format(generic_desc, app.message.author.name))
                    log_embed.set_footer(text=app.premade_ver)
                    await app.attempt_log(app.message.server.id, log_embed)
                else:
                    await app.say("Could not unmute the user.")
# Allows you to unmute a member.

unmute.description = "Allows you to unmute a member."
# Sets a description for "unmute".

unmute.requires_staff = True
# Set that this script requires staff.

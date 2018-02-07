import discord, random
# Imports go here.

async def strike(app):
    if app.args == []:
        embed=discord.Embed(title="I could not find any arguments.",
                    description="Please supply the user first and then the strike reason. For instance using the default prefix as an example, `{}strike @user reason`.".format(app.config["default_prefix"]),
                    color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        user = app.pass_user(app.args[0], app.message.server)
        if user == None:
            embed=discord.Embed(title="Rest in pepperonis.",
                                description="I couldn't find that user. Please make sure you put the user as the first argument.",
                                color=0xff0000)
            embed.set_footer(text=app.premade_ver)
            await app.say(embed=embed)
        else:
            del app.args[0]
            args = ' '.join(app.args).strip('|').strip(' ')
            x = True
            while x:
                charset = "0123456789"
                strike_id_length = 10
                strike_id = ""
                while not len(strike_id) == strike_id_length:
                    strike_id = strike_id + charset[random.randint(0, len(charset)-1)]
                sql = "SELECT COUNT(*) AS COUNT FROM strikes WHERE strike_id = %s"
                with app.mysql_connection.cursor() as cursor:
                    cursor.execute(sql, (strike_id,))
                    if cursor.fetchone()["COUNT"] == 0:
                        x = False
                    cursor.close()
            insert_sql = "INSERT INTO strikes(strike_id, staff_id, user_id, server_id, strike_reason) VALUES (%s, %s, %s, %s, %s)"
            strike_count_sql = "SELECT COUNT(*) AS COUNT FROM strikes WHERE user_id = %s AND server_id = %s"
            with app.mysql_connection.cursor() as cursor:
                cursor.execute(insert_sql, (strike_id, app.message.author.id, user.id, app.message.server.id, args, ))
                cursor.execute(strike_count_sql, (user.id, app.message.server.id, ))
                strike_count = cursor.fetchone()["COUNT"]
                cursor.close()
            embed=discord.Embed(title="✓ User stricken.",
                                description="{} has been stricken for `{}`. They now have {} strike(s).".format(user.name, args, strike_count),
                                color=0x00ff00)
            embed.set_footer(text=app.premade_ver)
            await app.say(embed=embed)
            embed=discord.Embed(title="User stricken.",
                                description="{} has been stricken by {} for `{}` in {}. They now have {} strike(s).".format(user.name, app.message.author.name, args,
                                app.message.channel.name, strike_count))
            await app.attempt_log(app.message.server.id, embed)
# Allows you to strike a user.

strike.description = "Allows you to strike a user."
# Sets a description for "strike".

strike.requires_staff = True
# Set that this script requires staff.

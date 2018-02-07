import discord, random
# Imports go here.

async def merit(app):
    if app.args == []:
        embed=discord.Embed(title="I could not find any arguments.", 
                    description="Please supply the user first and then the merit reason. For instance using the default prefix as an example, `{}merit @user reason`.".format(app.config["default_prefix"]),
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
                merit_id_length = 10
                merit_id = ""
                while not len(merit_id) == merit_id_length:
                    merit_id = merit_id + charset[random.randint(0, len(charset)-1)]
                sql = "SELECT COUNT(*) AS COUNT FROM merits WHERE merit_id = %s"
                with app.mysql_connection.cursor() as cursor:
                    cursor.execute(sql, (merit_id,))
                    if cursor.fetchone()["COUNT"] == 0:
                        x = False
                    cursor.close()
            insert_sql = "INSERT INTO merits(merit_id, staff_id, user_id, server_id, merit_reason) VALUES (%s, %s, %s, %s, %s)"
            merit_count_sql = "SELECT COUNT(*) AS COUNT FROM merits WHERE user_id = %s AND server_id = %s"
            with app.mysql_connection.cursor() as cursor:
                cursor.execute(insert_sql, (merit_id, app.message.author.id, user.id, app.message.server.id, args, ))
                cursor.execute(merit_count_sql, (user.id, app.message.server.id, ))
                merit_count = cursor.fetchone()["COUNT"]
                cursor.close()
            embed=discord.Embed(title="✓ User rewarded.", 
                                description="{} has been rewarded for `{}`. They now have {} merit(s).".format(user.name, args, merit_count),
                                color=0x00ff00)
            embed.set_footer(text=app.premade_ver)
            await app.say(embed=embed)
# Allows you to reward a user.

merit.description = "Allows you to reward a user."
# Sets a description for "merit".

merit.requires_staff = True
# Set that this script requires staff.

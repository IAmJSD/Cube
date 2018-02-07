import discord
# Imports go here.

async def rmstrikes(app):
    if app.args == []:
        embed=discord.Embed(title="I could not find any arguments.",
            description="Please supply each strike ID you want to remove.",
            color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        response = ""
        for strike in app.args:
            sql = "SELECT COUNT(*) AS COUNT FROM strikes WHERE server_id = %s AND strike_id = %s"
            with app.mysql_connection.cursor() as cursor:
                cursor.execute(sql, (app.message.server.id, strike, ))
                strike_count = cursor.fetchone()["COUNT"]
                cursor.close()
            if strike_count == 0:
                response = response + "\n{} - {}".format(strike, "Strike not found.")
            else:
                sql = "DELETE FROM strikes WHERE server_id = %s AND strike_id = %s"
                with app.mysql_connection.cursor() as cursor:
                    cursor.execute(sql, (app.message.server.id, strike, ))
                    cursor.close()
                response = response + "\n{} - {}".format(strike, "Strike removed.")
        main_embed = discord.Embed(title="Strike removal:", description="```{}```".format(response))
        main_embed.set_footer(text=app.premade_ver)
        await app.say(embed=main_embed)
        log_embed = discord.Embed(title="Strike removal:", description="`{}` removed the following strikes:```{}```".format(app.message.author.name, response))
        log_embed.set_footer(text=app.premade_ver)
        await app.attempt_log(app.message.server.id, log_embed)
# Allows you to remove strikes.

rmstrikes.description = "Allows you to remove strikes by using their strike ID's."
# Sets a description for "rmstrikes".

rmstrikes.requires_staff = True
# Set that this script requires staff.

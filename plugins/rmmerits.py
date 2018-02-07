import discord
# Imports go here.

async def rmmerits(app):
    if app.args == []:
        embed=discord.Embed(title="I could not find any arguments.",
            description="Please supply each merit ID you want to remove.",
            color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        response = ""
        for merit in app.args:
            sql = "SELECT COUNT(*) AS COUNT FROM merits WHERE server_id = %s AND merit_id = %s"
            with app.mysql_connection.cursor() as cursor:
                cursor.execute(sql, (app.message.server.id, merit, ))
                merit_count = cursor.fetchone()["COUNT"]
                cursor.close()
            if merit_count == 0:
                response = response + "\n{} - {}".format(merit, "Merit not found.")
            else:
                sql = "DELETE FROM merits WHERE server_id = %s AND merit_id = %s"
                with app.mysql_connection.cursor() as cursor:
                    cursor.execute(sql, (app.message.server.id, merit, ))
                    cursor.close()
                response = response + "\n{} - {}".format(merit, "Merit removed.")
        main_embed = discord.Embed(title="Merit removal:", description="```{}```".format(response))
        main_embed.set_footer(text=app.premade_ver)
        await app.say(embed=main_embed)
        log_embed = discord.Embed(title="Merit removal:", description="`{}` removed the following merits:```{}```".format(app.message.author.name, response))
        log_embed.set_footer(text=app.premade_ver)
        await app.attempt_log(app.message.server.id, log_embed)
# Allows you to remove merits.

rmmerits.description = "Allows you to remove merits by using their merit ID's."
# Sets a description for "rmmerits".

rmmerits.requires_staff = True
# Set that this script requires staff.

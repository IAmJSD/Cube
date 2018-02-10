import discord
# Imports go here.

async def silence(app):
    silenced_count_sql = "SELECT COUNT(*) AS silenced_count FROM silenced_servers WHERE server_id = %s"
    with app.mysql_connection.cursor() as cursor:
        cursor.execute(silenced_count_sql, (app.message.server.id,))
        silenced_count = cursor.fetchone()["silenced_count"]
        cursor.close()
    if silenced_count != 0:
        await app.say(embed=discord.Embed(title=":zipper_mouth: Already silenced."))
    else:
        silence_sql = "INSERT INTO silenced_servers(server_id) VALUES(%s)"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(silence_sql, (app.message.server.id,))
            cursor.close()
        app.mysql_connection.commit()
        await app.say(embed=discord.Embed(title=":zipper_mouth: Silenced."))
# Allows you to silence the join/leave messages from the bot.

silence.description = "Allows you to silence the join/leave messages from the bot."
# Sets a description for "silence".

silence.requires_staff = True
# Set that this script requires staff.

import discord
# Imports go here.

async def unsilence(app):
    silenced_count_sql = "SELECT COUNT(*) AS silenced_count FROM silenced_servers WHERE server_id = %s"
    with app.mysql_connection.cursor() as cursor:
        cursor.execute(silenced_count_sql, (app.message.server.id,))
        silenced_count = cursor.fetchone()["silenced_count"]
        cursor.close()
    if silenced_count == 0:
        await app.say(embed=discord.Embed(title="😃 Not silenced."))
    else:
        unsilence_sql = "DELETE FROM silenced_servers WHERE server_id = %s"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(unsilence_sql, (app.message.server.id,))
            cursor.close()
        app.mysql_connection.commit()
        await app.say(embed=discord.Embed(title="😃 Unsilenced."))
# Allows you to unsilence the join/leave messages from the bot.

unsilence.description = "Allows you to unsilence the join/leave messages from the bot."
# Sets a description for "unsilence".

unsilence.requires_staff = True
# Set that this script requires staff.

import discord
# Imports go here.

async def main_channel(app):
    channel = app.pass_channel(app.args[0])
    if channel is None:
        embed=discord.Embed(title="Well this is awkward.", 
                            description="I couldn't find the channel you tagged.",
                            color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        try:
            msg = await app.dclient.send_message(channel, "This is a test of inputing and deleting messages into the main channel.")
            await app.dclient.delete_message(msg)
            del_sql = "DELETE FROM main_channels WHERE server_id = %s"
            insert_sql = "INSERT INTO main_channels (server_id, channel_id) VALUES (%s, %s)"
            with app.mysql_connection.cursor() as cursor:
                cursor.execute(del_sql, (app.message.server.id, ))
                cursor.execute(insert_sql, (app.message.server.id, channel.id, ))
                cursor.close()
            app.mysql_connection.commit()
            embed=discord.Embed(title="✓ Done!", 
                    description="I set the main channel to {}.".format(channel.mention),
                    color=0x00ff00)
            embed.set_footer(text=app.premade_ver)
            await app.say(embed=embed)
        except discord.HTTPException:
            embed=discord.Embed(title="Ermmmmmmmm.", 
                    description="I couldn't post a test message to that channel.",
                    color=0xff0000)
            embed.set_footer(text=app.premade_ver)
            await app.say(embed=embed)
# Sets the logging channel.

main_channel.description = "Sets the main channel."
# Sets a description for "main_channel".

main_channel.requires_management = True
# Sets that this script requires the "Management" role.

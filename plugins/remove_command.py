import discord
# Imports go here.

async def remove_command(app):
    if app.args == []:
        embed=discord.Embed(title="I could not find any arguments.", 
                            description="Please supply the command you want to delete.",
                            color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        command = app.args[0].lower()
        cmd_count_sql = "SELECT COUNT(*) AS COUNT FROM custom_commands WHERE server_id = %s AND command = %s"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(cmd_count_sql, (app.message.server.id, command, ))
            cmd_count = cursor.fetchone()["COUNT"]
            cursor.close()
        if cmd_count == 0:
            embed=discord.Embed(title="I could not find any commands.", 
                                description="I could not find a command named `{}`.".format(command),
                                color=0xff0000)
            embed.set_footer(text=app.premade_ver)
            await app.say(embed=embed)
        else:
            cmd_delete_sql = "DELETE FROM custom_commands WHERE server_id = %s AND command = %s"
            with app.mysql_connection.cursor() as cursor:
                cursor.execute(cmd_delete_sql, (app.message.server.id, command, ))
                cursor.close()
            generic_desc = "`{}` has been deleted".format(command)
            embed=discord.Embed(title="✓ Custom command removed.", 
                    description="{}.".format(generic_desc),
                    color=0x00ff00)
            embed.set_footer(text=app.premade_ver)
            await app.say(embed=embed)
            embed=discord.Embed(title="Custom command added.", description="{} by `{}`.".format(generic_desc, app.message.author.name))
            await app.attempt_log(app.message.server.id, embed)
# Allows you to remove a custom command.

remove_command.description = "Allows you to remove a custom command."
# Sets a description for "remove_command".

remove_command.requires_staff = True
# Set that this script requires staff.

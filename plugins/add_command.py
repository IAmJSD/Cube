import discord
# Imports go here.

async def add_command(app):
    if app.args == []:
        embed=discord.Embed(title="I could not find any arguments.", 
            description="Please supply the command first and then the resonse. For instance using the default prefix as an example, `{}add_command command description`.".format(app.config["default_prefix"]),
            color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        if len(app.args) == 1:
            embed=discord.Embed(title="I could not find a response.", 
                description="Please provide one for the second argument.",
                color=0xff0000)
            embed.set_footer(text=app.premade_ver)
            await app.say(embed=embed)
        else:
            command = app.args[0].lower()
            del app.args[0]
            cmd_count_sql = "SELECT COUNT(*) AS COUNT FROM custom_commands WHERE server_id = %s AND command = %s"
            with app.mysql_connection.cursor() as cursor:
                cursor.execute(cmd_count_sql, (app.message.server.id, command, ))
                cmd_count = cursor.fetchone()["COUNT"]
                cursor.close()
            if cmd_count != 0:
                embed=discord.Embed(title="Command already exists.", 
                                    description="`{}` already exists. Please delete it if you want to recreate it.".format(command),
                                    color=0xff0000)
                embed.set_footer(text=app.premade_ver)
                await app.say(embed=embed)
            else:
                sql = "INSERT INTO custom_commands (server_id, command, response) VALUES (%s, %s, %s)"
                with app.mysql_connection.cursor() as cursor:
                    cursor.execute(sql, (app.message.server.id, command, ' '.join(app.args)))
                    cursor.close()
                app.mysql_connection.commit()
                generic_desc = "`{}` was set as a custom command".format(command)
                embed=discord.Embed(title="✓ Custom command added.", 
                                    description="{}.".format(generic_desc),
                                    color=0x00ff00)
                embed.set_footer(text=app.premade_ver)
                await app.say(embed=embed)
                embed=discord.Embed(title="Custom command added.", description="{} by `{}`.".format(generic_desc, app.message.author.name))
                await app.attempt_log(app.message.server.id, embed)
# Allows you to set a custom command.

add_command.description = "Allows you to set a custom command. Use $args$ to repersent any arguments for the command."
# Sets a description for "add_command".

add_command.requires_staff = True
# Set that this script requires staff.

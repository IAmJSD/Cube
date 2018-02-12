import discord
# Imports go here.

async def leave_message(app):
    if app.args == []:
        embed=discord.Embed(title="I could not find any arguments.",
                            description="Please supply the leave message for arguments (using $user$ to repersent the user and $server$ to repersent the server).",
                            color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        args = ' '.join(app.args)
        if args == "$rm$":
            just_delete = True
            visualised_bye = "{} has left this server.".format(app.message.author.name)
        else:
            just_delete = False
            visualised_bye = args.replace("$user$", app.message.author.mention).replace("$server$", app.message.server.name)
        join_delete_sql = "DELETE FROM custom_leave WHERE server_id = %s"
        join_insert_sql = "INSERT INTO custom_leave(server_id, message) VALUES(%s, %s)"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(join_delete_sql, (app.message.server.id,))
            if not just_delete:
                cursor.execute(join_insert_sql, (app.message.server.id, args,))
            cursor.close()
        app.mysql_connection.commit()
        embed=discord.Embed(title="Leave Message Set:", description="The leave message shows when someone leaves the server.", color=0x00ff00)
        embed.add_field(name="Leave Message Preview:", value=visualised_bye, inline=False)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
# Allows you to set the leave message.

leave_message.description = "Allows you to set the leave message. Use $server$ to repersent the server name and $user$ to repersent the user. if you want to remove the custom greeting, just use $rm$ on its own."
# Sets a description for "leave_message".

leave_message.requires_staff = True
# Set that this script requires staff.

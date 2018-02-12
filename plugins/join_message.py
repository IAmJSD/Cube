import discord
# Imports go here.

async def join_message(app):
    if app.args == []:
        embed=discord.Embed(title="I could not find any arguments.",
                            description="Please supply the join message for arguments (using $user$ to repersent the user and $server$ to repersent the server).",
                            color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        args = ' '.join(app.args)
        if args == "$rm$":
            just_delete = True
            visualised_greeting = "Hello {} and welcome to {}!".format(app.message.author.mention, app.message.server.name)
        else:
            just_delete = False
            visualised_greeting = args.replace("$user$", app.message.author.mention).replace("$server$", app.message.server.name)
        join_delete_sql = "DELETE FROM custom_greetings WHERE server_id = %s"
        join_insert_sql = "INSERT INTO custom_greetings(server_id, message) VALUES(%s, %s)"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(join_delete_sql, (app.message.server.id,))
            if not just_delete:
                cursor.execute(join_insert_sql, (app.message.server.id, args,))
            cursor.close()
        app.mysql_connection.commit()
        embed=discord.Embed(title="Join Message Set:", description="The join message shows when someone joins the server.", color=0x00ff00)
        embed.add_field(name="Join Message Preview:", value=visualised_greeting, inline=False)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
# Allows you to set the join message.

join_message.description = "Allows you to set the join message. Use $server$ to represent the server name and $user$ to represent the user. if you want to remove the custom greeting, just use $rm$ on its own."
# Sets a description for "join_message".

join_message.requires_staff = True
# Set that this script requires staff.

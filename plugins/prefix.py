import discord
# Imports go here.

async def prefix(app):
    args = ' '.join(app.args)
    if args == "":
        embed=discord.Embed(title="What do you want it to be?", 
                    description="I couldn't find any arguments to set as the prefix.",
                    color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        del_sql = "DELETE FROM custom_prefixes WHERE server_id = %s"
        insert_sql = "INSERT INTO custom_prefixes (prefix, server_id) VALUES (%s, %s)"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(del_sql, (app.message.server.id, ))
            if args != "$rm$":
                cursor.execute(insert_sql, (args, app.message.server.id, ))
            cursor.close()
        app.mysql_connection.commit()
        if args == "$rm$":
            args = app.config["bot_name"]
        main_title = "✓ New prefix set."
        main_desc = "`{}` has been set as the new prefix".format(args)
        main_colour = 0x00ff00
        embed=discord.Embed(title=main_title, description=main_desc+".", color=main_colour)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
        embed=discord.Embed(title=main_title, 
        description="{} by {}.".format(main_desc, app.message.author.name), 
        color=main_colour)
        embed.set_footer(text=app.premade_ver)
        await app.attempt_log(app.message.server.id, embed)
# Allows you to set the bots prefix. Use $rm$ to remove the custom prefix.

prefix.description = "Allows you to set the bots prefix. Use $rm$ to remove the custom prefix."
# Sets a description for "prefix".

prefix.requires_management = True
# Set that this script requires management.

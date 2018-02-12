import discord
# Imports go here,

async def botinfo(app):
    embed=discord.Embed(title="This bot is based off the Cube core created by JakeMakesStuff on GitHub.", url="https://github.com/JakeMakesStuff/Cube", color=0x00ff40)
    embed.set_author(name="{}:".format(app.config["bot_name"]))
    '''
    embed.set_thumbnail(url=app.dclient.user.avatar_url)
    Adds SIGNIFICANT lag to the embed generation sadly.
    '''
    embed.add_field(name="Using uvloop:", value=str(app.using_uvloop), inline=True)
    embed.add_field(name="Server Count:", value=len(app.dclient.servers), inline=True)
    bot_owner = app.dclient.get_server(app.config["owner_server_id"]).get_member(app.config["owner_user_id"])
    with app.mysql_connection.cursor() as cursor:
        msg_logs_count_sql = "SELECT COUNT(*) AS COUNT FROM msg_logs"
        cursor.execute(msg_logs_count_sql)
        msg_logs_count = cursor.fetchone()["COUNT"]
        cursor.close()
    embed.add_field(name="Logged Messages:", value=str(msg_logs_count), inline=True)
    embed.add_field(name="Instance Owner:", value="{} [{}]".format(bot_owner.name, bot_owner.id), inline=True)
    embed.add_field(name="Cube Support:", value="https://discord.gg/98wacKS (If you need help with a plugin **NOT natively in Cube**, please ask elsewhere unless you are developing it)", inline=True)
    embed.set_footer(text=app.premade_ver)
    await app.say(embed=embed)
# Gets info about the bot.

botinfo.description = "Gets info about the bot."
# Sets a description for "botinfo".

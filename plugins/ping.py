import discord, time
# Imports go here.

async def ping(app):
    before = time.perf_counter()
    msg = await app.say(embed=discord.Embed(title="Pinging..."))
    total = ((time.perf_counter()-before)*1000)
    embed=discord.Embed(title="🏓 Pong!", description="Generating a message took {} ms.".format(round(total, 1)), color=0x00ff00)
    embed.set_footer(text=app.premade_ver)
    await app.dclient.edit_message(msg, embed=embed)
# Pings the bot.

ping.description = "Pings the bot."
# Sets a description for "ping".

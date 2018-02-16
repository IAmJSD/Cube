import discord
# Imports go here.

async def serverinfo(app):
    embed=discord.Embed(title="Info About {}:".format(app.message.server.name), color=0x1bb3e4)
    embed.set_thumbnail(url=app.message.server.icon_url)
    embed.add_field(name="Server ID:", value=app.message.server.id, inline=False)
    embed.add_field(name="Server Owner:", value=app.message.server.owner.mention, inline=False)
    embed.add_field(name="Member Count:", value=str(len(app.message.server.members)), inline=False)
    embed.add_field(name="Created:", value=str(app.message.server.created_at), inline=False)
    embed.set_footer(text=app.premade_ver)
    await app.say(embed=embed)
# Gets info about the server.

serverinfo.description = "Gets info about the server."
# Sets a description for "serverinfo".

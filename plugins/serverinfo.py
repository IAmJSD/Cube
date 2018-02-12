import discord
# Imports go here.

async def serverinfo(app):
    embed=discord.Embed(title="Info About {}:".format(app.message.server.name))
    embed.set_thumbnail(url=app.message.server.icon_url)
    embed.add_field(name="Server ID:", value=app.message.server.id, inline=True)
    embed.add_field(name="Server Owner:", value=app.message.server.owner.mention, inline=True)
    embed.add_field(name="Member Count:", value=str(len(app.message.server.members)), inline=True)
    embed.add_field(name="Created:", value=str(app.message.server.created_at), inline=True)
    embed.set_footer(text=app.premade_ver)
    await app.say(embed=embed)
# Gets info about the server.

serverinfo.description = "Gets info about the server."
# Sets a description for "serverinfo".

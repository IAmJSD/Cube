import discord
# Imports go here.

async def invite(app):
    try:
        await app.whisper("https://discordapp.com/oauth2/authorize?client_id={}&scope=bot&permissions=2080898303".format(app.dclient.user.id))
        await app.say(embed=discord.Embed(title="📬 Check DM's"))
    except:
        await app.say(embed=discord.Embed(title="Could not DM.", color=0xff0000))
# Allows you to invite to another server.

invite.description = "Allows you to invite to another server."
# Sets a description for "invite".

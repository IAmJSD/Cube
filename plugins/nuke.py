import discord
# Imports go here.

async def nuke(app):
    if len(app.args) == 0:
        embed=discord.Embed(title="I could not find an amount of messages to nuke.",
            description="Please give an amount of messages to nuke as the first argument.",
            color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        t = True
        try:
            amount = int(app.args[0])
        except:
            embed=discord.Embed(title="Could not convert to an integer.",
                                description="Please make sure your first argument is an integer.",
                                color=0xff0000)
            embed.set_footer(text=app.premade_ver)
            await app.say(embed=embed)
            t = False
        if t:
            user = None
            if len(app.args) >= 2:
                user = app.pass_user(app.args[1], app.message.server)
            try:
                def is_user_listed(message):
                    if user is None or user == message.author:
                        return True
                    else:
                        return False
                del_msgs = await app.dclient.purge_from(app.message.channel, limit=amount, check=is_user_listed)
                z = len(del_msgs)
                if not user is None:
                    generic_desc = "{} message(s) created by `{}` were deleted".format(z, user.name)
                else:
                    generic_desc = "{} message(s) were deleted".format(z)
                embed=discord.Embed(title="✓ Messages deleted.",
                                    description="{}.".format(generic_desc),
                                    color=0x00ff00)
                embed.set_footer(text=app.premade_ver)
                await app.say(embed=embed)
                embed=discord.Embed(title="Message(s) deleted.",
                                    description="{} by {}.".format(generic_desc, app.message.author.name))
                embed.set_footer(text=app.premade_ver)
                await app.attempt_log(app.message.server.id, embed)
            except:
                await app.say("```Failed to nuke messages.```")
# Allows you to nuke messages.

nuke.description = "Allows you to nuke messages."
# Sets a description for "nuke".

nuke.requires_staff = True
# Set that this script requires staff.

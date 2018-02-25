# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import discord
# Imports go here.

def Plugin(app):

    @app.command("Gets info about the server.")
    async def serverinfo(app):
        embed=discord.Embed(title=f"Info About {app.message.guild.name}:", color=0x1bb3e4)
        embed.set_thumbnail(url=app.message.guild.icon_url)
        embed.add_field(name="Server ID:", value=str(app.message.guild.id), inline=False)
        embed.add_field(name="Server Owner:", value=app.message.guild.owner.mention, inline=False)
        embed.add_field(name="Member Count:", value=str(len(app.message.guild.members)), inline=False)
        embed.add_field(name="Created:", value=str(app.message.guild.created_at), inline=False)
        embed.set_footer(text=f'{app.config["bot_name"]} v{app.config["version"]}')
        await app.say(embed=embed)
    # Gets info about the server.

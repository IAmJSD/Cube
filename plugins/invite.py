# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import discord
# Imports go here.

def Plugin(app):

    @app.command("DM's a invite for this bot.")
    async def invite(app):
        try:
            await app.whisper("https://discordapp.com/oauth2/authorize?client_id={}&scope=bot&permissions=2080898303".format(app.dclient.user.id))
            await app.say(embed=discord.Embed(title="📬 Check DM's"))
        except:
            await app.say(embed=discord.Embed(title="Could not DM.", color=0xff0000))
    # Allows you to invite to another server.

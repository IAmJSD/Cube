# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import discord, time
# Imports go here.

def Plugin(app):

    @app.command("Pings the bot.")
    async def ping(app):
        before = time.perf_counter()
        msg = await app.say(embed=discord.Embed(title="Pinging..."))
        total = (time.perf_counter()-before)*1000
        await msg.edit(embed=app.create_embed("🏓 Pong!", "Generating a message took {} ms.".format(round(total, 1)), success=True))
    # Pings the bot.

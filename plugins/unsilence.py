# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import discord
# Imports go here.

def Plugin(app):

    @app.command("Allows you to unsilence the join/leave messages from the bot.", requires_management=True)
    async def unsilence(app):
        if (await app.run_mysql("SELECT COUNT(*) AS silenced_count FROM silenced_servers WHERE server_id = %s", (app.message.guild.id, ), get_one=True))[0] == 0:
            await app.say(embed=discord.Embed(title="😃 Not silenced."))
        else:
            await app.run_mysql("DELETE FROM silenced_servers WHERE server_id = %s", (app.message.guild.id, ), commit=True)
            await app.say(embed=discord.Embed(title="😃 Unsilenced."))
    # Allows you to unsilence the join/leave messages from the bot.

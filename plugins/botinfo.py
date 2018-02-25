# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import discord
# Imports go here.

def Plugin(app):

    @app.command("Gets info about the bot.")
    async def botinfo(app):
        embed=discord.Embed(title="This bot is based off the Cube core created by JakeMakesStuff on GitHub.", url="https://github.com/JakeMakesStuff/Cube", color=0x00ff40)
        embed.set_author(name=f'{app.config["bot_name"]}:')
        '''
        embed.set_thumbnail(url=app.dclient.user.avatar_url)
        Adds SIGNIFICANT lag to the embed generation sadly.
        '''
        embed.add_field(name="Using uvloop:", value=str(app.using_uvloop), inline=True)
        embed.add_field(name="Server Count:", value=len(app.dclient.guilds), inline=True)
        bot_owner = app.dclient.get_guild(app.config["owner_server_id"]).get_member(app.config["owner_user_id"])
        msg_logs_count = (await app.run_mysql("SELECT COUNT(*) AS COUNT FROM msg_logs", get_one=True))[0]
        embed.add_field(name="Logged Messages:", value=str(msg_logs_count), inline=True)
        embed.add_field(name="Shard ID:", value=str(app.message.guild.shard_id), inline=False)
        embed.add_field(name="Instance Owner:", value="{} [{}]".format(bot_owner.name, bot_owner.id), inline=True)
        embed.add_field(name="Cube Support:", value="https://discord.gg/98wacKS (If you need help with a plugin **NOT natively in Cube**, please ask elsewhere unless you are developing it)", inline=True)
        embed.set_footer(text=f'{app.config["bot_name"]} v{app.config["version"]}')
        await app.say(embed=embed)
    # Gets info about the bot.

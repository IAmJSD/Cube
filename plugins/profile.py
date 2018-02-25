# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import discord
# Imports go here.

def Plugin(app):

    @app.command("Allows you to get information about a Discord account.", usage="[user] (Optional)")
    async def profile(app):
        if len(app.args) != 0 and app.pass_user(app.message.guild, app.args[0]) is None:
                await app.say(embed=app.create_embed("Could not find the user.", "Make sure you tag the user as your first argument or leave it blank.", error=True))
        else:
            if app.args == []:
                user = app.message.author
            else:
                user = app.pass_user(app.message.guild, app.args[0])
            embed=discord.Embed(title=user.name + "'s Profile:", colour=user.colour)
            embed.add_field(name="Discriminator:", value=str(user.discriminator), inline=True)
            embed.add_field(name="ID:", value=str(user.id), inline=True)
            embed.add_field(name="Bot Account:", value=str(user.bot), inline=True)
            embed.add_field(name="Highest Role:", value=user.top_role.name, inline=True)
            embed.add_field(name="User:", value=user.mention, inline=True)
            embed.add_field(name="Game:", value=str(user.game), inline=True)
            embed.add_field(name="Created:", value=str(user.created_at).split(' ')[0], inline=True)
            embed.add_field(name="Joined:", value=str(user.joined_at).split(' ')[0], inline=True)
            strike_count = (await app.run_mysql("SELECT COUNT(*) AS COUNT FROM strikes WHERE user_id = %s AND server_id = %s",
                                                                        (user.id, app.message.guild.id, ), get_one=True))[0]
            merit_count = (await app.run_mysql("SELECT COUNT(*) AS COUNT FROM merits WHERE user_id = %s AND server_id = %s",
                                                                        (user.id, app.message.guild.id, ), get_one=True))[0]
            kick_count = (await app.run_mysql("SELECT COUNT(*) AS COUNT FROM kicks WHERE user_id = %s AND server_id = %s",
                                                                    (user.id, app.message.guild.id, ), get_one=True))[0]
            ban_count = (await app.run_mysql("SELECT COUNT(*) AS COUNT FROM bans WHERE user_id = %s AND server_id = %s",
                                                                    (user.id, app.message.guild.id, ), get_one=True))[0]
            embed.add_field(name="Strikes:", value=str(strike_count), inline=True)
            embed.add_field(name="Merits:", value=str(merit_count), inline=True)
            embed.add_field(name="Kicks:", value=str(kick_count), inline=True)
            embed.add_field(name="Bans:", value=str(ban_count), inline=True)
            embed.set_footer(text=f'{app.config["bot_name"]} v{app.config["version"]}')
            await app.say(embed=embed)
    # Allows you to get information about a Discord account.

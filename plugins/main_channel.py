# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

def Plugin(app):

    @app.command("Allows you to set a main channel (Used for things like join/leave messages).", requires_management=True, usage="[channel]")
    async def main_channel(app):
        if app.args == []:
            await app.say(embed=app.create_embed("Could not find arguments.",
                                "Please provide arguments for this command.",
                                                                error=True))
        else:
            channel = app.pass_channel(app.args[0], app.message.guild)
            if channel is None:
                await app.say(embed=app.create_embed("Could not find the channel.",
                            "Make sure you tag the channel as your first argument.",
                                                                        error=True))
            else:
                try:
                    test_msg = await channel.send("This is a test of the main channel.")
                    await test_msg.delete()
                except:
                    await app.say(embed=app.create_embed("Could not send a test message.",
                        "Make sure the bot has read and write permssions to the channel.",
                                                                            error=True))
                    return
                await app.run_mysql("DELETE FROM main_channels WHERE server_id = %s", (app.message.guild.id, ), commit=True)
                await app.run_mysql("INSERT INTO main_channels (server_id, channel_id) VALUES (%s, %s)",
                                                    (app.message.guild.id, channel.id, ), commit=True)
                embed = app.create_embed("Main channel set:",
                f"{app.message.author.mention} set the main channel to {channel.mention}.",
                success=True)
                await app.say(embed=embed)
                await app.attempt_log(app.message.guild.id, embed=embed)
    # Allows you to set a main channel (Used for things like join/leave messages).

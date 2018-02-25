# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import asyncio
# Imports go here.

def Plugin(app):

    @app.command("Allows you to kick a user.", requires_staff=True, usage="[user] [reason]")
    async def kick(app):
        if app.args == []:
            await app.say(embed=app.create_embed("Could not find arguments.",
                                "Please provide arguments for this command.",
                                                                error=True))
        else:
            user = app.pass_user(app.message.guild, app.args[0])
            if user is None:
                await app.say(embed=app.create_embed("Could not find the user.",
                            "Make sure you tag the user as your first argument.",
                                                                    error=True))
            else:
                if len(app.args) == 1:
                    reason = "No reason found."
                else:
                    reason = ' '.join(app.args[1:])
                await app.say(f"Are you sure you want to kick `{user.name}` for `{reason}`? Please say yes in order to run this action.")
                def check(m):
                    return m.author == app.message.author and m.channel == app.message.channel
                try:
                    msg2 = await app.dclient.wait_for('message', check=check, timeout=30)
                    if msg2.content.strip(' ').lower() == "yes":
                        x = False
                        try:
                            try:
                                await user.send(f"You were kicked from `{app.message.guild.name}` by `{app.message.author.name} ({app.message.author.id})` for `{reason}`.")
                            except:
                                pass
                            await user.kick(reason=reason)
                            x = True
                        except:
                            pass
                        if x:
                            e = app.create_embed("User kicked:",
                            f"`{user.name} ({user.id})` was kicked by {app.message.author.mention} for `{reason}`..",
                            success=True)
                            await app.say(embed=e)
                            await app.run_mysql("INSERT INTO `kicks`(`server_id`, `staff_id`, `user_id`, `reason`) VALUES(%s, %s, %s, %s)",
                            (app.message.guild.id, app.message.author.id, user.id, reason, ), commit=True)
                            await app.attempt_log(app.message.guild.id, embed=e)
                        else:
                            await app.say("Could not kick the specified user.")
                    else:
                        await app.say("Kick cancelled.")
                except asyncio.TimeoutError:
                    pass
    # Allows you to kick a user.

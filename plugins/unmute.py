# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import asyncio
# Imports go here.

def Plugin(app):

    @app.command('Allows you to unmute a user.', requires_staff=True, usage="[user]")
    async def unmute(app):
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
                if not "muted" in str([r.name.lower() for r in user.roles]):
                    await app.say("User not muted.")
                else:
                    m = None
                    for r in user.roles:
                        if m is None and "muted" in r.name.lower():
                            m = r
                    try:
                        await user.remove_roles(m, reason="UNMUTE")
                    except:
                        await app.say("Could not remove the Muted role.")
                        return
                    try:
                        await user.send(f"You were unmuted in `{app.message.guild.name}` by `{app.message.author.name} ({app.message.author.id})`.")
                    except:
                        pass
                    embed = app.create_embed("User unmuted:", f"`{user.name} ({user.id})` was unmuted by {app.message.author.mention}.", success=True)
                    await app.say(embed=embed)
                    await app.attempt_log(app.message.guild.id, embed=embed)
    # Allows you to unmute a user.

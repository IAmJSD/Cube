# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import asyncio, discord
# Imports go here.

def Plugin(app):

    @app.command("Allows you to ban a user by using their ID.", requires_staff=True, usage="[user_id] [reason]")
    async def hackban(app):
        if app.args == []:
            await app.say(embed=app.create_embed("Could not find arguments.",
                                "Please provide arguments for this command.",
                                                                error=True))
        else:
            if len(app.args) == 1:
                reason = "No reason found."
            else:
                reason = ' '.join(app.args[1:])
            await app.say(f"Are you sure you want to ban a user with the ID `{app.args[0]}` for `{reason}`? Please say yes in order to run this action.")
            def check(m):
                return m.author == app.message.author and m.channel == app.message.channel
            try:
                msg2 = await app.dclient.wait_for('message', check=check, timeout=30)
                if msg2.content.strip(' ').lower() == "yes":
                    x = False
                    try:
                        user = discord.Object(id=int(app.args[0]))
                        await app.message.guild.ban(user, reason=reason)
                        x = True
                    except:
                        pass
                    if x:
                        e = app.create_embed("User hackbanned:",
                        f"A user with the ID `{app.args[0]}` was hackbanned by {app.message.author.mention} for `{reason}`.",
                        success=True)
                        await app.say(embed=e)
                        await app.run_mysql("INSERT INTO `bans`(`server_id`, `staff_id`, `user_id`, `reason`) VALUES(%s, %s, %s, %s)",
                        (app.message.guild.id, app.message.author.id, int(app.args[0]), reason, ), commit=True)
                        await app.attempt_log(app.message.guild.id, embed=e)
                    else:
                        await app.say("Could not ban the specified user.")
                else:
                    await app.say("ban cancelled.")
            except asyncio.TimeoutError:
                pass
    # Allows you to ban a user by using their ID.

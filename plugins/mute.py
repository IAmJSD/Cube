# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import asyncio
# Imports go here.

def Plugin(app):

    @app.command('Allows you to mute a user. Requires a role with "muted" in (not case sensitive).', requires_staff=True, usage="[user]")
    async def mute(app):
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
                await app.say(f"Are you sure you want to mute `{user.name}` for `{reason}`? Please say yes in order to run this action.")
                def check(m):
                    return m.author == app.message.author and m.channel == app.message.channel
                try:
                    msg2 = await app.dclient.wait_for('message', check=check, timeout=30)
                    if msg2.content.strip(' ').lower() == "aye":
                        if "muted" in str([r.name.lower() for r in user.roles]):
                            await app.say("User already muted.")
                        else:
                            try:
                                roles_without_everyone = [r for r in user.roles if r.name != "@everyone"]
                                await user.remove_roles(*roles_without_everyone, reason=f"MUTE: {reason}")
                            except:
                                await app.say("Cannot remove roles. Does the bot have permission to do so?")
                                return
                            m = None
                            for r in app.message.guild.roles:
                                if m is None and "muted" in r.name.lower():
                                    m = r
                            if m is None:
                                await app.say("No Muted role found. Please create one for this command to work.")
                                return
                            try:
                                await user.add_roles(m, reason=f"MUTE: {reason}")
                            except:
                                await app.say("Could not apply the Muted role.")
                                return
                            embed = app.create_embed("User muted:", f"`{user.name} ({user.id})` was muted by {app.message.author.mention} for `{reason}`.", success=True)
                            try:
                                await user.send(f"You were muted in `{app.message.guild.name}` by `{app.message.author.name} ({app.message.author.id})` for `{reason}`.")
                            except:
                                pass
                            await app.say(embed=embed)
                            await app.attempt_log(app.message.guild.id, embed=embed)
                    else:
                        await app.say("Mute cancelled.")
                except asyncio.TimeoutError:
                    await app.say("Mute cancelled.")
    # Allows you to mute a user. Requires a role with "muted" in (not case sensitive).

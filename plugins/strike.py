# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import random
# Imports go here.

def Plugin(app):

    @app.command("Allows you to strike a user.", requires_staff=True, usage="[user] [reason]")
    async def strike(app):
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
                x = True
                while x:
                    charset = "0123456789"
                    strike_id_length = 10
                    strike_id = ""
                    while not len(strike_id) == strike_id_length:
                        strike_id = strike_id + charset[random.randint(0, len(charset)-1)]
                        if (await app.run_mysql("SELECT COUNT(*) AS COUNT FROM strikes WHERE strike_id = %s", (strike_id,), get_one=True))[0] == 0:
                            x = False
                await app.run_mysql("INSERT INTO strikes(strike_id, staff_id, user_id, server_id, strike_reason) VALUES (%s, %s, %s, %s, %s)",
                (strike_id, app.message.author.id, user.id, app.message.guild.id, reason, ), commit=True)
                strike_count = (await app.run_mysql("SELECT COUNT(*) AS COUNT FROM strikes WHERE user_id = %s AND server_id = %s", (user.id, app.message.guild.id, ), get_one=True))[0]
                try:
                    await user.send(f"You were stricken in `{app.message.guild.name}` by `{app.message.author.name} ({app.message.author.id})` for `{reason}`.")
                except:
                    pass
                embed = app.create_embed("User stricken:",
                                        f"`{user.name} ({user.id})` was stricken by {app.message.author.mention} for `{reason}`. They now have {strike_count} strike(s).",
                                        success=True)
                await app.say(embed=embed)
                await app.attempt_log(app.message.guild.id, embed=embed)
    # Allows you to strike a user.

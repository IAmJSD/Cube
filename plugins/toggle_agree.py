# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import asyncio
# Imports go here.

def Plugin(app):

    @app.command('Allows the "agree" command to be turned on and off. If you want to reconfigure "agree", simply run it twice.', requires_management=True)
    async def toggle_agree(app):
        is_agree_active = not (await app.run_mysql("SELECT COUNT(*) AS COUNT FROM member_roles WHERE server_id = %s", (app.message.guild.id, ), get_one=True))[0] == 0
        if is_agree_active:
            await app.run_mysql("DELETE FROM member_roles WHERE server_id = %s", (app.message.guild.id, ), commit=True)
            embed = app.create_embed("Agree disabled:", f"The agree function was disabled by {app.message.author.mention}.", error=True)
            await app.say(embed=embed)
            await app.attempt_log(app.message.guild.id, embed=embed)
        else:
            await app.say("Please say something that is in the name of the Member role **but not in any other role**. This is not case sensitive.")
            def check(m):
                return m.author == app.message.author and m.channel == app.message.channel
            try:
                msg2 = await app.dclient.wait_for('message', check=check, timeout=30)
            except asyncio.TimeoutError:
                return
            await app.run_mysql("INSERT INTO member_roles (server_id, role_part) VALUES(%s, %s)", (app.message.guild.id, msg2.content, ), commit=True)
            embed = app.create_embed("Agree enabled:", f"The agree function was enabled by {app.message.author.mention} to apply a role which contains `{msg2.content}`.", success=True)
            await app.say(embed=embed)
            await app.attempt_log(app.message.guild.id, embed=embed)
    # Allows the "agree" command to be turned on and off. If you want to reconfigure "agree", simply run it twice.

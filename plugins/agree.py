# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

def Plugin(app):

    @app.command("If the agree system is configured properly, this will give the Member role.")
    async def agree(app):
        try:
            await app.message.delete()
        except:
            pass
        member_role_part = await app.run_mysql("SELECT * FROM member_roles WHERE server_id = %s", (app.message.guild.id, ), get_one=True)
        if member_role_part is None:
            try:
                await app.whisper("The agree command is not enabled in this server.")
            except:
                pass
        else:
            member_role_part = member_role_part[1]
            x = False
            if "muted" in str([r.name.lower() for r in app.message.author.roles]):
                await app.whisper("You are muted!")
                return
            for role in app.message.guild.roles:
                if not x:
                    if member_role_part in role.name.lower():
                        try:
                            await app.message.author.add_roles(role)
                            x = True
                        except:
                            pass
            try:
                if x:
                    await app.whisper(f"You were granted the Member role in {app.message.guild.name}.")
                else:
                    await app.whisper(f"I could not give you the Member role in {app.message.guild.name}. This could be either because the bot does not have permission or the role is not found.")
            except:
                pass
    # If the agree system is configured properly, this will give the Member role.

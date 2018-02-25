# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import discord
# Imports go here.

def Plugin(app):

    @app.command("Gets help for the bot.")
    async def help(app):
        prefix = await app.get_prefix(app.message.guild.id)
        is_management = False
        for role in app.message.author.roles:
            if "manager" in role.name.lower() or "management" in role.name.lower():
                is_management = True
        is_staff = False
        for role in app.message.author.roles:
            if "staff" in role.name.lower():
                is_staff = True
        bot_admin_check_sql = await app.run_mysql("SELECT * FROM `bot_admins` WHERE `user_id` = %s", (app.message.author.id,), get_one=True)
        is_bot_admin = not bot_admin_check_sql is None
        embeds = [discord.Embed(title="Cube Help:")]
        embed_fields = 0
        cmd_attrs = []
        for cmd in app.discord_commands:
            a = app.discord_commands[cmd].cmd_attr
            a["command"] = cmd
            cmd_attrs.append(a)
        cmd_attrs.sort(key=lambda y: y["command"])
        for a in cmd_attrs:
            show = True
            if (a["requires_staff"] and not is_staff) or (a["requires_management"] and not is_management) or (a["requires_bot_admin"] and not is_bot_admin):
                show = False
            if show:
                if embed_fields == 10:
                    embeds.append(discord.Embed())
                    embed_fields = 0
                if not a["usage"] is None:
                    cmd_name = f"{prefix}{a['command']} {a['usage']}"
                else:
                    cmd_name = f"{prefix}{a['command']}"
                embeds[-1].add_field(name=cmd_name, value=a["description"], inline=False)
                embed_fields = embed_fields + 1
        try:
            for y in embeds:
                await app.whisper(embed=y)
            await app.say(embed=discord.Embed(title="📬 Check DM's"))
        except:
            await app.say(embed=discord.Embed(title="Could not DM.", color=0xff0000))
    # Allows you to get help.

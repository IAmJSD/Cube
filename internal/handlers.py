# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

from copy import copy
import discord, asyncio
# Imports go here.

async def no_perm_handler(app, msg):
    await msg.channel.send(embed=app.create_embed("Uh oh.", "You do not have permission to do this.", error=True))
# Handles when the user does not have permission to do something.

async def exception_handler(app, msg, e):
    await msg.channel.send(embed=app.create_embed("Bug alert!",
    f"I had a problem executing that command.```{e}```Simply reply `yes` in the next 30 seconds for this to be reported or anything else to ignore this.",
    error=True))
    def check(m):
        return m.author == msg.author and m.channel == msg.channel
    try:
        msg2 = await app.dclient.wait_for('message', check=check, timeout=30)
        if msg2.content.lower() == "yes":
            owner = app.dclient.get_guild(app.config["owner_server_id"]).get_member(app.config["owner_user_id"])
            formatted_msg = "New command bug report!\n\nMessage that triggered: `{}`\n\nError: `{}`".format(msg.content, e)
            await owner.send(formatted_msg)
            await msg.channel.send("Bug report submitted! Thanks for making Cube better.")
    except asyncio.TimeoutError:
        pass
# The exception handler for the bot.

async def cmd_handler(app, msg):
    p = await app.get_prefix(msg.guild.id)
    p_len = len(p)
    if msg.content[:p_len] == p:
        msg_p_strip_split = [x for x in msg.content[p_len:].split(' ') if x != ""]
        cmd = msg_p_strip_split[0].lower()
        args = msg_p_strip_split[1:]
        if cmd in app.discord_commands:
            c = app.discord_commands[cmd]
            a = c.cmd_attr
            if (a["requires_management"] and not "manager" in [r.name.lower() for r in msg.author.roles] and not "management" in [r.name.lower() for r in msg.author.roles]) or (a["requires_staff"] and not "staff" in [r.name.lower() for r in msg.author.roles]):
                return await no_perm_handler(app, msg)
            elif a["requires_bot_admin"]:
                bot_admin_check_sql = await app.run_mysql("SELECT * FROM `bot_admins` WHERE `user_id` = %s", (msg.author.id,), get_one=True)
                if bot_admin_check_sql is None:
                    return await no_perm_handler(app, msg)
            try:
                app_ctx = copy(app)
                app_ctx.message = msg
                app_ctx.args = args
                await c(app_ctx)
            except Exception as e:
                await exception_handler(app, msg, e)
        else:
            custom_cmd = await app.run_mysql("SELECT * FROM `custom_commands` WHERE `server_id` = %s AND `command` = %s", (msg.guild.id, cmd,), get_one=True)
            if not custom_cmd is None:
                try:
                    await msg.channel.send(custom_cmd[2])
                except:
                    pass
# The command handler for the bot.

async def server_restriction_handler(server):
    owner = server.owner
    members = 0
    bots = 0
    for m in server.members:
        members = members + 1
        if m.bot:
            bots = bots + 1
    percentage_bots = 100 * float(bots) / float(members)
    if percentage_bots >= 75 and members >= 100:
        try:
            await owner.send("This bot will not work on servers that are >75% bots and have >100 members.")
        except:
            pass
        await server.leave()
        return True
    else:
        return False
# Handles server spam prevention.

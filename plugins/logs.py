# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

import discord, os
# Imports go here.

def Plugin(app):

    @app.event
    async def on_message(app):
        await app.run_mysql("INSERT INTO `msg_logs` (`message_id`, `user_id`, `server_id`, `channel_id`, `attachments`, `content`) VALUES (%s, %s, %s, %s, %s, %s)",
        (app.message.id, app.message.author.id, app.message.guild.id, app.message.channel.id, str(app.message.attachments), app.message.content, ), commit=True)
    # Logs each message.

    @app.command("Allows you to get logs about a user.", requires_staff=True, usage="[user]")
    async def logs(app):
        if app.args == []:
            user = app.message.author
        else:
            user = app.pass_user(app.message.guild, app.args[0])
        if user is None:
            await app.say(embed=app.create_embed("Could not find the user.",
                        "Make sure you tag the user as your first argument.",
                                                                error=True))
        else:
            line = "-----------"
            log = line+"\nMessage logs for {} ({})\n".format(user.name, user.id)+line
            msgs = await app.run_mysql("SELECT * FROM msg_logs WHERE user_id = %s AND server_id = %s",
                                                    (user.id, app.message.guild.id, ), get_many=True)
            for msg in msgs:
                channel = app.pass_channel(msg[3])
                if channel is None:
                    channel_layout = "{}".format(msg[3])
                else:
                    channel_layout = "{} ({})".format(channel.name, msg[3])
                log=log+"\nMessage ID: {}\nChannel: {}\nAttachments: {}\nMessage Content: {}\n{}".format(msg[0], channel_layout, msg[4], msg[5], line)

            with open(f"{user.id}.txt", "w+", encoding="utf8") as logfile:
                logfile.write(log)

            try:
                x = True
                await app.whisper(file=discord.File(fp=f"{user.id}.txt"))
            except:
                x = False

            os.remove(f"{user.id}.txt")

            if x:
                await app.say(embed=discord.Embed(title="📬 Check DM's"))
            else:
                await app.say(embed=discord.Embed(title="Could not DM.", color=0xff0000))
    # Allows you to get logs about a user.

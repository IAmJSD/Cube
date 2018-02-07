import discord, os
# Imports go here.

async def on_message(app):
    msg = app.message
    if not msg.channel.is_private:
        sql = "INSERT INTO `msg_logs` (`message_id`, `user_id`, `server_id`, `channel_id`, `attachments`, `content`) VALUES (%s, %s, %s, %s, %s, %s)"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(sql, (msg.id, msg.author.id, msg.server.id, msg.channel.id, str(msg.attachments), msg.content, ))
            cursor.close()
        app.mysql_connection.commit()
# Logs every message sent.

async def logs(app):
    if app.args == []:
        user = app.message.author
    else:
        user = app.pass_user(app.args[0], app.message.server)
    if user == None:
        embed=discord.Embed(title="Huston, we have a problem.",
            description="I couldn't find that user.",
            color=0xff0000)
        embed.set_footer(text=app.premade_ver)
        await app.say(embed=embed)
    else:
        line = "-----------"
        log = line+"\nMessage logs for {} ({})\n".format(user.name, user.id)+line
        sql = "SELECT * FROM msg_logs WHERE user_id = %s AND server_id = %s"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(sql, (user.id, app.message.server.id, ))
            msgs=cursor.fetchall()
            for msg in msgs:
                channel = app.pass_channel(msg["channel_id"])
                if channel == None:
                    channel_layout = "{}".format(msg["channel_id"])
                else:
                    channel_layout = "{} ({})".format(channel.name, msg["channel_id"])
                log=log+"\nMessage ID: {}\nChannel: {}\nAttachments: {}\nMessage Content: {}\n{}".format(msg["message_id"], channel_layout, msg["attachments"], msg["content"], line)
            cursor.close()
        with open(user.id+".txt", "w+", encoding="utf8") as logfile:
            logfile.write(log)
        try:
            x = True
            await app.dclient.send_file(app.message.author, user.id+".txt")
        except:
            x = False
        os.remove(user.id+".txt")
        if x:
            await app.say(embed=discord.Embed(title="📬 Check DM's"))
        else:
            await app.say(embed=discord.Embed(title="Could not DM.", color=0xff0000))
# Allows you to check logs of a user.

logs.description = "Allows you to check logs of a user."
# Sets a description for "logs".

logs.requires_staff = True
# Set that this script requires staff.

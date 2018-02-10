import discord
# Imports go here.

async def on_member_join(app, member):
    silenced_check_sql = "SELECT COUNT(*) AS silenced_count FROM silenced_servers WHERE server_id = %s"
    with app.mysql_connection.cursor() as cursor:
        cursor.execute(silenced_check_sql, (member.server.id,))
        if cursor.fetchone()["silenced_count"] == 0:
            main_channel_count_sql = "SELECT COUNT(*) AS main_channel_count FROM main_channels WHERE server_id = %s"
            cursor.execute(main_channel_count_sql, (member.server.id,))
            main_channel_count = cursor.fetchone()["main_channel_count"]
            if main_channel_count != 0:
                main_channel_sql = "SELECT * FROM main_channels WHERE server_id = %s"
                greeting_sql = "SELECT * FROM custom_greetings WHERE server_id = %s"
                greeting_count_sql = "SELECT COUNT(*) AS greeting_count FROM custom_greetings WHERE server_id = %s"
                cursor.execute(main_channel_sql, (member.server.id,))
                main_channel = cursor.fetchone()["channel_id"]
                cursor.execute(greeting_count_sql, (member.server.id,))
                if cursor.fetchone()["greeting_count"] != 0:
                    cursor.execute(greeting_sql, (member.server.id,))
                    greeting = cursor.fetchone()["message"].replace("$user$", member.mention).replace("$server$", member.server.name)
                else:
                    greeting = "Hello {} and welcome to {}!".format(member.mention, member.server.name)
                embed = discord.Embed(title="👋 Welcome!", description=greeting, color=0x00ff00)
                embed.set_footer(text=app.premade_ver)
            cursor.close()
        try:
            await app.dclient.send_message(app.dclient.get_channel(main_channel), embed=embed)
        except:
            pass
# Used to welcome members into a server.

async def on_member_remove(app, member):
    silenced_check_sql = "SELECT COUNT(*) AS silenced_count FROM silenced_servers WHERE server_id = %s"
    with app.mysql_connection.cursor() as cursor:
        cursor.execute(silenced_check_sql, (member.server.id,))
        if cursor.fetchone()["silenced_count"] == 0:
            main_channel_count_sql = "SELECT COUNT(*) AS main_channel_count FROM main_channels WHERE server_id = %s"
            cursor.execute(main_channel_count_sql, (member.server.id,))
            main_channel_count = cursor.fetchone()["main_channel_count"]
            if main_channel_count != 0:
                main_channel_sql = "SELECT * FROM main_channels WHERE server_id = %s"
                leave_message_sql = "SELECT * FROM custom_leave WHERE server_id = %s"
                leave_message_count_sql = "SELECT COUNT(*) AS leave_message_count FROM custom_leave WHERE server_id = %s"
                cursor.execute(main_channel_sql, (member.server.id,))
                main_channel = cursor.fetchone()["channel_id"]
                cursor.execute(leave_message_count_sql, (member.server.id,))
                if cursor.fetchone()["leave_message_count"] != 0:
                    cursor.execute(leave_message_sql, (member.server.id,))
                    bye = cursor.fetchone()["message"].replace("$user$", member.name).replace("$server$", member.server.name)
                else:
                    bye = "{} has left this server.".format(member.name)
                embed = discord.Embed(title="User left.", description=bye, color=0xff0000)
                embed.set_footer(text=app.premade_ver)
            cursor.close()
        try:
            await app.dclient.send_message(app.dclient.get_channel(main_channel), embed=embed)
        except:
            pass
# Used to say bye to members.

import discord
# Imports go here.

async def toggle_agree(app):
    member_name_count_sql = "SELECT COUNT(*) AS COUNT FROM member_roles WHERE server_id = %s"
    with app.mysql_connection.cursor() as cursor:
        cursor.execute(member_name_count_sql, (app.message.server.id, ))
        member_name_count = cursor.fetchone()["COUNT"]
        cursor.close()
    if member_name_count == 0:
        await app.say("Please say something that is in the name of the Member role **but not in any other role**. This is not case sensitive.")
        msg2 = await app.dclient.wait_for_message(author=app.message.author, timeout=30)
        if not msg2 is None:
            role = msg2.content.lower()
            member_name_insert_sql = "INSERT INTO member_roles (part_of_role, server_id) VALUES(%s, %s)"
            with app.mysql_connection.cursor() as cursor:
                cursor.execute(member_name_insert_sql, (role, app.message.server.id, ))
                cursor.close()
            await app.say("Alright! The agree command is active now.")
            log_embed = discord.Embed(title="Agree setup:", description="The agree command was setup by `{}` to give a role where part of the role name was `{}`.".format(app.message.author.name, role))
            await app.attempt_log(app.message.server.id, log_embed)
    else:
            disable_agree_sql = "DELETE FROM member_roles WHERE server_id = %s"
            with app.mysql_connection.cursor() as cursor:
                cursor.execute(disable_agree_sql, (app.message.server.id, ))
                cursor.close()
            await app.say("The agree command was disabled.")
            log_embed = discord.Embed(title="Agree disabled:", description="The agree command was disabled by `{}`.".format(app.message.author.name))
            await app.attempt_log(app.message.server.id, log_embed)
# Allows the "agree" command to be turned on and off.

toggle_agree.description = 'Allows the "agree" command to be turned on and off.'
# Sets a description for "toggle_agree".

toggle_agree.requires_staff = True
# Set that this script requires staff.

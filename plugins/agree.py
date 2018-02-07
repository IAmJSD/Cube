async def agree(app):
    member_name_count_sql = "SELECT COUNT(*) AS COUNT FROM member_roles WHERE server_id = %s"
    with app.mysql_connection.cursor() as cursor:
        cursor.execute(member_name_count_sql, (app.message.server.id, ))
        member_name_count = cursor.fetchone()["COUNT"]
        cursor.close()
    if member_name_count == 0:
        await app.whisper("The agree command is not enabled in this server.")
    else:
        get_role = "SELECT * FROM member_roles WHERE server_id = %s"
        with app.mysql_connection.cursor() as cursor:
            cursor.execute(get_role, (app.message.server.id, ))
            role_part = cursor.fetchone()["part_of_role"]
            cursor.close()
        x = False
        for role in app.message.server.roles:
            if not x:
                if role_part in role.name.lower():
                    try:
                        await app.dclient.add_roles(app.message.author, role)
                        x = True
                    except:
                        pass
        try:
            if x:
                await app.whisper("You were granted the Member role in {}.".format(app.message.server.name))
            else:
                await app.whisper("I could not give you the Member role in {}. This could be either because the bot does not have permission or the role is not found.".format(app.message.server.name))
        except:
            pass
    try:
        await app.dclient.delete_message(app.message)
    except:
        pass
# If the agree system is configured properly, this will give the Member role.

agree.description = "If the agree system is configured properly, this will give the Member role."
# Sets a description for "agree".

# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

def Plugin(app):

    @app.event
    async def on_member_join(app, member):
        if (await app.run_mysql("SELECT COUNT(*) AS silenced_count FROM silenced_servers WHERE server_id = %s", (member.guild.id, ), get_one=True))[0] == 0:
            main_channel = await app.run_mysql("SELECT * FROM main_channels WHERE server_id = %s", (member.guild.id,), get_one=True)
            if not main_channel is None:
                main_channel = main_channel[1]
                greeting = await app.run_mysql("SELECT * FROM custom_greetings WHERE server_id = %s", (member.guild.id, ), get_one=True)
                if not greeting is None:
                    greeting = greeting[1].replace("$user$", member.mention).replace("$server$", member.guild.name)
                else:
                    greeting = "Hello {} and welcome to {}!".format(member.mention, member.guild.name)
                embed = app.create_embed("👋 Welcome!", greeting, success=True)
                try:
                    await app.pass_channel(main_channel).send(embed=embed)
                except:
                    pass
    # Used to welcome members into a server.

    @app.event
    async def on_member_remove(app, member):
        if (await app.run_mysql("SELECT COUNT(*) AS silenced_count FROM silenced_servers WHERE server_id = %s", (member.guild.id, ), get_one=True))[0] == 0:
            main_channel = await app.run_mysql("SELECT * FROM main_channels WHERE server_id = %s", (member.guild.id, ), get_one=True)
            if not main_channel is None:
                main_channel = main_channel[1]
                leave_message = await app.run_mysql("SELECT * FROM custom_leave WHERE server_id = %s", (member.guild.id, ), get_one=True)
                if not leave_message is None:
                    bye = leave_message[1].replace("$user$", member.name).replace("$server$", member.guild.name)
                else:
                    bye = "{} has left this server.".format(member.name)
                embed = app.create_embed("User left.", bye, error=True)
                try:
                    await app.pass_channel(main_channel).send(embed=embed)
                except:
                    pass
    # Used to say bye to members.

    @app.command("Allows you to set the join message.", requires_management=True, usage="[join_message] (using $user$ to repersent the user, $server$ to repersent the server or just $rm$ on its own to delete the message)")
    async def join_message(app):
        if app.args == []:
            await app.say(embed=app.create_embed("Could not find arguments.",
                                "Please provide arguments for this command.",
                                                                error=True))
        else:
            args = ' '.join(app.args)
            if args == "$rm$":
                just_delete = True
                visualised_greeting = "Hello {} and welcome to {}!".format(app.message.author.mention, app.message.guild.name)
            else:
                just_delete = False
                visualised_greeting = args.replace("$user$", app.message.author.mention).replace("$server$", app.message.guild.name)
            await app.run_mysql("DELETE FROM custom_greetings WHERE server_id = %s", (app.message.guild.id, ), commit=True)
            if not just_delete:
                await app.run_mysql("INSERT INTO custom_greetings(server_id, greeting) VALUES(%s, %s)", (app.message.guild.id, args, ), commit=True)
            embed=app.create_embed("Join Message Set:", f"The join message shows when someone joins the server. This was set by {app.message.author.mention}.", success=True)
            embed.add_field(name="Join Message Preview:", value=visualised_greeting, inline=False)
            await app.say(embed=embed)
            await app.attempt_log(app.message.guild.id, embed=embed)
    # Allows you to set the join message.

    @app.command("Allows you to set the leave message.", requires_management=True, usage="[leave_message] (using $user$ to repersent the user, $server$ to repersent the server or just $rm$ on its own to delete the message)")
    async def leave_message(app):
        if app.args == []:
            await app.say(embed=app.create_embed("Could not find arguments.",
                                "Please provide arguments for this command.",
                                                                error=True))
        else:
            args = ' '.join(app.args)
            if args == "$rm$":
                just_delete = True
                visualised_bye = "{} has left this server.".format(app.message.author.name)
            else:
                just_delete = False
                visualised_bye = args.replace("$user$", app.message.author.name).replace("$server$", app.message.guild.name)
            await app.run_mysql("DELETE FROM custom_leave WHERE server_id = %s", (app.message.guild.id, ), commit=True)
            if not just_delete:
                await app.run_mysql("INSERT INTO custom_leave(server_id, bye) VALUES(%s, %s)", (app.message.guild.id, args, ), commit=True)
            embed=app.create_embed("Leave Message Set:", f"The leave message shows when someone leaves the server. This was set by {app.message.author.mention}.", success=True)
            embed.add_field(name="Leave Message Preview:", value=visualised_bye, inline=False)
            await app.say(embed=embed)
            await app.attempt_log(app.message.guild.id, embed=embed)
    # Allows you to set the leave message.

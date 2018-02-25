# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

def Plugin(app):

    @app.command("Allows you to create a custom command.", requires_management=True, usage="[command] [response]")
    async def add_command(app):
        if app.args == []:
            await app.say(embed=app.create_embed("Could not find arguments.",
                                "Please provide arguments for this command.",
                                                                error=True))
        else:
            if len(app.args) == 1:
                await app.say(embed=app.create_embed("No response found.",
                    "Please provide the response for the second argument.",
                                                                error=True))
            else:
                command = app.args[0]
                response = ' '.join(app.args[1:])
                cmd_count = (await app.run_mysql("SELECT COUNT(*) AS COUNT FROM custom_commands WHERE server_id = %s AND command = %s",
                            (app.message.guild.id, command, ), get_one=True))[0]
                if cmd_count != 0:
                    await app.say(embed=app.create_embed("Command already exists.",
                                f"`{command}` already exists. Please delete it if you want to recreate it.",
                                error=True))
                else:
                    await app.run_mysql("INSERT INTO custom_commands (server_id, command, response) VALUES (%s, %s, %s)",
                                                                (app.message.guild.id, command, response, ), commit=True)
                    embed = app.create_embed("Custom command added:",
                    f"`{command}` was added as a custom command by {app.message.author.mention}.", success=True)
                    await app.say(embed=embed)
                    await app.attempt_log(app.message.guild.id, embed=embed)
    # Allows you to create a custom command.

# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

def Plugin(app):

    @app.command("Allows you to set the prefix.", requires_management=True, usage="[prefix] (or $rm$ to use the default prefix)")
    async def prefix(app):
        prefix = ' '.join(app.args)
        if prefix == "":
            await app.say(embed=app.create_embed("Could not find arguments.",
                                "Please provide arguments for this command.",
                                                                error=True))
        else:
            del_sql = "DELETE FROM custom_prefixes WHERE server_id = %s"
            insert_sql = "INSERT INTO custom_prefixes (server_id, prefix) VALUES (%s, %s)"
            await app.run_mysql(del_sql, (app.message.guild.id,), commit=True)
            if prefix != "$rm$":
                await app.run_mysql(insert_sql, (app.message.guild.id, prefix, ), commit=True)
            else:
                prefix = app.config["default_prefix"]
            embed = app.create_embed("New prefix set:",
            f"`{prefix}` was set as the new server prefix by {app.message.author.mention}.",
            success=True)
            await app.attempt_log(app.message.guild.id, embed=embed)
            await app.say(embed=embed)
    # Allows you to set the prefix.

# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

def Plugin(app):

    @app.command("Allows you to remove merits.", requires_management=True, usage="[merit_id] (You can put multiple ID's)")
    async def rmmerits(app):
        if app.args == []:
            await app.say(embed=app.create_embed("Could not find arguments.",
                                "Please provide arguments for this command.",
                                                                error=True))
        else:
            response = ""
            for merit in app.args:
                merit_count = (await app.run_mysql("SELECT COUNT(*) AS COUNT FROM merits WHERE merit_id = %s AND server_id = %s",
                                                                            (merit, app.message.guild.id, ), get_one=True))[0]
                if merit_count == 0:
                    response = response + "\n{} - {}".format(merit, "Merit not found.")
                else:
                    await app.run_mysql("DELETE FROM merits WHERE server_id = %s AND merit_id = %s",
                                                        (app.message.guild.id, merit, ), commit=True)
                    response = response + "\n{} - {}".format(merit, "Merit removed.")
            embed = app.create_embed("Merits removed:",
                    f"{app.message.author.mention} removed the following merit(s):\n```{response}```", success=True)
            await app.say(embed=embed)
            await app.attempt_log(app.message.guild.id, embed=embed)
    # Allows you to remove merits.

# Cube. Copyright (C) Jake Gealer <jake@gealer.email> 2017-2018.
# This Source Code Form is subject to the terms of the Mozilla Public
# License, v. 2.0. If a copy of the MPL was not distributed with this
# file, You can obtain one at http://mozilla.org/MPL/2.0/.

def Plugin(app):

    @app.command("Allows you to remove strikes.", requires_management=True, usage="[strike_id] (You can put multiple ID's)")
    async def rmstrikes(app):
        if app.args == []:
            await app.say(embed=app.create_embed("Could not find arguments.",
                                "Please provide arguments for this command.",
                                                                error=True))
        else:
            response = ""
            for strike in app.args:
                strike_count = (await app.run_mysql("SELECT COUNT(*) AS COUNT FROM strikes WHERE strike_id = %s AND server_id = %s",
                                                                            (strike, app.message.guild.id, ), get_one=True))[0]
                if strike_count == 0:
                    response = response + "\n{} - {}".format(strike, "Strike not found.")
                else:
                    await app.run_mysql("DELETE FROM strikes WHERE server_id = %s AND strike_id = %s",
                                                        (app.message.guild.id, strike, ), commit=True)
                    response = response + "\n{} - {}".format(strike, "Strike removed.")
            embed = app.create_embed("Strikes removed:",
                    f"{app.message.author.mention} removed the following strike(s):\n```{response}```", success=True)
            await app.say(embed=embed)
            await app.attempt_log(app.message.guild.id, embed=embed)
    # Allows you to remove strikes.
